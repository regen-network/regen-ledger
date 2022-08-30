package server

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	marketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormstore"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	coretypes "github.com/regen-network/regen-ledger/x/ecocredit/core"
	marketkeeper "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/keeper"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/core"
)

type serverImpl struct {
	legacySubspace paramtypes.Subspace
	bankKeeper     ecocredit.BankKeeper
	accountKeeper  ecocredit.AccountKeeper

	CoreKeeper        core.Keeper
	BasketKeeper      basket.Keeper
	MarketplaceKeeper marketkeeper.Keeper

	db               ormdb.ModuleDB
	stateStore       api.StateStore
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
	coreAddr := s.accountKeeper.GetModuleAddress(ecocredit.ModuleName)
	if coreAddr == nil {
		panic(fmt.Sprintf("%s module account has not been set", ecocredit.ModuleName))
	}

	// ensure basket submodule account is set
	basketAddr := s.accountKeeper.GetModuleAddress(baskettypes.BasketSubModuleName)
	if basketAddr == nil {
		panic(fmt.Sprintf("%s module account has not been set", baskettypes.BasketSubModuleName))
	}

	var err error
	s.db, err = ormstore.NewStoreKeyDB(&ecocredit.ModuleSchema, storeKey, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	coreStore, basketStore, marketStore := getStateStores(s.db)
	s.stateStore = coreStore
	s.basketStore = basketStore
	s.marketplaceStore = marketStore
	s.CoreKeeper = core.NewKeeper(coreStore, bankKeeper, coreAddr, basketStore, authority)
	s.BasketKeeper = basket.NewKeeper(basketStore, coreStore, bankKeeper, s.legacySubspace, basketAddr, authority)
	s.MarketplaceKeeper = marketkeeper.NewKeeper(marketStore, coreStore, bankKeeper, s.legacySubspace, authority)

	return s
}

func getStateStores(db ormdb.ModuleDB) (api.StateStore, basketapi.StateStore, marketapi.StateStore) {
	coreStore, err := api.NewStateStore(db)
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
	return coreStore, basketStore, marketStore
}

func (s serverImpl) QueryServers() (coretypes.QueryServer, baskettypes.QueryServer, markettypes.QueryServer) {
	return s.CoreKeeper, s.BasketKeeper, s.MarketplaceKeeper
}

func (s serverImpl) GetStateStores() (api.StateStore, basketapi.StateStore, marketapi.StateStore) {
	return s.stateStore, s.basketStore, s.marketplaceStore
}
