package server

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/proto/tendermint/crypto"
)

/*
* This mock module is needed to bypass the requirement of validator updates in the SDK module manager.
* Since there are no validator updates when testing, we inject one via this mock module.
 */

var _ module.AppModule = MockModule{}

type MockModule struct{}

func (m MockModule) Name() string {
	return "mocker"
}

func (m MockModule) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	return
}

func (m MockModule) RegisterInterfaces(registry types.InterfaceRegistry) {
	return
}

func (m MockModule) DefaultGenesis(codec codec.JSONCodec) json.RawMessage {
	return nil
}

func (m MockModule) ValidateGenesis(codec codec.JSONCodec, config client.TxEncodingConfig, message json.RawMessage) error {
	return nil
}

func (m MockModule) RegisterGRPCGatewayRoutes(context client.Context, mux *runtime.ServeMux) {
	return
}

func (m MockModule) GetTxCmd() *cobra.Command {
	return nil
}

func (m MockModule) GetQueryCmd() *cobra.Command {
	return nil
}

func (m MockModule) InitGenesis(context sdk.Context, codec codec.JSONCodec, message json.RawMessage) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{{crypto.PublicKey{}, 40}}
}

func (m MockModule) ExportGenesis(context sdk.Context, codec codec.JSONCodec) json.RawMessage {
	return nil
}

func (m MockModule) RegisterInvariants(registry sdk.InvariantRegistry) {

}

func (m MockModule) Route() sdk.Route {
	return sdk.Route{}
}

func (m MockModule) QuerierRoute() string {
	return ""
}

func (m MockModule) LegacyQuerierHandler(amino *codec.LegacyAmino) sdk.Querier {
	return nil
}

func (m MockModule) RegisterServices(configurator module.Configurator) {
}

func (m MockModule) ConsensusVersion() uint64 {
	return 1
}
