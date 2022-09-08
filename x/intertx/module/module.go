package module

import (
	"context"
	"encoding/json"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/x/intertx"
	intertxClient "github.com/regen-network/regen-ledger/x/intertx/client"
	"github.com/regen-network/regen-ledger/x/intertx/keeper"
	v1 "github.com/regen-network/regen-ledger/x/intertx/types/v1"
)

var (
	_ module.AppModule = AppModule{}
)

// AppModule implements the AppModule interface for the capability module.
type AppModule struct {
	keeper keeper.Keeper
}

func NewModule(k keeper.Keeper) AppModule {
	return AppModule{k}
}

func (a AppModule) Name() string {
	return intertx.ModuleName
}

func (a AppModule) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	v1.RegisterCodec(amino)
}

func (a AppModule) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	v1.RegisterInterfaces(registry)
}

func (a AppModule) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	return nil
}

func (a AppModule) ValidateGenesis(_ codec.JSONCodec, _ client.TxEncodingConfig, _ json.RawMessage) error {
	return nil // TODO(Tyler): validate here?
}

func (a AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	err := v1.RegisterQueryHandlerClient(context.Background(), mux, v1.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
}

func (a AppModule) GetTxCmd() *cobra.Command {
	return intertxClient.GetTxCmd()
}

func (a AppModule) GetQueryCmd() *cobra.Command {
	return intertxClient.GetQueryCmd()
}

func (a AppModule) InitGenesis(context sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) []types.ValidatorUpdate {
	return []types.ValidatorUpdate{} // TODO(Tyler): update?
}

func (a AppModule) ExportGenesis(context sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	return nil // TODO(Tyler): impl?
}

func (a AppModule) RegisterInvariants(registry sdk.InvariantRegistry) {}

func (a AppModule) Route() sdk.Route {
	return sdk.NewRoute(intertx.RouterKey, nil)
}

func (a AppModule) QuerierRoute() string {
	return intertx.QuerierRoute
}

func (a AppModule) LegacyQuerierHandler(amino *codec.LegacyAmino) sdk.Querier {
	return nil
}

func (a AppModule) RegisterServices(cfg module.Configurator) {
	v1.RegisterMsgServer(cfg.MsgServer(), a.keeper)
	v1.RegisterQueryServer(cfg.QueryServer(), a.keeper)
}

func (a AppModule) ConsensusVersion() uint64 {
	return 1
}
