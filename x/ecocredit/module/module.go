package module

import (
	"encoding/json"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
)

type AppModule struct {
	Key sdk.StoreKey
	Srv server.Server
}

var _ module.AppModule = AppModule{}

func NewAppModule() AppModule {
	key := sdk.NewKVStoreKey(ecocredit.ModuleName)
	return AppModule{key, server.NewServer(key)}
}

func (a AppModule) Name() string { return a.Key.Name() }

func (a AppModule) DefaultGenesis(codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (a AppModule) ValidateGenesis(codec.JSONMarshaler, sdkclient.TxEncodingConfig, json.RawMessage) error {
	return nil
}
func (a AppModule) InitGenesis(sdk.Context, codec.JSONMarshaler, json.RawMessage) []abci.ValidatorUpdate {
	return nil
}

func (a AppModule) ExportGenesis(sdk.Context, codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (a AppModule) GetTxCmd() *cobra.Command {
	return nil
}

func (a AppModule) GetQueryCmd() *cobra.Command {
	return nil
}

func (a AppModule) RegisterGRPCGatewayRoutes(sdkclient.Context, *runtime.ServeMux) {}

func (a AppModule) RegisterInvariants(sdk.InvariantRegistry) {}

func (a AppModule) RegisterServices(cfg module.Configurator) {
	ecocredit.RegisterMsgServer(cfg.MsgServer(), a.Srv)
	ecocredit.RegisterQueryServer(cfg.QueryServer(), a.Srv)
}

func (a AppModule) BeginBlock(sdk.Context, abci.RequestBeginBlock) {}

func (a AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}

/**** DEPRECATED ****/

// Should we deprecate it? RegisterServices already registers everything for the module.
func (a AppModule) RegisterInterfaces(codectypes.InterfaceRegistry) {}

func (a AppModule) Route() sdk.Route {
	return sdk.Route{}
}

func (a AppModule) QuerierRoute() string { return a.Name() }

func (a AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	return nil
}

func (a AppModule) RegisterRESTRoutes(sdkclient.Context, *mux.Router) {}

func (a AppModule) RegisterLegacyAminoCodec(*codec.LegacyAmino) {}
