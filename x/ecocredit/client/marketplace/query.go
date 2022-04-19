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

// QuerySellOrdersByAddressCmd returns a query command that retrieves all sell orders by address with pagination.
func QuerySellOrdersByAddressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell-orders-by-address [address]",
		Short: "List all sell orders by owner address with pagination",
		Long: strings.TrimSpace(
			`Retrieve sell orders by owner address with pagination
	
Example:
$ regen q sell-orders-by-address regen1fv85...zkfu
$ regen q sell-orders-by-address regen1fv85...zkfu --pagination.limit 10 --pagination.offset 2
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
			res, err := client.SellOrdersByAddress(cmd.Context(), &marketplace.QuerySellOrdersByAddressRequest{
				Address:    args[0],
				Pagination: pagination,
			})
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "sell-orders-by-address")

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
