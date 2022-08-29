package client

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	baseclient "github.com/regen-network/regen-ledger/x/ecocredit/base/client"
	basketcli "github.com/regen-network/regen-ledger/x/ecocredit/client/basket"
	marketplacecli "github.com/regen-network/regen-ledger/x/ecocredit/client/marketplace"
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
		basketcli.QueryBasketCmd(),
		basketcli.QueryBasketsCmd(),
		basketcli.QueryBasketBalanceCmd(),
		basketcli.QueryBasketBalancesCmd(),
		basketcli.QueryBasketFeesCmd(),
		marketplacecli.QuerySellOrderCmd(),
		marketplacecli.QuerySellOrdersCmd(),
		marketplacecli.QuerySellOrdersBySellerCmd(),
		marketplacecli.QuerySellOrdersByBatchCmd(),
		marketplacecli.QueryAllowedDenomsCmd(),
	)
	return cmd
}
