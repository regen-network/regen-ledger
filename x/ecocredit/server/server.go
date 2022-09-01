package server

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	marketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	baseapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormstore"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	basekeeper "github.com/regen-network/regen-ledger/x/ecocredit/base/keeper"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	basketkeeper "github.com/regen-network/regen-ledger/x/ecocredit/basket/keeper"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
	marketkeeper "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/keeper"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
)

type serverImpl struct {
	legacySubspace paramtypes.Subspace
	bankKeeper     ecocredit.BankKeeper
	accountKeeper  ecocredit.AccountKeeper

	BaseKeeper        basekeeper.Keeper
	BasketKeeper      basketkeeper.Keeper
	MarketplaceKeeper marketkeeper.Keeper

	db               ormdb.ModuleDB
	stateStore       baseapi.StateStore
	basketStore      basketapi.StateStore
	marketplaceStore marketapi.StateStore
}

//nolint:revive
func NewServer(storeKey storetypes.StoreKey, legacySubspace paramtypes.Subspace,
	accountKeeper ecocredit.AccountKeeper, bankKeeper ecocredit.BankKeeper, authority sdk.AccAddress) serverImpl {
	s := serverImpl{
		legacySubspace: legacySubspace,
		bankKeeper:     bankKeeper,
		accountKeeper:  accountKeeper,
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

	var err error
	s.db, err = ormstore.NewStoreKeyDB(&ecocredit.ModuleSchema, storeKey, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	baseStore, basketStore, marketStore := getStateStores(s.db)
	s.stateStore = baseStore
	s.basketStore = basketStore
	s.marketplaceStore = marketStore
	s.BaseKeeper = basekeeper.NewKeeper(baseStore, bankKeeper, baseAddr, basketStore, marketStore, authority)
	s.BasketKeeper = basketkeeper.NewKeeper(basketStore, baseStore, bankKeeper, s.legacySubspace, basketAddr, authority)
	s.MarketplaceKeeper = marketkeeper.NewKeeper(marketStore, baseStore, bankKeeper, s.legacySubspace, authority)

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
	return s.BaseKeeper, s.BasketKeeper, s.MarketplaceKeeper
}

func (s serverImpl) GetStateStores() (baseapi.StateStore, basketapi.StateStore, marketapi.StateStore) {
	return s.stateStore, s.basketStore, s.marketplaceStore
}
