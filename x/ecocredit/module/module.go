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
	"github.com/regen-network/regen-ledger/x/ecocredit/client"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
)

const (
	StoreKey = ecocredit.ModuleName
)

type AppModuleBasic struct{}

type AppModule struct {
	AppModuleBasic
	key sdk.StoreKey
}

var _ module.AppModule = AppModule{}
var _ module.AppModuleBasic = AppModuleBasic{}

func NewAppModule(key sdk.StoreKey) module.AppModule {
	return AppModule{key: key}
}

func (a AppModuleBasic) Name() string { return ecocredit.ModuleName }

func (a AppModuleBasic) DefaultGenesis(codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (a AppModuleBasic) ValidateGenesis(codec.JSONMarshaler, sdkclient.TxEncodingConfig, json.RawMessage) error {
	return nil
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	return client.QueryCmd()
}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return client.TxCmd()
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(sdkclient.Context, *runtime.ServeMux) {}

func (a AppModuleBasic) RegisterInterfaces(r codectypes.InterfaceRegistry) {
	ecocredit.RegisterTypes(r)
}

func (a AppModule) InitGenesis(sdk.Context, codec.JSONMarshaler, json.RawMessage) []abci.ValidatorUpdate {
	return nil
}

func (a AppModule) ExportGenesis(sdk.Context, codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (a AppModule) RegisterInvariants(sdk.InvariantRegistry) {}

func (a AppModule) RegisterServices(cfg module.Configurator) {
	server.RegisterServices(a.key, cfg)
}

func (a AppModule) BeginBlock(sdk.Context, abci.RequestBeginBlock) {}

func (a AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}

/**** DEPRECATED ****/

func (a AppModuleBasic) RegisterLegacyAminoCodec(*codec.LegacyAmino)       {}
func (a AppModuleBasic) RegisterRESTRoutes(sdkclient.Context, *mux.Router) {}

func (a AppModule) Route() sdk.Route     { return sdk.Route{} }
func (a AppModule) QuerierRoute() string { return a.Name() }

func (a AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier { return nil }
