package module

import (
	"context"
	"encoding/json"
	"fmt"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	climodule "github.com/regen-network/regen-ledger/types/module/client/cli"
	restmodule "github.com/regen-network/regen-ledger/types/module/client/grpc_gateway"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	bond "github.com/regen-network/regen-ledger/v2/x/bond"
	"github.com/regen-network/regen-ledger/v2/x/bond/client"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
	"log"
	"math/rand"

	"github.com/regen-network/regen-ledger/v2/x/bond/server"
)

type Module struct {
	log           log.Logger
	paramSpace    paramtypes.Subspace
	accountKeeper bond.AccountKeeper
	bankKeeper    bond.BankKeeper
	keeper        bond.Keeper
}

// NewModule returns a new Module object.
func NewModule(paramSpace paramtypes.Subspace, accountKeeper bond.AccountKeeper, bankKeeper bond.BankKeeper) Module {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(bond.ParamKeyTable())
	}

	return Module{
		paramSpace:    paramSpace,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
	}
}

var _ module.AppModuleBasic = Module{}
var _ servermodule.Module = Module{}
var _ restmodule.Module = Module{}
var _ climodule.Module = Module{}

// TODO: Add tests
// var _ module.AppModuleSimulation = Module{}

func (a Module) Name() string {
	return bond.ModuleName
}

func (a Module) RegisterInterfaces(registry types.InterfaceRegistry) {
	bond.RegisterInterfaces(registry)
}

func (a Module) RegisterServices(configurator servermodule.Configurator) {
	a.keeper = server.RegisterServices(configurator, a.paramSpace, a.accountKeeper, a.bankKeeper)
}

//nolint:errcheck
func (a Module) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	bond.RegisterQueryHandlerClient(context.Background(), mux, bond.NewQueryClient(clientCtx))
}

func (a Module) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(bond.DefaultGenesisState())
}

func (a Module) ValidateGenesis(cdc codec.JSONCodec, _ sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	var data bond.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", bond.ModuleName, err)
	}

	return data.Validate()
}

func (a Module) GetQueryCmd() *cobra.Command {
	return client.GetQueryCmd()
}

func (a Module) GetTxCmd() *cobra.Command {
	return client.GetTxCmd(a.Name())
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (Module) ConsensusVersion() uint64 { return 1 }

/**** DEPRECATED ****/
func (a Module) RegisterRESTRoutes(sdkclient.Context, *mux.Router) {}
func (a Module) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	bond.RegisterLegacyAminoCodec(cdc)
}

//
// AppModuleSimulation functions
//

// GenerateGenesisState creates a randomized GenesisState of the bond module.
func (Module) GenerateGenesisState(simState *module.SimulationState) {
	// simulation.RandomizedGenState(simState)
}

// ProposalContents returns all the bond content functions used to
// simulate proposals.
func (Module) ProposalContents(simState module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized bond param changes for the simulator.
func (Module) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	return nil
}

// RegisterStoreDecoder registers a decoder for bond module's types
func (Module) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
}

// WeightedOperations returns all the bond module operations with their respective weights.
// NOTE: This is no longer needed for the modules which uses ADR-33, bond module `WeightedOperations`
// registered in the `x/bond/server` package.
func (Module) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return nil
}

// BeginBlock checks if there are any expired sell or buy orders and removes them from state.
func (a Module) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
	err := bond.BeginBlocker(ctx, a.keeper)
	if err != nil {
		panic(err)
	}
}
