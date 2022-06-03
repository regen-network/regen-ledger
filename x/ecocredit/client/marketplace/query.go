package marketplace

import (
	"strconv"
	"strings"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

// QuerySellOrderCmd returns a query command that retrieves information for a given sell order.
func QuerySellOrderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell-order [sell_order_id]",
		Short: "Retrieve information for a given sell order",
		Long: strings.TrimSpace(`Retrieve information for a given sell order
	
Example:
$ regen q sell-order 1
$ regen q sell-order 1 --output json
		`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := sdkclient.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			client := marketplace.NewQueryClient(ctx)
			sellOrderId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return ecocredit.ErrInvalidSellOrder.Wrap(err.Error())
			}
			res, err := client.SellOrder(cmd.Context(), &marketplace.QuerySellOrderRequest{
				SellOrderId: sellOrderId,
			})
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QuerySellOrdersCmd returns a query command that retrieves all sell orders with pagination.
func QuerySellOrdersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell-orders",
		Short: "List all sell orders with pagination",
		Long: strings.TrimSpace(`Retrieve sell orders with pagination
	
Example:
$ regen q sell-orders
$ regen q sell-orders --pagination.limit 10 --pagination.offset 2
		`),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := sdkclient.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			client := marketplace.NewQueryClient(ctx)
			pagination, err := sdkclient.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := client.SellOrders(cmd.Context(), &marketplace.QuerySellOrdersRequest{
				Pagination: pagination,
			})
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "sell-orders")

	return cmd
}

// QuerySellOrdersBySellerCmd returns a query command that retrieves all sell orders by address with pagination.
func QuerySellOrdersBySellerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell-orders-by-seller [seller]",
		Short: "List all sell orders by seller address with pagination",
		Long: strings.TrimSpace(
			`Retrieve sell orders by seller address with pagination
	
Example:
$ regen q sell-orders-by-seller regen1fv85...zkfu
$ regen q sell-orders-by-seller regen1fv85...zkfu --pagination.limit 10 --pagination.offset 2
		`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := sdkclient.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			client := marketplace.NewQueryClient(ctx)
			pagination, err := sdkclient.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := client.SellOrdersBySeller(cmd.Context(), &marketplace.QuerySellOrdersBySellerRequest{
				Seller:     args[0],
				Pagination: pagination,
			})
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "sell-orders-by-seller")

	return cmd
}

// QuerySellOrdersByBatchDenomCmd returns a query command that retrieves all sell orders by batch denom with pagination.
func QuerySellOrdersByBatchDenomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell-orders-by-batch-denom [batch_denom]",
		Short: "List all sell orders by batch denom with pagination",
		Long: strings.TrimSpace(`
		Retrieve sell orders by batch by denom with pagination
	
Example:
$ regen q sell-orders-by-batch-denom C01-20210101-20210201-001
$ regen q sell-orders-by-batch-denom C01-20210101-20210201-001 --pagination.limit 10 --pagination.offset 2
		`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := sdkclient.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			client := marketplace.NewQueryClient(ctx)
			pagination, err := sdkclient.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := client.SellOrdersByBatchDenom(cmd.Context(), &marketplace.QuerySellOrdersByBatchDenomRequest{
				BatchDenom: args[0],
				Pagination: pagination,
			})
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "sell-orders-by-batch-denom")
	return cmd
}

// QueryAllowedDenomsCmd returns a query command that retrieves allowed denoms with pagination.
func QueryAllowedDenomsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allowed-denoms",
		Short: "List all allowed denoms with pagination",
		Long: strings.TrimSpace(`
		Retrieve allowed denoms with pagination
	
Example:
$ regen q allowed-denoms
$ regen q allowed-denoms --pagination.limit 10 --pagination.offset 2
		`),
		Example: "$ regen q allowed-denoms",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := sdkclient.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			client := marketplace.NewQueryClient(ctx)
			pagination, err := sdkclient.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := client.AllowedDenoms(cmd.Context(), &marketplace.QueryAllowedDenomsRequest{
				Pagination: pagination,
			})
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "allowed-denoms")

	return cmd
}
