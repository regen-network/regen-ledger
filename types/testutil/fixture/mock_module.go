package fixture

import (
	"encoding/json"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/proto/tendermint/crypto"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
)

/*
* This mock module is needed to bypass the requirement of validator updates in the SDK module manager.
* Since there are no validator updates when testing, we inject one via this mock module.
 */

var _ module.AppModule = MockModule{}

type MockModule struct{}

// IsAppModule implements module.AppModule.
func (m MockModule) IsAppModule() {
	panic("unimplemented")
}

// IsOnePerModuleType implements module.AppModule.
func (m MockModule) IsOnePerModuleType() {
	panic("unimplemented")
}

// Name implements module.AppModule.
func (m MockModule) Name() string {
	return "mocker"
}

// RegisterLegacyAminoCodec implements module.AppModule.
func (m MockModule) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

// RegisterInterfaces implements module.AppModule.
func (m MockModule) RegisterInterfaces(_ types.InterfaceRegistry) {}

// DefaultGenesis implements module.AppModule.
func (m MockModule) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	return nil
}

// ValidateGenesis implements module.AppModule.
func (m MockModule) ValidateGenesis(_ codec.JSONCodec, _ client.TxEncodingConfig, _ json.RawMessage) error {
	return nil
}

// RegisterGRPCGatewayRoutes implements module.AppModule.
func (m MockModule) RegisterGRPCGatewayRoutes(_ client.Context, _ *runtime.ServeMux) {}

// GetTxCmd implements module.AppModule.
func (m MockModule) GetTxCmd() *cobra.Command {
	return nil
}

// GetQueryCmd implements module.AppModule.
func (m MockModule) GetQueryCmd() *cobra.Command {
	return nil
}

// InitGenesis implements module.AppModule.
func (m MockModule) InitGenesis(_ sdk.Context, _ codec.JSONCodec, _ json.RawMessage) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{{PubKey: crypto.PublicKey{}, Power: 40}}
}

// ExportGenesis implements module.AppModule.
func (m MockModule) ExportGenesis(_ sdk.Context, _ codec.JSONCodec) json.RawMessage {
	return nil
}

// RegisterInvariants implements module.AppModule.
func (m MockModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// RegisterServices implements module.AppModule.
func (m MockModule) RegisterServices(_ module.Configurator) {}

// ConsensusVersion implements module.AppModule.
func (m MockModule) ConsensusVersion() uint64 {
	return 1
}
