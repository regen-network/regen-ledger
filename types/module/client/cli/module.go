package cli

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/regen-network/regen-ledger/types/module"
)

// Module is an interface that modules should implement to integrate with the CLI framework.
type Module interface {
	module.TypeModule

	DefaultGenesis(codec.JSONCodec) json.RawMessage
	ValidateGenesis(codec.JSONCodec, client.TxEncodingConfig, json.RawMessage) error
	GetTxCmd() *cobra.Command
	GetQueryCmd() *cobra.Command
}
