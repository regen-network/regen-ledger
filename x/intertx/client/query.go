package client

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/regen-network/regen-ledger/x/intertx"
	intertxv1 "github.com/regen-network/regen-ledger/x/intertx/types/v1"
)

// GetQueryCmd creates and returns the intertx query command
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        intertx.ModuleName,
		Short:                      "Querying commands for the intertx module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(getInterchainAccountCmd())

	return cmd
}

func getInterchainAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ica [owner-account] [connection-id]",
		Short:   "query the interchain account address",
		Long:    "query the interchain account address associated with the owner address and connection id.",
		Example: "regen q intertx ica regen1drn830y2l24pne08t7k7p7z6zms3x8p8zc3u0h channel-5",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := intertxv1.NewQueryClient(clientCtx)
			req := intertxv1.QueryInterchainAccountRequest{
				Owner:        args[0],
				ConnectionId: args[1],
			}
			res, err := queryClient.InterchainAccount(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
