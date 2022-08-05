package wasm

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/CosmWasm/wasmd/x/wasm/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string {
	return ModuleName
}

func (a AppModuleBasic) RegisterLegacyAminoCodec(*codec.LegacyAmino) {}

func (a AppModuleBasic) RegisterInterfaces(cdctypes.InterfaceRegistry) {}

func (a AppModuleBasic) DefaultGenesis(codec.JSONCodec) json.RawMessage {
	return []byte("[]")
}

func (a AppModuleBasic) ValidateGenesis(codec.JSONCodec, client.TxEncodingConfig, json.RawMessage) error {
	return nil
}

func (a AppModuleBasic) RegisterRESTRoutes(client.Context, *mux.Router) {}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux) {}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return &cobra.Command{}
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	return &cobra.Command{}
}

type AppModule struct{}

func (a AppModule) Name() string {
	return ModuleName
}

func (a AppModule) RegisterLegacyAminoCodec(*codec.LegacyAmino) {}

func (a AppModule) RegisterInterfaces(cdctypes.InterfaceRegistry) {}

func (a AppModule) DefaultGenesis(codec.JSONCodec) json.RawMessage {
	return []byte("[]")
}

func (a AppModule) ValidateGenesis(codec.JSONCodec, client.TxEncodingConfig, json.RawMessage) error {
	return nil
}

func (a AppModule) RegisterRESTRoutes(client.Context, *mux.Router) {}

func (a AppModule) RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux) {}

func (a AppModule) GetTxCmd() *cobra.Command {
	return &cobra.Command{}
}

func (a AppModule) GetQueryCmd() *cobra.Command {
	return &cobra.Command{}
}

func (a AppModule) InitGenesis(sdk.Context, codec.JSONCodec, json.RawMessage) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func (a AppModule) ExportGenesis(sdk.Context, codec.JSONCodec) json.RawMessage {
	return []byte("[]")
}

func (a AppModule) RegisterInvariants(sdk.InvariantRegistry) {}

func (a AppModule) Route() sdk.Route {
	return sdk.Route{}
}

func (a AppModule) QuerierRoute() string {
	return ""
}

func (a AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	return nil
}

func (a AppModule) RegisterServices(module.Configurator) {}

func (a AppModule) ConsensusVersion() uint64 {
	return 0
}

func (a AppModule) BeginBlock(sdk.Context, abci.RequestBeginBlock) {}

func (a AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func NewAppModule(_ codec.Codec, _ *Keeper, _ types.StakingKeeper) AppModule {
	return AppModule{}
}

func AddModuleInitFlags(_ *cobra.Command) {}

func ReadWasmConfig(_ servertypes.AppOptions) (types.WasmConfig, error) {
	return types.WasmConfig{}, nil
}
