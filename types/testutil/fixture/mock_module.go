package fixture

import (
	"encoding/json"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/proto/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
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

func (m MockModule) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

func (m MockModule) RegisterInterfaces(_ types.InterfaceRegistry) {}

func (m MockModule) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	return nil
}

func (m MockModule) ValidateGenesis(_ codec.JSONCodec, _ client.TxEncodingConfig, _ json.RawMessage) error {
	return nil
}

func (m MockModule) RegisterGRPCGatewayRoutes(_ client.Context, _ *runtime.ServeMux) {}

func (m MockModule) GetTxCmd() *cobra.Command {
	return nil
}

func (m MockModule) GetQueryCmd() *cobra.Command {
	return nil
}

func (m MockModule) InitGenesis(_ sdk.Context, _ codec.JSONCodec, _ json.RawMessage) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{{PubKey: crypto.PublicKey{}, Power: 40}}
}

func (m MockModule) ExportGenesis(_ sdk.Context, _ codec.JSONCodec) json.RawMessage {
	return nil
}

func (m MockModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (m MockModule) Route() sdk.Route {
	return sdk.Route{}
}

func (m MockModule) QuerierRoute() string {
	return ""
}

func (m MockModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier {
	return nil
}

func (m MockModule) RegisterServices(_ module.Configurator) {}

func (m MockModule) ConsensusVersion() uint64 {
	return 1
}
