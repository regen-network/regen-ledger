package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/spf13/cobra"
)

// TxCmd returns a root CLI command handler for all x/data transaction commands.
func TxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        data.ModuleName,
		Short:                      "Data transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		MsgAnchorDataCmd(),
	)

	return cmd
}

func MsgAnchorDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "send [from_key_or_address] [to_address] [amount]",
		Short: `Send funds from one account to another. Note, the'--from' flag is
ignored as it is implied from [from_key_or_address].`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			// TODO

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags())
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
