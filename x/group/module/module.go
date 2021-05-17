package module

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	climodule "github.com/regen-network/regen-ledger/types/module/client/cli"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/regen-network/regen-ledger/x/group/client"

	"github.com/regen-network/regen-ledger/x/group/exported"
	"github.com/regen-network/regen-ledger/x/group/server"
	"github.com/regen-network/regen-ledger/x/group/simulation"

	servermodule "github.com/regen-network/regen-ledger/types/module/server"
)

type Module struct {
	Registry      types.InterfaceRegistry
	BankKeeper    exported.BankKeeper
	AccountKeeper exported.AccountKeeper
}

var _ module.AppModuleBasic = Module{}
var _ module.AppModuleSimulation = Module{}
var _ servermodule.Module = Module{}
var _ climodule.Module = Module{}
var _ servermodule.LegacyRouteModule = Module{}

func (a Module) Name() string {
	return group.ModuleName
}

// RegisterInterfaces registers module concrete types into protobuf Any.
func (a Module) RegisterInterfaces(registry types.InterfaceRegistry) {
	group.RegisterTypes(registry)
}

func (a Module) RegisterServices(configurator servermodule.Configurator) {
	server.RegisterServices(configurator, a.AccountKeeper, a.BankKeeper)
}

func (a Module) DefaultGenesis(marshaler codec.JSONMarshaler) json.RawMessage {
	return marshaler.MustMarshalJSON(group.NewGenesisState())
}

func (a Module) ValidateGenesis(cdc codec.JSONMarshaler, config sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	var data group.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", group.ModuleName, err)
	}
	return data.Validate()
}

func (a Module) RegisterGRPCGatewayRoutes(sdkclient.Context, *runtime.ServeMux) {}

func (a Module) GetTxCmd() *cobra.Command {
	return client.TxCmd(a.Name())
}

func (a Module) GetQueryCmd() *cobra.Command {
	return client.QueryCmd(a.Name())
}

/**** DEPRECATED ****/
func (a Module) RegisterRESTRoutes(sdkclient.Context, *mux.Router) {}
func (a Module) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	group.RegisterLegacyAminoCodec(cdc)
}

func (a Module) Route(configurator servermodule.Configurator) sdk.Route {
	return sdk.NewRoute(group.RouterKey, server.NewHandler(configurator, a.AccountKeeper, a.BankKeeper))
}

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenesisState of the group module.
func (Module) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}

// ProposalContents returns all the group content functions used to
// simulate proposals.
func (Module) ProposalContents(simState module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized group param changes for the simulator.
func (Module) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	return nil
}

// RegisterStoreDecoder registers a decoder for group module's types
func (Module) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
}

// WeightedOperations returns all the group module operations with their respective weights.
// NOTE: This is no longer needed for the modules which uses ADR-33, group module `WeightedOperations`
// registered in the `x/group/server` package.
func (Module) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return nil
}
