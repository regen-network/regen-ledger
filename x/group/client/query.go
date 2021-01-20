package client

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group gov queries under a subcommand
	queryCmd := &cobra.Command{
		Use:                        group.ModuleName,
		Short:                      "Querying commands for the governance module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	queryCmd.AddCommand(
		GetCmdQueryProposal(),
		GetCmdQueryProposals(),
		GetCmdQueryVote(),
		GetCmdQueryVotes(),
		GetCmdQueryParam(),
		GetCmdQueryParams(),
		GetCmdQueryProposer(),
		GetCmdQueryDeposit(),
		GetCmdQueryDeposits(),
		GetCmdQueryTally(),
	)

	return govQueryCmd
}