package server

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormstore"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/marketplace"
)

type serverImpl struct {
	storeKey storetypes.StoreKey

	paramSpace    paramtypes.Subspace
	bankKeeper    ecocredit.BankKeeper
	accountKeeper ecocredit.AccountKeeper

	CoreKeeper        core.Keeper
	BasketKeeper      basket.Keeper
	MarketplaceKeeper marketplace.Keeper

	db          ormdb.ModuleDB
	stateStore  api.StateStore
	basketStore basketapi.StateStore
}

func NewServer(storeKey storetypes.StoreKey, paramSpace paramtypes.Subspace,
	accountKeeper ecocredit.AccountKeeper, bankKeeper ecocredit.BankKeeper, authority sdk.AccAddress) serverImpl {
	s := serverImpl{
		storeKey:      storeKey,
		paramSpace:    paramSpace,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
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
	s.CoreKeeper = core.NewKeeper(coreStore, bankKeeper, s.paramSpace, coreAddr, authority)
	s.BasketKeeper = basket.NewKeeper(basketStore, coreStore, bankKeeper, s.paramSpace, basketAddr)
	s.MarketplaceKeeper = marketplace.NewKeeper(marketStore, coreStore, bankKeeper, s.paramSpace, authority)

	return s
}

func getStateStores(db ormdb.ModuleDB) (api.StateStore, basketapi.StateStore, marketApi.StateStore) {
	coreStore, err := api.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	basketStore, err := basketapi.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	marketStore, err := marketApi.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	return coreStore, basketStore, marketStore
}

// TODO(Tyler): still need to figure out weighted sim handlers??

//func RegisterServices(
//	configurator server.Configurator,
//	paramSpace paramtypes.Subspace,
//	accountKeeper ecocredit.AccountKeeper,
//	bankKeeper ecocredit.BankKeeper,
//	authority sdk.AccAddress,
//) Keeper {
//	impl := NewServer(configurator.ModuleKey(), paramSpace, accountKeeper, bankKeeper, authority)
//
//	coretypes.RegisterMsgServer(configurator.MsgServer(), impl.CoreKeeper)
//	coretypes.RegisterQueryServer(configurator.QueryServer(), impl.CoreKeeper)
//
//	baskettypes.RegisterMsgServer(configurator.MsgServer(), impl.BasketKeeper)
//	baskettypes.RegisterQueryServer(configurator.QueryServer(), impl.BasketKeeper)
//
//	marketplacetypes.RegisterMsgServer(configurator.MsgServer(), impl.MarketplaceKeeper)
//	marketplacetypes.RegisterQueryServer(configurator.QueryServer(), impl.MarketplaceKeeper)
//
//	configurator.RegisterGenesisHandlers(impl.InitGenesis, impl.ExportGenesis)
//	configurator.RegisterMigrationHandler(impl.RunMigrations)
//
//	configurator.RegisterWeightedOperationsHandler(impl.WeightedOperations)
//	configurator.RegisterInvariantsHandler(impl.RegisterInvariants)
//	return impl
//}
