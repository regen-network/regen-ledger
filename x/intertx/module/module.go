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

// NewModule returns a new Module.
func NewModule(k keeper.Keeper) AppModule {
	return AppModule{k}
}

// Name implements AppModule/Name.
func (a AppModule) Name() string {
	return intertx.ModuleName
}

// RegisterLegacyAminoCodec implements AppModule/RegisterLegacyAminoCodec.
func (a AppModule) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	v1.RegisterLegacyAminoCodec(amino)
}

// RegisterInterfaces implements AppModule/RegisterTypes.
func (a AppModule) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	v1.RegisterTypes(registry)
}

// DefaultGenesis implements AppModule/DefaultGenesis.
func (a AppModule) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	return nil
}

// ValidateGenesis implements AppModule/ValidateGenesis.
func (a AppModule) ValidateGenesis(_ codec.JSONCodec, _ client.TxEncodingConfig, _ json.RawMessage) error {
	return nil
}

// RegisterGRPCGatewayRoutes implements AppModule/RegisterGRPCGatewayRoutes.
func (a AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	err := v1.RegisterQueryHandlerClient(context.Background(), mux, v1.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
}

// GetTxCmd implements AppModule/GetTxCmd.
func (a AppModule) GetTxCmd() *cobra.Command {
	return intertxClient.GetTxCmd()
}

// GetQueryCmd implements AppModule/GetQueryCmd.
func (a AppModule) GetQueryCmd() *cobra.Command {
	return intertxClient.GetQueryCmd()
}

// InitGenesis implements AppModule/InitGenesis.
func (a AppModule) InitGenesis(context sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) []types.ValidatorUpdate {
	return []types.ValidatorUpdate{}
}

// ExportGenesis implements AppModule/ExportGenesis.
func (a AppModule) ExportGenesis(context sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	return nil
}

// RegisterInvariants implements AppModule/RegisterInvariants.
func (a AppModule) RegisterInvariants(registry sdk.InvariantRegistry) {}

// Route implements AppModule/Route.
func (a AppModule) Route() sdk.Route {
	return sdk.NewRoute(intertx.RouterKey, nil)
}

// QuerierRoute implements AppModule/QuerierRoute.
func (a AppModule) QuerierRoute() string {
	return intertx.QuerierRoute
}

// LegacyQuerierHandler implements AppModule/LegacyQuerierHandler.
func (a AppModule) LegacyQuerierHandler(amino *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices implements AppModule/RegisterServices.
func (a AppModule) RegisterServices(cfg module.Configurator) {
	v1.RegisterMsgServer(cfg.MsgServer(), a.keeper)
	v1.RegisterQueryServer(cfg.QueryServer(), a.keeper)
}

// ConsensusVersion is the module consensus version
func (a AppModule) ConsensusVersion() uint64 {
	return 1
}
