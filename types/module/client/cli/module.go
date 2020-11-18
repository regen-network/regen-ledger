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

func AddQueryCommands(rootQueryCmd *cobra.Command, moduleMap module.ModuleMap) {
	for _, m := range moduleMap {
		cliMod, ok := m.(Module)
		if !ok {
			continue
		}

		if cmd := cliMod.GetQueryCmd(); cmd != nil {
			rootQueryCmd.AddCommand(cmd)
		}
	}
}

func AddTxCommands(rootTxCmd *cobra.Command, moduleMap module.ModuleMap) {
	for _, m := range moduleMap {
		cliMod, ok := m.(Module)
		if !ok {
			continue
		}

		if cmd := cliMod.GetTxCmd(); cmd != nil {
			rootTxCmd.AddCommand(cmd)
		}
	}
}
