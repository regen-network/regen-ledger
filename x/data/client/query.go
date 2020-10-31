package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	gocid "github.com/ipfs/go-cid"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/spf13/cobra"
)

// QueryCmd returns the parent command for all x/data CLI query commands
func QueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        data.ModuleName,
		Short:                      "Querying commands for the data module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		QueryDataCmd(),
	)

	return cmd
}

func QueryDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "data [cid]",
		Short: "Query for CID timestamp, signers and content (if available)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			cid, err := gocid.Decode(args[0])
			if err != nil {
				return err
			}

			queryClient := data.NewQueryClient(clientCtx)

			res, err := queryClient.Data(cmd.Context(), &data.QueryDataRequest{Cid: cid.Bytes()})
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
