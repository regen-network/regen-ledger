package client

import (
	"strconv"

	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
)

// QuerySellOrderCmd returns a query command that retrieves information for a given sell order.
func QuerySellOrderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sell-order [sell-order-id]",
		Short:   "Retrieve information for a given sell order",
		Long:    "Retrieve information for a given sell order.",
		Example: "regen q ecocredit sell-order 1",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := sdkclient.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			client := types.NewQueryClient(ctx)
			sellOrderID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return ecocredit.ErrInvalidSellOrder.Wrap(err.Error())
			}
			res, err := client.SellOrder(cmd.Context(), &types.QuerySellOrderRequest{
				SellOrderId: sellOrderID,
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
		Short: "List all sell orders",
		Long:  `List all sell orders with optional pagination flags.`,
		Example: `regen q sell-orders
regen q sell-orders --limit 10 --offset 10`,
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := sdkclient.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			client := types.NewQueryClient(ctx)
			pagination, err := sdkclient.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := client.SellOrders(cmd.Context(), &types.QuerySellOrdersRequest{
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
		Short: "List all sell orders by seller",
		Long:  `List all sell orders by seller with optional pagination flags.`,
		Example: `regen q ecocredit sell-orders-by-seller regen18xvpj53vaupyfejpws5sktv5lnas5xj2phm3cf
regen q ecocredit sell-orders-by-seller regen18xvpj53vaupyfejpws5sktv5lnas5xj2phm3cf --limit 10 --offset 10`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := sdkclient.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			client := types.NewQueryClient(ctx)
			pagination, err := sdkclient.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := client.SellOrdersBySeller(cmd.Context(), &types.QuerySellOrdersBySellerRequest{
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

// QuerySellOrdersByBatchCmd returns a query command that retrieves all sell orders by batch denom with pagination.
func QuerySellOrdersByBatchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell-orders-by-batch [batch_denom]",
		Short: "List all sell orders by batch denom",
		Long:  "List all sell orders by batch by denom with optional pagination flags.",
		Example: `regen q ecocredit sell-orders-by-batch C01-001-20200101-20210101-001
regen q ecocredit sell-orders-by-batch C01-001-20200101-20210101-001 --limit 10 --offset 10`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := sdkclient.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			client := types.NewQueryClient(ctx)
			pagination, err := sdkclient.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := client.SellOrdersByBatch(cmd.Context(), &types.QuerySellOrdersByBatchRequest{
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
	flags.AddPaginationFlagsToCmd(cmd, "sell-orders-by-batch")
	return cmd
}

// QueryAllowedDenomsCmd returns a query command that retrieves allowed denoms with pagination.
func QueryAllowedDenomsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allowed-denoms",
		Short: "List all allowed denoms",
		Long:  "List all allowed denoms with optional pagination flags.",
		Example: `regen q ecocredit allowed-denoms
regen q ecocredit allowed-denoms --limit 10 --offset 10`,
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := sdkclient.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			client := types.NewQueryClient(ctx)
			pagination, err := sdkclient.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := client.AllowedDenoms(cmd.Context(), &types.QueryAllowedDenomsRequest{
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
