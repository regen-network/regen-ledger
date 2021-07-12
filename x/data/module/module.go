package module

import (
	"context"
	"encoding/json"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	restmodule "github.com/regen-network/regen-ledger/types/module/client/grpc_gateway"
	"github.com/spf13/cobra"

	climodule "github.com/regen-network/regen-ledger/types/module/client/cli"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/data/client"
	"github.com/regen-network/regen-ledger/x/data/server"
)

type Module struct{}

var _ module.AppModuleBasic = Module{}
var _ servermodule.Module = Module{}
var _ restmodule.Module = Module{}
var _ climodule.Module = Module{}

func (a Module) Name() string {
	return "data"
}

func (a Module) RegisterInterfaces(registry types.InterfaceRegistry) {
	data.RegisterTypes(registry)
}

func (a Module) RegisterServices(configurator servermodule.Configurator) {
	server.RegisterServices(configurator)
}

//nolint:errcheck
func (a Module) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	data.RegisterQueryHandlerClient(context.Background(), mux, data.NewQueryClient(clientCtx))
}

func (a Module) DefaultGenesis(codec.JSONCodec) json.RawMessage { return nil }

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
