package module

import (
	"context"
	"encoding/json"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	restmodule "github.com/regen-network/regen-ledger/types/module/client/grpc_gateway"
	"github.com/spf13/cobra"

	climodule "github.com/regen-network/regen-ledger/types/module/client/cli"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/client"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
)

type Module struct {
	paramSpace paramtypes.Subspace
	bankKeeper ecocredit.BankKeeper
}

func NewModule(paramSpace paramtypes.Subspace, bankKeeper ecocredit.BankKeeper) Module {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(ecocredit.ParamKeyTable())
	}

	return Module{
		paramSpace: paramSpace,
		bankKeeper: bankKeeper,
	}
}

var _ module.AppModuleBasic = Module{}
var _ servermodule.Module = Module{}
var _ restmodule.Module = Module{}
var _ climodule.Module = Module{}

func (a Module) Name() string {
	return ecocredit.ModuleName
}

func (a Module) RegisterInterfaces(registry types.InterfaceRegistry) {
	ecocredit.RegisterTypes(registry)
}

func (a Module) RegisterServices(configurator servermodule.Configurator) {
	server.RegisterServices(configurator, a.paramSpace, a.bankKeeper)
}

//nolint:errcheck
func (a Module) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	ecocredit.RegisterQueryHandlerClient(context.Background(), mux, ecocredit.NewQueryClient(clientCtx))
}

func (a Module) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(ecocredit.DefaultGenesisState())
}

func (a Module) ValidateGenesis(codec.JSONCodec, sdkclient.TxEncodingConfig, json.RawMessage) error {
	return nil
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
func (a Module) RegisterLegacyAminoCodec(*codec.LegacyAmino)       {}
