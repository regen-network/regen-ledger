package module

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/data/client"
	"github.com/regen-network/regen-ledger/x/data/server"
)

const (
	StoreKey = data.ModuleName
)

type AppModuleBasic struct{}

type AppModule struct {
	AppModuleBasic
	storeKey sdk.StoreKey
}

func NewAppModule(storeKey sdk.StoreKey) *AppModule {
	return &AppModule{storeKey: storeKey}
}

var _ module.AppModule = AppModule{}
var _ module.AppModuleBasic = AppModuleBasic{}

func (a AppModuleBasic) Name() string { return data.ModuleName }

func (a AppModuleBasic) RegisterLegacyAminoCodec(*codec.LegacyAmino) {}

func (a AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	data.RegisterTypes(registry)
}

func (a AppModuleBasic) DefaultGenesis(codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (a AppModuleBasic) ValidateGenesis(codec.JSONMarshaler, sdkclient.TxEncodingConfig, json.RawMessage) error {
	return nil
}

func (a AppModuleBasic) RegisterRESTRoutes(sdkclient.Context, *mux.Router) {}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(sdkclient.Context, *runtime.ServeMux) {}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return client.TxCmd()
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	return client.QueryCmd()
}

func (a AppModule) InitGenesis(sdk.Context, codec.JSONMarshaler, json.RawMessage) []abci.ValidatorUpdate {
	return nil
}

func (a AppModule) ExportGenesis(sdk.Context, codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (a AppModule) RegisterInvariants(sdk.InvariantRegistry) {}

func (a AppModule) Route() sdk.Route { return sdk.Route{} }

func (a AppModule) QuerierRoute() string { return "" }

func (a AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	return nil
}

func (a AppModule) RegisterServices(configurator module.Configurator) {
	server.RegisterServices(a.storeKey, configurator)
}

func (a AppModule) BeginBlock(sdk.Context, abci.RequestBeginBlock) {
}

func (a AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}
