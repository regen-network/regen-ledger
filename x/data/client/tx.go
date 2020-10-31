package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/spf13/cobra"
)

// TxCmd returns a root CLI command handler for all x/bank transaction commands.
func TxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        data.ModuleName,
		Short:                      "Data transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand()

	return cmd
}
