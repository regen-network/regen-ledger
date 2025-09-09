package server

import (
	"fmt"

	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/orm/model/ormdb"

	basketapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	marketapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	baseapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/v2/ormstore"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	basekeeper "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/keeper"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/basket"
	basketkeeper "github.com/regen-network/regen-ledger/x/ecocredit/v4/basket/keeper"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/basket/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace"
	marketkeeper "github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace/keeper"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace/types/v1"
)

type serverImpl struct {
	bankKeeper    ecocredit.BankKeeper
	accountKeeper ecocredit.AccountKeeper

	baseKeeper   basekeeper.Keeper
	basketKeeper basketkeeper.Keeper
	marketKeeper marketkeeper.Keeper

	db               ormdb.ModuleDB
	stateStore       baseapi.StateStore
	basketStore      basketapi.StateStore
	marketplaceStore marketapi.StateStore
}

//nolint:revive
func NewServer(storeKey storetypes.StoreKey,
	accountKeeper ecocredit.AccountKeeper, bankKeeper ecocredit.BankKeeper, authority sdk.AccAddress,
) serverImpl {
	s := serverImpl{
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
	}

	// ensure ecocredit module account is set
	baseAddr := s.accountKeeper.GetModuleAddress(ecocredit.ModuleName)
	if baseAddr == nil {
		panic(fmt.Sprintf("%s module account has not been set", ecocredit.ModuleName))
	}

	// ensure basket submodule account is set
	basketAddr := s.accountKeeper.GetModuleAddress(basket.BasketSubModuleName)
	if basketAddr == nil {
		panic(fmt.Sprintf("%s module account has not been set", basket.BasketSubModuleName))
	}

	// ensure marketplace fee pool account is set
	marketplaceAddr := s.accountKeeper.GetModuleAddress(marketplace.FeePoolName)
	if marketplaceAddr == nil {
		panic(fmt.Sprintf("%s module account has not been set", marketplace.FeePoolName))
	}

	var err error
	s.db, err = ormstore.NewStoreKeyDB(&ecocredit.ModuleSchema, storeKey, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	baseStore, basketStore, marketStore := getStateStores(s.db)
	s.stateStore = baseStore
	s.basketStore = basketStore
	s.marketplaceStore = marketStore
	s.baseKeeper = basekeeper.NewKeeper(baseStore, bankKeeper, baseAddr, basketStore, marketStore, authority)
	s.basketKeeper = basketkeeper.NewKeeper(basketStore, baseStore, bankKeeper, basketAddr, authority)
	s.marketKeeper = marketkeeper.NewKeeper(marketStore, baseStore, bankKeeper, authority)

	return s
}

func getStateStores(db ormdb.ModuleDB) (baseapi.StateStore, basketapi.StateStore, marketapi.StateStore) {
	baseStore, err := baseapi.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	basketStore, err := basketapi.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	marketStore, err := marketapi.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	return baseStore, basketStore, marketStore
}

func (s serverImpl) QueryServers() (basetypes.QueryServer, baskettypes.QueryServer, markettypes.QueryServer) {
	return s.baseKeeper, s.basketKeeper, s.marketKeeper
}

func (s serverImpl) GetStateStores() (baseapi.StateStore, basketapi.StateStore, marketapi.StateStore) {
	return s.stateStore, s.basketStore, s.marketplaceStore
}

func (s serverImpl) GetBaseKeeper() basekeeper.Keeper {
	return s.baseKeeper
}

func (s serverImpl) GetBasketKeeper() basketkeeper.Keeper {
	return s.basketKeeper
}

func (s serverImpl) GetMarketKeeper() marketkeeper.Keeper {
	return s.marketKeeper
}
