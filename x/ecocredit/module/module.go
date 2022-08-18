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
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	climodule "github.com/regen-network/regen-ledger/types/module/client/cli"
	restmodule "github.com/regen-network/regen-ledger/types/module/client/grpc_gateway"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/client"
	coretypes "github.com/regen-network/regen-ledger/x/ecocredit/core"
	corev1alpha1 "github.com/regen-network/regen-ledger/x/ecocredit/core/v1alpha1"
	"github.com/regen-network/regen-ledger/x/ecocredit/genesis"
	marketplacetypes "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation"
)

type Module struct {
	paramSpace    paramtypes.Subspace
	accountKeeper ecocredit.AccountKeeper
	bankKeeper    ecocredit.BankKeeper
	Keeper        server.Keeper
	authority     sdk.AccAddress
}

// NewModule returns a new Module object.
func NewModule(
	paramSpace paramtypes.Subspace,
	accountKeeper ecocredit.AccountKeeper,
	bankKeeper ecocredit.BankKeeper,
	authority sdk.AccAddress,
) *Module {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(coretypes.ParamKeyTable())
	}

	return &Module{
		paramSpace:    paramSpace,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		authority:     authority,
	}
}

var _ module.AppModuleBasic = &Module{}
var _ servermodule.Module = &Module{}
var _ restmodule.Module = &Module{}
var _ climodule.Module = &Module{}
var _ module.AppModuleSimulation = &Module{}

func (a Module) Name() string {
	return ecocredit.ModuleName
}

func (a Module) RegisterInterfaces(registry types.InterfaceRegistry) {
	baskettypes.RegisterTypes(registry)
	coretypes.RegisterTypes(registry)
	marketplacetypes.RegisterTypes(registry)

	// legacy types to support querying historical events
	corev1alpha1.RegisterTypes(registry)
}

func (a *Module) RegisterServices(configurator servermodule.Configurator) {
	a.Keeper = server.RegisterServices(configurator, a.paramSpace, a.accountKeeper, a.bankKeeper, a.authority)
}

//nolint:errcheck
func (a Module) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	ctx := context.Background()
	baskettypes.RegisterQueryHandlerClient(ctx, mux, baskettypes.NewQueryClient(clientCtx))
	marketplacetypes.RegisterQueryHandlerClient(ctx, mux, marketplacetypes.NewQueryClient(clientCtx))
	coretypes.RegisterQueryHandlerClient(ctx, mux, coretypes.NewQueryClient(clientCtx))
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

	params := coretypes.DefaultParams()
	err = genesis.MergeParamsIntoTarget(cdc, &params, jsonTarget)
	if err != nil {
		panic(err)
	}

	creditTypes := coretypes.DefaultCreditTypes()
	err = genesis.MergeCreditTypesIntoTarget(creditTypes, jsonTarget)
	if err != nil {
		panic(err)
	}

	allowedDenoms := marketplacetypes.DefaultAllowedDenoms()
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

	var params coretypes.Params
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
func (Module) ConsensusVersion() uint64 { return 2 }

/**** DEPRECATED ****/
func (a Module) RegisterRESTRoutes(sdkclient.Context, *mux.Router) {}
func (a Module) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	baskettypes.RegisterLegacyAminoCodec(cdc)
	coretypes.RegisterLegacyAminoCodec(cdc)
	marketplacetypes.RegisterLegacyAminoCodec(cdc)
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
	return simulation.ParamChanges()
}

// RegisterStoreDecoder registers a decoder for ecocredit module's types
func (Module) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
}

// WeightedOperations returns all the ecocredit module operations with their respective weights.
// NOTE: This is no longer needed for the modules which uses ADR-33, ecocredit module `WeightedOperations`
// registered in the `x/ecocredit/server` package.
func (Module) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return nil
}

// BeginBlock checks if there are any expired sell or buy orders and removes them from state.
func (a Module) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
	err := server.BeginBlocker(ctx, a.Keeper)
	if err != nil {
		panic(err)
	}
}
