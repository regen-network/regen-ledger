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

type AppModule struct {
	storeKey sdk.StoreKey
}

var _ module.AppModule = AppModule{}

func (a AppModule) Name() string { return data.ModuleName }

func (a AppModule) RegisterLegacyAminoCodec(*codec.LegacyAmino) {}

func (a AppModule) RegisterInterfaces(codectypes.InterfaceRegistry) {}

func (a AppModule) DefaultGenesis(codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (a AppModule) ValidateGenesis(codec.JSONMarshaler, sdkclient.TxEncodingConfig, json.RawMessage) error {
	return nil
}

func (a AppModule) RegisterRESTRoutes(sdkclient.Context, *mux.Router) {}

func (a AppModule) RegisterGRPCGatewayRoutes(sdkclient.Context, *runtime.ServeMux) {}

func (a AppModule) GetTxCmd() *cobra.Command {
	return client.TxCmd()
}

func (a AppModule) GetQueryCmd() *cobra.Command {
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
	impl := server.NewServer(a.storeKey)
	data.RegisterMsgServer(configurator.MsgServer(), impl)
	data.RegisterQueryServer(configurator.QueryServer(), impl)
}

func (a AppModule) BeginBlock(sdk.Context, abci.RequestBeginBlock) {
}

func (a AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}
