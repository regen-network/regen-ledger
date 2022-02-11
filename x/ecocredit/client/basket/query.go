package basketclient

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

// QueryBasketCmd returns a query command that retrieves a basket.
func QueryBasketCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "basket [basket-denom]",
		Short:   "Gets the info for a basket.",
		Long:    "Retrieves the information for a basket definition, given a specific basket denom.",
		Example: "regen q ecocredit basket BASKET",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			client := basket.NewQueryClient(ctx)

			res, err := client.Basket(cmd.Context(), &basket.QueryBasketRequest{BasketDenom: args[0]})
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
	return cmd
}

// QueryBasketBalanceCmd returns a query command that retrieves the balance of a credit batch in the basket.
func QueryBasketBalanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "basket-balance [basket-denom] [batch-denom]",
		Short: "Retrieves the balance of a credit batch in the basket",
		Long: strings.TrimSpace(`Retrieves the balance of a credit batch in the basket
Example:
		$regen q ecocredit basket-balance BASKET C01-20210101-20220101-001
		`),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			client := basket.NewQueryClient(ctx)

			res, err := client.BasketBalance(cmd.Context(), &basket.QueryBasketBalanceRequest{
				BasketDenom: args[0],
				BatchDenom:  args[1],
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

// QueryBasketBalancesCmd returns a query command that retrieves the the balance of each credit batch for the given basket denom.
func QueryBasketBalancesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "basket-balances [basket-denom]",
		Short: "Retrieves the the balance of each credit batch for the given basket denom",
		Long: strings.TrimSpace(`Retrieves the the balance of each credit batch for the given basket denom

Examples:
		$regen q ecocredit basket-balances BASKET1
		$regen q ecocredit basket-balances BASKET1 --pagination.offset 1 --pagination.limit 10
		`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			client := basket.NewQueryClient(ctx)
			res, err := client.BasketBalances(cmd.Context(), &basket.QueryBasketBalancesRequest{
				BasketDenom: args[0],
				Pagination:  pagination,
			})
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "basket-balances")

	return cmd
}
