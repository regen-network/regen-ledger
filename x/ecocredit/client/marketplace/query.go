package marketplace

import (
	"strconv"

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
		Args:  cobra.ExactArgs(1),
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
		Args:  cobra.ExactArgs(0),
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
		Args:  cobra.ExactArgs(1),
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
		Args:  cobra.ExactArgs(1),
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

// QueryBuyOrderCmd returns a query command that retrieves information for a given buy order.
func QueryBuyOrderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-order [buy_order_id]",
		Short: "Retrieve information for a given buy order",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := sdkclient.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			client := marketplace.NewQueryClient(ctx)
			buyOrderId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return ecocredit.ErrInvalidBuyOrder.Wrap(err.Error())
			}
			res, err := client.BuyOrder(cmd.Context(), &marketplace.QueryBuyOrderRequest{
				BuyOrderId: buyOrderId,
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

// QueryBuyOrdersCmd returns a query command that retrieves all buy orders with pagination.
func QueryBuyOrdersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-orders",
		Short: "List all buy orders with pagination",
		Args:  cobra.ExactArgs(0),
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
			res, err := client.BuyOrders(cmd.Context(), &marketplace.QueryBuyOrdersRequest{
				Pagination: pagination,
			})
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "buy-orders")

	return cmd
}

// QueryBuyOrdersByAddressCmd returns a query command that retrieves all buy orders by address with pagination.
func QueryBuyOrdersByAddressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-orders-by-address [address]",
		Short: "List all buy orders by buyer address with pagination",
		Args:  cobra.ExactArgs(1),
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
			res, err := client.BuyOrdersByAddress(cmd.Context(), &marketplace.QueryBuyOrdersByAddressRequest{
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
	flags.AddPaginationFlagsToCmd(cmd, "buy-orders-by-address")

	return cmd
}

// QueryAllowedDenomsCmd returns a query command that retrieves all allowed ask denoms with pagination.
func QueryAllowedDenomsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allowed-denoms",
		Short: "List all allowed denoms with pagination",
		Args:  cobra.ExactArgs(0),
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
