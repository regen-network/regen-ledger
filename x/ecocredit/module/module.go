package module

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/reflect/protoreflect"

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

	"github.com/regen-network/regen-ledger/x/ecocredit"
	basesims "github.com/regen-network/regen-ledger/x/ecocredit/base/simulation"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	basetypesv1alpha1 "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1alpha1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/client"
	"github.com/regen-network/regen-ledger/x/ecocredit/genesis"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation"
)

var (
	_ module.AppModule           = &Module{}
	_ module.AppModuleBasic      = Module{}
	_ module.AppModuleSimulation = Module{}
)

type Module struct {
	key storetypes.StoreKey
	// legacySubspace is used solely for migration of x/ecocredit managed parameters
	legacySubspace paramtypes.Subspace
	accountKeeper  ecocredit.AccountKeeper
	bankKeeper     ecocredit.BankKeeper
	Keeper         server.Keeper
	authority      sdk.AccAddress
}

func (a Module) InitGenesis(s sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) []abci.ValidatorUpdate {
	update, err := a.Keeper.InitGenesis(s, jsonCodec, message)
	if err != nil {
		panic(err)
	}
	return update
}

func (a Module) ExportGenesis(s sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	m, err := a.Keeper.ExportGenesis(s, jsonCodec)
	if err != nil {
		panic(err)
	}
	return m
}

func (a Module) RegisterInvariants(reg sdk.InvariantRegistry) {
	a.Keeper.RegisterInvariants(reg)
}

func (a Module) Route() sdk.Route {
	return sdk.Route{}
}

func (a Module) QuerierRoute() string {
	return ecocredit.ModuleName
}

func (a Module) LegacyQuerierHandler(amino *codec.LegacyAmino) sdk.Querier { return nil }

// NewModule returns a new Module object.
func NewModule(
	storeKey storetypes.StoreKey,
	legacySubspace paramtypes.Subspace,
	accountKeeper ecocredit.AccountKeeper,
	bankKeeper ecocredit.BankKeeper,
	authority sdk.AccAddress,
) *Module {
	if !legacySubspace.HasKeyTable() {
		legacySubspace = legacySubspace.WithKeyTable(basetypes.ParamKeyTable())
	}

	return &Module{
		key:            storeKey,
		legacySubspace: legacySubspace,
		bankKeeper:     bankKeeper,
		accountKeeper:  accountKeeper,
		authority:      authority,
	}
}

var _ module.AppModuleBasic = &Module{}
var _ module.AppModuleSimulation = &Module{}

func (a Module) Name() string {
	return ecocredit.ModuleName
}

func (a Module) RegisterInterfaces(registry types.InterfaceRegistry) {
	baskettypes.RegisterTypes(registry)
	basetypes.RegisterTypes(registry)
	markettypes.RegisterTypes(registry)

	// legacy types to support querying historical events
	basetypesv1alpha1.RegisterTypes(registry)
}

func (a *Module) RegisterServices(cfg module.Configurator) {
	svr := server.NewServer(a.key, a.legacySubspace, a.accountKeeper, a.bankKeeper, a.authority)
	basetypes.RegisterMsgServer(cfg.MsgServer(), svr.CoreKeeper)
	basetypes.RegisterQueryServer(cfg.QueryServer(), svr.CoreKeeper)

	baskettypes.RegisterMsgServer(cfg.MsgServer(), svr.BasketKeeper)
	baskettypes.RegisterQueryServer(cfg.QueryServer(), svr.BasketKeeper)

	markettypes.RegisterMsgServer(cfg.MsgServer(), svr.MarketplaceKeeper)
	markettypes.RegisterQueryServer(cfg.QueryServer(), svr.MarketplaceKeeper)

	m := server.NewMigrator(svr, a.legacySubspace)
	if err := cfg.RegisterMigration(ecocredit.ModuleName, 2, m.Migrate2to3); err != nil {
		panic(err)
	}
	a.Keeper = svr
}

//nolint:errcheck
func (a Module) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	ctx := context.Background()
	basetypes.RegisterQueryHandlerClient(ctx, mux, basetypes.NewQueryClient(clientCtx))
	baskettypes.RegisterQueryHandlerClient(ctx, mux, baskettypes.NewQueryClient(clientCtx))
	markettypes.RegisterQueryHandlerClient(ctx, mux, markettypes.NewQueryClient(clientCtx))
}

func (a Module) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
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

	creditClassFees := genesis.DefaultCreditClassFees()
	err = genesis.MergeCreditClassFeesIntoTarget(cdc, creditClassFees, jsonTarget)
	if err != nil {
		panic(err)
	}

	basketFees := genesis.DefaultBasketFees()
	err = genesis.MergeBasketFeesIntoTarget(cdc, basketFees, jsonTarget)
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

func (a Module) ValidateGenesis(cdc codec.JSONCodec, _ sdkclient.TxEncodingConfig, bz json.RawMessage) error {
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

	var params basetypes.Params
	r, err := jsonSource.OpenReader(protoreflect.FullName(proto.MessageName(&params)))
	if err != nil {
		return err
	}

	if r == nil {
		return nil
	}

	if err := (&jsonpb.Unmarshaler{AllowUnknownFields: true}).Unmarshal(r, &params); err != nil {
		return fmt.Errorf("failed to unmarshal %s params state: %w", ecocredit.ModuleName, err)
	}

	return genesis.ValidateGenesis(bz, params)
}

func (a Module) GetQueryCmd() *cobra.Command {
	return client.QueryCmd(a.Name())
}

func (a Module) GetTxCmd() *cobra.Command {
	return client.TxCmd(a.Name())
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (Module) ConsensusVersion() uint64 { return 3 }

/**** DEPRECATED ****/
func (a Module) RegisterRESTRoutes(sdkclient.Context, *mux.Router) {}
func (a Module) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	basetypes.RegisterLegacyAminoCodec(cdc)
	baskettypes.RegisterLegacyAminoCodec(cdc)
	markettypes.RegisterLegacyAminoCodec(cdc)
}

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenesisState of the ecocredit module.
func (Module) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}

// ProposalContents returns all the ecocredit content functions used to
// simulate proposals.
func (Module) ProposalContents(simState module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized ecocredit param changes for the simulator.
func (Module) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	return nil
}

// RegisterStoreDecoder registers a decoder for ecocredit module's types
func (Module) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns all the ecocredit module operations with their respective weights.
func (a Module) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	coreQuerier, basketQuerier, marketQuerier := a.Keeper.QueryServers()
	return basesims.WeightedOperations(
		simState.AppParams, simState.Cdc,
		a.accountKeeper, a.bankKeeper,
		coreQuerier,
		basketQuerier,
		marketQuerier,
	)
}

// BeginBlock checks if there are any expired sell or buy orders and removes them from state.
func (a Module) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
	err := server.BeginBlocker(ctx, a.Keeper)
	if err != nil {
		panic(err)
	}
}
