package module

import (
	"encoding/json"
	"math/rand"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	climodule "github.com/regen-network/regen-ledger/types/module/client/cli"
	restmodule "github.com/regen-network/regen-ledger/types/module/client/grpc_gateway"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/bridge"
	"github.com/regen-network/regen-ledger/x/bridge/client"
	"github.com/regen-network/regen-ledger/x/bridge/server"
	"github.com/regen-network/regen-ledger/x/bridge/simulation"
)

type Module struct {
	router *baseapp.MsgServiceRouter
}

var _ module.AppModuleBasic = Module{}
var _ servermodule.Module = Module{}
var _ restmodule.Module = Module{}
var _ climodule.Module = Module{}
var _ module.AppModuleSimulation = &Module{}

func NewModule(router *baseapp.MsgServiceRouter) Module {
	return Module{
		router: router,
	}
}

func (a Module) Name() string {
	return "bridge"
}

func (a Module) RegisterInterfaces(registry types.InterfaceRegistry) {
	bridge.RegisterTypes(registry)
}

func (a Module) RegisterServices(configurator servermodule.Configurator) {
	server.RegisterServices(configurator, a.ak, a.bk)
}

//nolint:errcheck
func (a Module) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {

}

func (a Module) DefaultGenesis(codec.JSONCodec) json.RawMessage {
	db, err := ormdb.NewModuleDB(&server.ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	jsonTarget := ormjson.NewRawMessageTarget()
	err = db.DefaultJSON(jsonTarget)
	if err != nil {
		panic(err)
	}

	bz, err := jsonTarget.JSON()
	if err != nil {
		panic(err)
	}

	return bz
}

func (a Module) ValidateGenesis(_ codec.JSONCodec, _ sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	db, err := ormdb.NewModuleDB(&server.ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		return err
	}

	jsonSource, err := ormjson.NewRawMessageSource(bz)
	if err != nil {
		return err
	}

	return db.ValidateJSON(jsonSource)
}

func (a Module) GetQueryCmd() *cobra.Command {
	return client.QueryCmd(a.Name())
}

func (a Module) GetTxCmd() *cobra.Command {
	return client.TxCmd(a.Name())
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (Module) ConsensusVersion() uint64 { return 1 }

/**** DEPRECATED ****/
func (a Module) RegisterRESTRoutes(sdkclient.Context, *mux.Router) {}

// RegisterLegacyAminoCodec registers the bridge module's types on the given LegacyAmino codec.
func (a Module) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	bridge.RegisterLegacyAminoCodec(cdc)
}

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenesisState of the bridge module.
func (Module) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}

// ProposalContents returns all the bridge content functions used to
// simulate proposals.
func (Module) ProposalContents(simState module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized bridge param changes for the simulator.
func (Module) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	return nil
}

// RegisterStoreDecoder registers a decoder for bridge module's types
func (Module) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
}

// WeightedOperations returns all the bridge module operations with their respective weights.
// NOTE: This is no longer needed for the modules which uses ADR-33, bridge module `WeightedOperations`
// registered in the `x/bridge/server` package.
func (Module) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return nil
}
