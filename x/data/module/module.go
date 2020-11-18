package module

import (
	"encoding/json"
	"github.com/regen-network/regen-ledger/types/module/client/cli"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/data"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/regen-network/regen-ledger/x/data/client"
	"github.com/regen-network/regen-ledger/x/data/server"
	"github.com/spf13/cobra"
)

const DefaultModuleName = "data"

type Module struct{}

var _ servermodule.Module = Module{}
var _ cli.Module = Module{}

func (a Module) RegisterTypes(registry codectypes.InterfaceRegistry) {
	data.RegisterTypes(registry)
}

func (a Module) RegisterServices(configurator servermodule.Configurator) {
	server.RegisterServices(configurator.ModuleKey(), configurator)
}

func (a Module) GetTxCmd() *cobra.Command {
	return client.TxCmd()
}

func (a Module) GetQueryCmd() *cobra.Command {
	return client.QueryCmd()
}

func (a Module) DefaultGenesis(codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (a Module) ValidateGenesis(codec.JSONMarshaler, sdkclient.TxEncodingConfig, json.RawMessage) error {
	return nil
}
