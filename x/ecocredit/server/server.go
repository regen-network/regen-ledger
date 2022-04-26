package server

import (
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
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
	storeKey sdk.StoreKey

	paramSpace    paramtypes.Subspace
	bankKeeper    ecocredit.BankKeeper
	accountKeeper ecocredit.AccountKeeper

	basketKeeper      basket.Keeper
	coreKeeper        core.Keeper
	marketplaceKeeper marketplace.Keeper

	db         ormdb.ModuleDB
	stateStore api.StateStore
}

func (s serverImpl) AddCreditType(ctx sdk.Context, ctp *coretypes.CreditTypeProposal) error {
	return s.coreKeeper.AddCreditType(ctx, ctp)
}

func newServer(storeKey sdk.StoreKey, paramSpace paramtypes.Subspace,
	accountKeeper ecocredit.AccountKeeper, bankKeeper ecocredit.BankKeeper, distKeeper ecocredit.DistributionKeeper) serverImpl {
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
	s.marketplaceKeeper = marketplace.NewKeeper(marketStore, coreStore, bankKeeper, s.paramSpace)

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
) Keeper {
	impl := newServer(configurator.ModuleKey(), paramSpace, accountKeeper, bankKeeper, distKeeper)

	coretypes.RegisterMsgServer(configurator.MsgServer(), impl.coreKeeper)
	coretypes.RegisterQueryServer(configurator.QueryServer(), impl.coreKeeper)

	baskettypes.RegisterMsgServer(configurator.MsgServer(), impl.basketKeeper)
	baskettypes.RegisterQueryServer(configurator.QueryServer(), impl.basketKeeper)

	marketplacetypes.RegisterMsgServer(configurator.MsgServer(), impl.marketplaceKeeper)
	marketplacetypes.RegisterQueryServer(configurator.QueryServer(), impl.marketplaceKeeper)

	configurator.RegisterGenesisHandlers(impl.InitGenesis, impl.ExportGenesis)
	configurator.RegisterMigrationHandler(impl.RunMigrations)

	configurator.RegisterWeightedOperationsHandler(impl.WeightedOperations)
	configurator.RegisterInvariantsHandler(impl.RegisterInvariants)
	return impl
}
