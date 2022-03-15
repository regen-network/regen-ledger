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

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	abci "github.com/tendermint/tendermint/abci/types"

	climodule "github.com/regen-network/regen-ledger/types/module/client/cli"
	restmodule "github.com/regen-network/regen-ledger/types/module/client/grpc_gateway"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/client"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation"
)

type Module struct {
	paramSpace         paramtypes.Subspace
	accountKeeper      ecocredit.AccountKeeper
	bankKeeper         ecocredit.BankKeeper
	distributionKeeper ecocredit.DistributionKeeper
	keeper             ecocredit.Keeper
}

// NewModule returns a new Module object.
func NewModule(
	paramSpace paramtypes.Subspace,
	accountKeeper ecocredit.AccountKeeper,
	bankKeeper ecocredit.BankKeeper,
	distributionKeeper ecocredit.DistributionKeeper,
) *Module {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(ecocredit.ParamKeyTable())
	}

	return &Module{
		paramSpace:         paramSpace,
		bankKeeper:         bankKeeper,
		accountKeeper:      accountKeeper,
		distributionKeeper: distributionKeeper,
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
	ecocredit.RegisterTypes(registry)
	baskettypes.RegisterTypes(registry)
}

func (a *Module) RegisterServices(configurator servermodule.Configurator) {
	a.keeper = server.RegisterServices(configurator, a.paramSpace, a.accountKeeper, a.bankKeeper, a.distributionKeeper)
}

//nolint:errcheck
func (a Module) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	ctx := context.Background()
	ecocredit.RegisterQueryHandlerClient(ctx, mux, ecocredit.NewQueryClient(clientCtx))
	baskettypes.RegisterQueryHandlerClient(ctx, mux, baskettypes.NewQueryClient(clientCtx))
}

func (a Module) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	db, err := ormdb.NewModuleDB(server.ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	jsonTarget := ormjson.NewRawMessageTarget()
	err = db.DefaultJSON(jsonTarget)
	if err != nil {
		panic(err)
	}

	err = server.MergeLegacyJSONIntoTarget(cdc, ecocredit.DefaultGenesisState(), jsonTarget)
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
	db, err := ormdb.NewModuleDB(server.ModuleSchema, ormdb.ModuleDBOptions{})
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

	var data ecocredit.GenesisState

	r, err := jsonSource.OpenReader(protoreflect.FullName(proto.MessageName(&data)))
	if err != nil {
		return err
	}

	if r == nil {
		return nil
	}

	if err := (&jsonpb.Unmarshaler{AllowUnknownFields: true}).Unmarshal(r, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", ecocredit.ModuleName, err)
	}

	return data.Validate()
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
	ecocredit.RegisterLegacyAminoCodec(cdc)
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
	return simulation.ParamChanges(r)
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
	err := ecocredit.BeginBlocker(ctx, a.keeper)
	if err != nil {
		panic(err)
	}
}
