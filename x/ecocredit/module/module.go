package module

import (
	"context"
	"encoding/json"
	"math/rand"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	basesims "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/simulation"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	basetypesv1alpha1 "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1alpha1"
	basketsims "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/simulation"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/client"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/genesis"
	marketsims "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/simulation"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/simulation"
)

var (
	_ module.AppModule           = &Module{}
	_ module.AppModuleBasic      = Module{}
	_ module.AppModuleSimulation = Module{}
)

const (
	ConsensusVersion = 3 // ConsensusVersion is the module consensus version
)

// Module implements the AppModule interface.
type Module struct {
	key           storetypes.StoreKey
	authority     sdk.AccAddress
	Keeper        server.Keeper
	accountKeeper ecocredit.AccountKeeper
	bankKeeper    ecocredit.BankKeeper
	govKeeper     ecocredit.GovKeeper

	// legacySubspace is used solely for migration of x/ecocredit managed parameters
	legacySubspace paramtypes.Subspace
}

// NewModule returns a new Module.
func NewModule(
	storeKey storetypes.StoreKey,
	authority sdk.AccAddress,
	accountKeeper ecocredit.AccountKeeper,
	bankKeeper ecocredit.BankKeeper,
	legacySubspace paramtypes.Subspace,
	govKeeper ecocredit.GovKeeper,
) *Module {

	// legacySubspace is used solely for migration of x/ecocredit managed parameters
	if !legacySubspace.HasKeyTable() {
		legacySubspace = legacySubspace.WithKeyTable(basetypes.ParamKeyTable())
	}

	return &Module{
		key:            storeKey,
		legacySubspace: legacySubspace,
		bankKeeper:     bankKeeper,
		accountKeeper:  accountKeeper,
		authority:      authority,
		govKeeper:      govKeeper,
	}
}

/* -------------------- AppModule -------------------- */

// ConsensusVersion implements AppModule/ConsensusVersion.
func (Module) ConsensusVersion() uint64 { return ConsensusVersion }

// Name implements AppModule/Name.
func (m Module) Name() string {
	return ecocredit.ModuleName
}

// Route implements AppModule/Route.
func (m Module) Route() sdk.Route {
	return sdk.Route{}
}

// QuerierRoute implements AppModule/QuerierRoute.
func (m Module) QuerierRoute() string {
	return ecocredit.ModuleName
}

// RegisterInvariants implements AppModule/RegisterInvariants.
func (m Module) RegisterInvariants(reg sdk.InvariantRegistry) {
	m.Keeper.RegisterInvariants(reg)
}

// RegisterInterfaces implements AppModule/RegisterInterfaces.
func (m Module) RegisterInterfaces(registry types.InterfaceRegistry) {
	baskettypes.RegisterTypes(registry)
	basetypes.RegisterTypes(registry)
	markettypes.RegisterTypes(registry)

	// legacy types to support querying historical events
	basetypesv1alpha1.RegisterTypes(registry)
}

// RegisterServices implements AppModule/RegisterServices.
func (m *Module) RegisterServices(cfg module.Configurator) {
	svr := server.NewServer(m.key, m.accountKeeper, m.bankKeeper, m.authority)
	basetypes.RegisterMsgServer(cfg.MsgServer(), svr.BaseKeeper)
	basetypes.RegisterQueryServer(cfg.QueryServer(), svr.BaseKeeper)

	baskettypes.RegisterMsgServer(cfg.MsgServer(), svr.BasketKeeper)
	baskettypes.RegisterQueryServer(cfg.QueryServer(), svr.BasketKeeper)

	markettypes.RegisterMsgServer(cfg.MsgServer(), svr.MarketplaceKeeper)
	markettypes.RegisterQueryServer(cfg.QueryServer(), svr.MarketplaceKeeper)

	migrator := server.NewMigrator(svr, m.legacySubspace)
	if err := cfg.RegisterMigration(ecocredit.ModuleName, 2, migrator.Migrate2to3); err != nil {
		panic(err)
	}

	m.Keeper = svr
}

