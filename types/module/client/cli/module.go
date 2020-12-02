package cli

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/regen-network/regen-ledger/types/module"
	"github.com/spf13/cobra"
)

type Module interface {
	module.ModuleBase

	DefaultGenesis(codec.JSONMarshaler) json.RawMessage
	ValidateGenesis(codec.JSONMarshaler, client.TxEncodingConfig, json.RawMessage) error
	GetTxCmd() *cobra.Command
	GetQueryCmd() *cobra.Command
}

func AddTxCommands(rootTxCmd *cobra.Command, modules module.Modules) {
	for _, m := range modules {
		if cliMod, ok := m.(Module); ok {
			if cmd := cliMod.GetTxCmd(); cmd != nil {
				rootTxCmd.AddCommand(cmd)
			}
		}
	}
}

func AddQueryCommands(rootTxCmd *cobra.Command, modules module.Modules) {
	for _, m := range modules {
		if cliMod, ok := m.(Module); ok {
			if cmd := cliMod.GetQueryCmd(); cmd != nil {
				rootTxCmd.AddCommand(cmd)
			}
		}
	}
}
