package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Keeper defines the expected interface needed to prune expired buy and sell orders.
type Keeper interface {
	PruneOrders(ctx sdk.Context) error
}

// ProposalKeeper defines the expected interface for ecocredit module proposals.
type ProposalKeeper interface {
	NewCreditType(ctx sdk.Context, ctp *core.CreditTypeProposal) error
}
