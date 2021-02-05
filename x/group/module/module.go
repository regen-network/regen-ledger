package module

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	climodule "github.com/regen-network/regen-ledger/types/module/client/cli"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/regen-network/regen-ledger/x/group/server"
	"github.com/regen-network/regen-ledger/x/group/simulation"
)

type Module struct {
	registry      types.InterfaceRegistry
	bankKeeper    group.BankKeeper
	AccountKeeper server.AccountKeeper
}

var _ module.AppModuleBasic = Module{}
var _ module.AppModuleSimulation = Module{}
var _ servermodule.Module = Module{}
var _ climodule.Module = Module{}

func (a Module) Name() string {
	return group.ModuleName
}

// RegisterInterfaces registers module concrete types into protobuf Any.
func (a Module) RegisterInterfaces(registry types.InterfaceRegistry) {
	group.RegisterTypes(registry)
}

func (a Module) RegisterServices(configurator servermodule.Configurator) {
	server.RegisterServices(configurator, a.AccountKeeper)
}

func (a Module) DefaultGenesis(marshaler codec.JSONMarshaler) json.RawMessage {
	return marshaler.MustMarshalJSON(group.NewGenesisState())
}

func (a Module) ValidateGenesis(cdc codec.JSONMarshaler, config client.TxEncodingConfig, bz json.RawMessage) error {
	var data group.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", group.ModuleName, err)
	}
	return data.Validate()
}

func (a Module) RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux) {}

func (a Module) GetTxCmd() *cobra.Command {
	// TODO #209
	return nil
}

func (a Module) GetQueryCmd() *cobra.Command {
	// TODO #209
	return nil
}

func (a Module) GenerateGenesisState(input *module.SimulationState) {
}

func (a Module) ProposalContents(simState module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

func (a Module) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	return nil
}

func (a Module) RegisterStoreDecoder(registry sdk.StoreDecoderRegistry) {
}

func (a Module) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	protoCdc := codec.NewProtoCodec(a.registry)
	return simulation.WeightedOperations(
		simState.AppParams, simState.Cdc,
		a.AccountKeeper, a.bankKeeper, protoCdc,
	)
}

/**** DEPRECATED ****/
func (a Module) RegisterRESTRoutes(client.Context, *mux.Router) {}
func (a Module) RegisterLegacyAminoCodec(*codec.LegacyAmino)    {}
