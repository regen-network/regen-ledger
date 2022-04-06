package server

import (
	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/types/ormstore"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	coretypes "github.com/regen-network/regen-ledger/x/ecocredit/core"
	marketplacetypes "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/marketplace"
)

type serverImpl struct {
	basketKeeper basket.Keeper
	coreKeeper   core.Keeper
	marketKeeper marketplace.Keeper

	storeKey sdk.StoreKey

	paramSpace    paramtypes.Subspace
	bankKeeper    ecocredit.BankKeeper
	accountKeeper ecocredit.AccountKeeper

	db ormdb.ModuleDB
}

func newServer(storeKey sdk.StoreKey, paramSpace paramtypes.Subspace,
	accountKeeper ecocredit.AccountKeeper, bankKeeper ecocredit.BankKeeper, distKeeper ecocredit.DistributionKeeper, cdc codec.Codec) serverImpl {
	s := serverImpl{
		storeKey:      storeKey,
		paramSpace:    paramSpace,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
	}

	var err error
	s.db, err = ormstore.NewStoreKeyDB(&ecocredit.ModuleSchema, storeKey, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	coreStore, basketStore, marketStore := getStateStores(s.db)
	s.basketKeeper = basket.NewKeeper(basketStore, coreStore, bankKeeper, distKeeper, s.paramSpace)
	s.coreKeeper = core.NewKeeper(coreStore, bankKeeper, s.paramSpace)
	s.marketKeeper = marketplace.NewKeeper(marketStore, coreStore, bankKeeper, s.paramSpace)

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

func RegisterServices(
	configurator server.Configurator,
	paramSpace paramtypes.Subspace,
	accountKeeper ecocredit.AccountKeeper,
	bankKeeper ecocredit.BankKeeper,
	distKeeper ecocredit.DistributionKeeper,
) ecocredit.Keeper {
	impl := newServer(configurator.ModuleKey(), paramSpace, accountKeeper, bankKeeper, distKeeper, configurator.Marshaler())

	baskettypes.RegisterMsgServer(configurator.MsgServer(), impl.basketKeeper)
	baskettypes.RegisterQueryServer(configurator.QueryServer(), impl.basketKeeper)

	coretypes.RegisterMsgServer(configurator.MsgServer(), impl.coreKeeper)
	coretypes.RegisterQueryServer(configurator.QueryServer(), impl.coreKeeper)

	marketplacetypes.RegisterMsgServer(configurator.MsgServer(), impl.marketKeeper)
	marketplacetypes.RegisterQueryServer(configurator.QueryServer(), impl.marketKeeper)

	configurator.RegisterGenesisHandlers(impl.InitGenesis, impl.ExportGenesis)
	// TODO: uncomment when sims are refactored https://github.com/regen-network/regen-ledger/issues/920
	// configurator.RegisterWeightedOperationsHandler(impl.WeightedOperations)
	configurator.RegisterInvariantsHandler(impl.RegisterInvariants)
	return impl
}
