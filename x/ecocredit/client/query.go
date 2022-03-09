package client

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	basketcli "github.com/regen-network/regen-ledger/x/ecocredit/client/basket"
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
		QueryProjectsCmd(),
		QueryProjectInfoCmd(),
		QueryParamsCmd(),
		QuerySellOrderCmd(),
		QuerySellOrdersCmd(),
		QuerySellOrdersByAddressCmd(),
		QuerySellOrdersByBatchDenomCmd(),
		QueryBuyOrderCmd(),
		QueryBuyOrdersCmd(),
		QueryBuyOrdersByAddressCmd(),
		QueryAllowedAskDenomsCmd(),
		basketcli.QueryBasketCmd(),
		basketcli.QueryBasketsCmd(),
		basketcli.QueryBasketBalanceCmd(),
		basketcli.QueryBasketBalancesCmd(),
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

// QueryProjectsCmd returns a query command that retrieves projects.
func QueryProjectsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects [class_id]",
		Short: "List all projects in the given class with pagination flags",
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

			res, err := c.Projects(cmd.Context(), &ecocredit.QueryProjectsRequest{
				ClassId:    args[0],
				Pagination: pagination,
			})
			return print(ctx, res, err)
		},
	}
	flags.AddPaginationFlagsToCmd(cmd, "projects")
	return qflags(cmd)
}

// QueryProjectInfoCmd returns a query command that retrieves project information.
func QueryProjectInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project-info [project_id]",
		Short: "Retrive project info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.ProjectInfo(cmd.Context(), &ecocredit.QueryProjectInfoRequest{
				ProjectId: args[0],
			})
			return print(ctx, res, err)
		},
	}

	return qflags(cmd)
}

// QueryBatchesCmd returns a query command that retrieves credit batches for a
// given project.
func QueryBatchesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batches [project_id]",
		Short: "List all credit batches in the given project with pagination flags",
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
				ProjectId:  args[0],
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

// QueryParamsCmd returns ecocredit module parameters.
func QueryParamsCmd() *cobra.Command {
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

// QuerySellOrderCmd returns a query command that retrieves information for a given sell order.
func QuerySellOrderCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "sell-order [sell_order_id]",
		Short: "Retrieve information for a given sell order",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			sellOrderId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return ecocredit.ErrInvalidSellOrder.Wrap(err.Error())
			}
			res, err := c.SellOrder(cmd.Context(), &ecocredit.QuerySellOrderRequest{
				SellOrderId: sellOrderId,
			})
			return print(ctx, res, err)
		},
	})
}

// QuerySellOrdersCmd returns a query command that retrieves all sell orders with pagination.
func QuerySellOrdersCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "sell-orders",
		Short: "List all sell orders with pagination",
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
			res, err := c.SellOrders(cmd.Context(), &ecocredit.QuerySellOrdersRequest{
				Pagination: pagination,
			})
			return print(ctx, res, err)
		},
	})
}

// QuerySellOrdersByAddressCmd returns a query command that retrieves all sell orders by address with pagination.
func QuerySellOrdersByAddressCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "sell-orders-by-address [address]",
		Short: "List all sell orders by owner address with pagination",
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
			res, err := c.SellOrdersByAddress(cmd.Context(), &ecocredit.QuerySellOrdersByAddressRequest{
				Address:    args[0],
				Pagination: pagination,
			})
			return print(ctx, res, err)
		},
	})
}

// QuerySellOrdersByBatchDenomCmd returns a query command that retrieves all sell orders by batch denom with pagination.
func QuerySellOrdersByBatchDenomCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "sell-orders-by-batch-denom [batch_denom]",
		Short: "List all sell orders by batch denom with pagination",
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
			res, err := c.SellOrdersByBatchDenom(cmd.Context(), &ecocredit.QuerySellOrdersByBatchDenomRequest{
				BatchDenom: args[0],
				Pagination: pagination,
			})
			return print(ctx, res, err)
		},
	})
}

// QueryBuyOrderCmd returns a query command that retrieves information for a given buy order.
func QueryBuyOrderCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "buy-order [buy_order_id]",
		Short: "Retrieve information for a given buy order",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			buyOrderId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return ecocredit.ErrInvalidBuyOrder.Wrap(err.Error())
			}
			res, err := c.BuyOrder(cmd.Context(), &ecocredit.QueryBuyOrderRequest{
				BuyOrderId: buyOrderId,
			})
			return print(ctx, res, err)
		},
	})
}

// QueryBuyOrdersCmd returns a query command that retrieves all buy orders with pagination.
func QueryBuyOrdersCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "buy-orders",
		Short: "List all buy orders with pagination",
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
			res, err := c.BuyOrders(cmd.Context(), &ecocredit.QueryBuyOrdersRequest{
				Pagination: pagination,
			})
			return print(ctx, res, err)
		},
	})
}

// QueryBuyOrdersByAddressCmd returns a query command that retrieves all buy orders by address with pagination.
func QueryBuyOrdersByAddressCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "buy-orders-by-address [address]",
		Short: "List all buy orders by buyer address with pagination",
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
			res, err := c.BuyOrdersByAddress(cmd.Context(), &ecocredit.QueryBuyOrdersByAddressRequest{
				Address:    args[0],
				Pagination: pagination,
			})
			return print(ctx, res, err)
		},
	})
}

// QueryAllowedAskDenomsCmd returns a query command that retrieves all allowed ask denoms with pagination.
func QueryAllowedAskDenomsCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "allowed-ask-denoms",
		Short: "List all allowed ask denoms with pagination",
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
			res, err := c.AllowedAskDenoms(cmd.Context(), &ecocredit.QueryAllowedAskDenomsRequest{
				Pagination: pagination,
			})
			return print(ctx, res, err)
		},
	})
}
