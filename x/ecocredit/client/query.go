package client

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// QueryCmd returns the parent command for all x/ecocredit query commands.
func QueryCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		SuggestionsMinimumDistance: 2,
		DisableFlagParsing:         true,

		Args:  cobra.ExactArgs(1),
		Use:   name,
		Short: "Query commands for the ecocredit module",
		RunE:  client.ValidateCmd,
	}
	cmd.AddCommand(
		QueryClassesCmd(),
		QueryClassInfoCmd(),
		QueryBatchesCmd(),
		QueryBatchInfoCmd(),
		QueryBalanceCmd(),
		QuerySupplyCmd(),
		QueryCreditTypesCmd(),
		QueryParams(),
	)
	return cmd
}

func qflags(cmd *cobra.Command) *cobra.Command {
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// QueryClassesCmd returns a query command that lists all credit classes.
func QueryClassesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "classes",
		Short: "List all credit classes with pagination flags",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.Classes(cmd.Context(), &ecocredit.QueryClassesRequest{
				Pagination: pagination,
			})
			return print(ctx, res, err)
		},
	}
	flags.AddPaginationFlagsToCmd(cmd, "classes")
	return qflags(cmd)
}

// QueryClassInfoCmd returns a query command that retrieves information for a
// given credit class.
func QueryClassInfoCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "class-info [class_id]",
		Short: "Retrieve credit class info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.ClassInfo(cmd.Context(), &ecocredit.QueryClassInfoRequest{
				ClassId: args[0],
			})
			return print(ctx, res, err)
		},
	})
}

// QueryBatchesCmd returns a query command that retrieves credit batches for a
// given credit class.
func QueryBatchesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batches [class_id]",
		Short: "List all credit batches in the given class with pagination flags",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.Batches(cmd.Context(), &ecocredit.QueryBatchesRequest{
				ClassId:    args[0],
				Pagination: pagination,
			})
			return print(ctx, res, err)
		},
	}
	flags.AddPaginationFlagsToCmd(cmd, "batches")
	return qflags(cmd)
}

// QueryBatchInfoCmd returns a query command that retrieves information for a
// given credit batch.
func QueryBatchInfoCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "batch-info [batch_denom]",
		Short: "Retrieve the credit issuance batch info",
		Long:  "Retrieve the credit issuance batch info based on the bach_denom (ID)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.BatchInfo(cmd.Context(), &ecocredit.QueryBatchInfoRequest{
				BatchDenom: args[0],
			})
			return print(ctx, res, err)
		},
	})
}

// QueryBalanceCmd returns a query command that retrieves the tradable and
// retired balances for a given credit batch and account address.
func QueryBalanceCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "balance [batch_denom] [account]",
		Short: "Retrieve the tradable and retired balances of the credit batch",
		Long:  "Retrieve the tradable and retired balances of the credit batch for a given account address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.Balance(cmd.Context(), &ecocredit.QueryBalanceRequest{
				BatchDenom: args[0], Account: args[1],
			})
			return print(ctx, res, err)
		},
	})
}

// QuerySupplyCmd returns a query command that retrieves the tradable and
// retired supply of credits for a given credit batch.
func QuerySupplyCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "supply [batch_denom]",
		Short: "Retrieve the tradable and retired supply of the credit batch",
		Long:  "Retrieve the tradable and retired supply of the credit batch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.Supply(cmd.Context(), &ecocredit.QuerySupplyRequest{
				BatchDenom: args[0],
			})
			return print(ctx, res, err)
		},
	})
}

// QueryCreditTypesCmd returns a query command that retrieves the list of
// approved credit types.
func QueryCreditTypesCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "types",
		Short: "Retrieve the list of credit types",
		Long:  "Retrieve the list of credit types that contains the type name, measurement unit and precision",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.CreditTypes(cmd.Context(), &ecocredit.QueryCreditTypesRequest{})
			return print(ctx, res, err)
		},
	})
}

// QueryParams returns ecocredit module parameters.
func QueryParams() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "params",
		Short: "Query the current ecocredit module parameters",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the current ecocredit module parameters
			
Examples:
$%s query %s params
$%s q %s params
			`, version.AppName, ecocredit.ModuleName, version.AppName, ecocredit.ModuleName),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.Params(cmd.Context(), &ecocredit.QueryParamsRequest{})
			return print(ctx, res, err)
		},
	})
}