// RegisterGRPCGatewayRoutes implements AppModule/RegisterGRPCGatewayRoutes.
func (m Module) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	ctx := context.Background()
	err := basetypes.RegisterQueryHandlerClient(ctx, mux, basetypes.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
	err = baskettypes.RegisterQueryHandlerClient(ctx, mux, baskettypes.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
	err = markettypes.RegisterQueryHandlerClient(ctx, mux, markettypes.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
}

// RegisterLegacyAminoCodec implements AppModule/RegisterLegacyAminoCodec.
func (m Module) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	basetypes.RegisterLegacyAminoCodec(cdc)
	baskettypes.RegisterLegacyAminoCodec(cdc)
	markettypes.RegisterLegacyAminoCodec(cdc)
}

// InitGenesis implements AppModule/InitGenesis.
func (m Module) InitGenesis(s sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) []abci.ValidatorUpdate {
	update, err := m.Keeper.InitGenesis(s, jsonCodec, message)
	if err != nil {
		panic(err)
	}
	return update
}

// ExportGenesis implements AppModule/ExportGenesis.
func (m Module) ExportGenesis(s sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	genesis, err := m.Keeper.ExportGenesis(s, jsonCodec)
	if err != nil {
		panic(err)
	}
	return genesis
}

// DefaultGenesis implements AppModule/DefaultGenesis.
func (m Module) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	db, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	jsonTarget := ormjson.NewRawMessageTarget()
	err = db.DefaultJSON(jsonTarget)
	if err != nil {
		panic(err)
	}

	creditTypes := genesis.DefaultCreditTypes()
	err = genesis.MergeCreditTypesIntoTarget(creditTypes, jsonTarget)
	if err != nil {
		panic(err)
	}

	classFee := genesis.DefaultClassFee()
	err = genesis.MergeClassFeeIntoTarget(cdc, classFee, jsonTarget)
	if err != nil {
		panic(err)
	}

	basketFee := genesis.DefaultBasketFee()
	err = genesis.MergeBasketFeeIntoTarget(cdc, basketFee, jsonTarget)
	if err != nil {
		panic(err)
	}

	allowedDenoms := genesis.DefaultAllowedDenoms()
	err = genesis.MergeAllowedDenomsIntoTarget(allowedDenoms, jsonTarget)
	if err != nil {
		panic(err)
	}

	bz, err := jsonTarget.JSON()
	if err != nil {
		panic(err)
	}

	return bz
}

// ValidateGenesis implements AppModule/ValidateGenesis.
func (m Module) ValidateGenesis(_ codec.JSONCodec, _ sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	db, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		return err
	}

	jsonSource, err := ormjson.NewRawMessageSource(bz)
	if err != nil {
		return err
	}

	err = db.ValidateJSON(jsonSource)
	if err != nil {
		return err
	}

	return genesis.ValidateGenesis(bz)
}

// GetTxCmd implements AppModule/GetTxCmd.
func (m Module) GetTxCmd() *cobra.Command {
	return client.TxCmd(m.Name())
}

// GetQueryCmd implements AppModule/GetQueryCmd.
func (m Module) GetQueryCmd() *cobra.Command {
	return client.QueryCmd(m.Name())
}

// BeginBlock checks if there are any expired sell or buy orders and removes them from state.
func (m Module) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	err := BeginBlocker(ctx, m.Keeper)
	if err != nil {
		panic(err)
	}
}

// LegacyQuerierHandler implements AppModule/LegacyQuerierHandler.
func (m Module) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier { return nil }

/* -------------------- AppModuleSimulation -------------------- */

// GenerateGenesisState creates a randomized GenesisState of the ecocredit module.
func (Module) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}

// ProposalContents implements AppModuleSimulation/ProposalContents.
func (Module) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams implements AppModuleSimulation/RandomizedParams.
func (Module) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	return nil
}

// RegisterStoreDecoder implements AppModuleSimulation/RegisterStoreDecoder.
func (Module) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations implements AppModuleSimulation/WeightedOperations.
func (m Module) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	baseServer, basketServer, marketServer := m.Keeper.QueryServers()
	baseOps := basesims.WeightedOperations(
		simState.AppParams,
		simState.Cdc,
		m.accountKeeper,
		m.bankKeeper,
		m.govKeeper,
		baseServer,
		basketServer,
		marketServer,
		m.authority,
	)

	basketOps := basketsims.WeightedOperations(
		simState.AppParams,
		simState.Cdc,
		m.accountKeeper,
		m.bankKeeper,
		m.govKeeper,
		baseServer,
		basketServer,
		m.authority,
	)

	marketplaceOps := marketsims.WeightedOperations(simState.AppParams,
		simState.Cdc,
		m.accountKeeper,
		m.bankKeeper,
		baseServer,
		marketServer,
		m.govKeeper,
		m.authority,
	)

	return append(append(baseOps, basketOps...), marketplaceOps...)
}
