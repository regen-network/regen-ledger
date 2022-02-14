package client

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/golang/protobuf/proto"
	"github.com/regen-network/regen-ledger/v2/x/bond"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Bond queries under a subcommand
	cmd := &cobra.Command{
		Use:                        bond.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", bond.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(QueryBondsCmd())

	return cmd
}

// QueryBondsCmd returns a query command that retrieves bonds.
func QueryBondsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bonds",
		Short: "List all bonds",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.Bonds(cmd.Context(), &bond.QueryBondsRequest{
				Pagination: pagination,
			})
			return print(ctx, res, err)
		},
	}
	flags.AddPaginationFlagsToCmd(cmd, "bonds")
	return qflags(cmd)
}

func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Shows the parameters of the module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := bond.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &bond.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func qflags(cmd *cobra.Command) *cobra.Command {
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func mkQueryClient(cmd *cobra.Command) (bond.QueryClient, sdkclient.Context, error) {
	ctx, err := sdkclient.GetClientQueryContext(cmd)
	if err != nil {
		return nil, sdkclient.Context{}, err
	}
	return bond.NewQueryClient(ctx), ctx, err
}

// prints a query client response
func print(cctx sdkclient.Context, res proto.Message, err error) error {
	if err != nil {
		return err
	}
	return cctx.PrintProto(res)
}
