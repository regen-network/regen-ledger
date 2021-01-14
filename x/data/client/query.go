package client

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	gocid "github.com/ipfs/go-cid"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/spf13/cobra"
)

// QueryCmd returns the parent command for all x/data CLI query commands
func QueryCmd(name string) *cobra.Command {
	queryByCidCmd := QueryByCidCmd()

	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Use:   fmt.Sprintf("%s [cid]", name),
		Short: "Querying commands for the data module",
		Long: strings.TrimSpace(`Querying commands for the data module.
If a CID is passed as first argument, then this command will query timestamp, signers and content (if available) for the given CID. Otherwise, this command will run the given subcommand.

Example (the two following commands are equivalent):
$ regen query data bafzbeigai3eoy2ccc7ybwjfz5r3rdxqrinwi4rwytly24tdbh6yk7zslrm
$ regen query data by-cid bafzbeigai3eoy2ccc7ybwjfz5r3rdxqrinwi4rwytly24tdbh6yk7zslrm`),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE: func(cmd *cobra.Command, args []string) error {
			// If 1st arg is NOT a CID, parse subcommands as usual.
			_, err := gocid.Decode(args[0])
			if err != nil {
				return client.ValidateCmd(cmd, args)
			}

			// Or else, we call QueryByCidCmd.
			return queryByCidCmd.RunE(cmd, args)
		},
	}

	cmd.AddCommand(
		queryByCidCmd,
	)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryByCidCmd creates a CLI command for Query/Data.
func QueryByCidCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "by-cid [cid]",
		Short: "Query for CID timestamp, signers and content (if available)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			cid, err := gocid.Decode(args[0])
			if err != nil {
				return err
			}

			queryClient := data.NewQueryClient(clientCtx)

			res, err := queryClient.ByCid(cmd.Context(), &data.QueryByCidRequest{Cid: cid.Bytes()})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
