package module

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	restmodule "github.com/regen-network/regen-ledger/types/module/client/grpc_gateway"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	climodule "github.com/regen-network/regen-ledger/types/module/client/cli"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/client"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation"
)

type Module struct {
	paramSpace    paramtypes.Subspace
	accountKeeper ecocredit.AccountKeeper
	bankKeeper    ecocredit.BankKeeper
	keeper        ecocredit.Keeper
}

// NewModule returns a new Module object.
func NewModule(paramSpace paramtypes.Subspace, accountKeeper ecocredit.AccountKeeper, bankKeeper ecocredit.BankKeeper) *Module {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(ecocredit.ParamKeyTable())
	}

	return &Module{
		paramSpace:    paramSpace,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
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
}

func (a *Module) RegisterServices(configurator servermodule.Configurator) {
	a.keeper = server.RegisterServices(configurator, a.paramSpace, a.accountKeeper, a.bankKeeper)
}

//nolint:errcheck
func (a Module) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	ecocredit.RegisterQueryHandlerClient(context.Background(), mux, ecocredit.NewQueryClient(clientCtx))
}

func (a Module) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(ecocredit.DefaultGenesisState())
}

func (a Module) ValidateGenesis(cdc codec.JSONCodec, _ sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	var data ecocredit.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
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
func (Module) ConsensusVersion() uint64 { return 1 }

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
