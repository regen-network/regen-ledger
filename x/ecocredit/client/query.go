package client

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	baseclient "github.com/regen-network/regen-ledger/x/ecocredit/base/client"
	basketclient "github.com/regen-network/regen-ledger/x/ecocredit/basket/client"
	marketclient "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/client"
)

// QueryCmd returns the parent command for all x/ecocredit query commands.
func QueryCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		SuggestionsMinimumDistance: 2,
		DisableFlagParsing:         true,
		Args:                       cobra.ExactArgs(1),
		Use:                        name,
		Short:                      "Query commands for the ecocredit module",
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(
		baseclient.QueryClassesCmd(),
		baseclient.QueryClassCmd(),
		baseclient.QueryClassIssuersCmd(),
		baseclient.QueryBatchesCmd(),
		baseclient.QueryBatchesByIssuerCmd(),
		baseclient.QueryBatchesByClassCmd(),
		baseclient.QueryBatchesByProjectCmd(),
		baseclient.QueryBatchCmd(),
		baseclient.QueryBatchBalanceCmd(),
		baseclient.QueryBatchSupplyCmd(),
		baseclient.QueryCreditTypesCmd(),
		baseclient.QueryProjectsCmd(),
		baseclient.QueryProjectsByClassCmd(),
		baseclient.QueryProjectsByReferenceIDCmd(),
		baseclient.QueryProjectsByAdminCmd(),
		baseclient.QueryProjectCmd(),
		baseclient.QueryParamsCmd(),
		baseclient.QueryCreditTypeCmd(),
		baseclient.QueryClassCreatorAllowlistCmd(),
		baseclient.QueryAllowedClassCreatorsCmd(),
		baseclient.QueryCreditClassFeesCmd(),
		baseclient.QueryAllBalances(),
		basketclient.QueryBasketCmd(),
		basketclient.QueryBasketsCmd(),
		basketclient.QueryBasketBalanceCmd(),
		basketclient.QueryBasketBalancesCmd(),
		basketclient.QueryBasketFeesCmd(),
		marketclient.QuerySellOrderCmd(),
		marketclient.QuerySellOrdersCmd(),
		marketclient.QuerySellOrdersBySellerCmd(),
		marketclient.QuerySellOrdersByBatchCmd(),
		marketclient.QueryAllowedDenomsCmd(),
	)
	return cmd
}
