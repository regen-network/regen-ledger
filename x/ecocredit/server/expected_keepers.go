package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/server/core"
)

// Keeper defines the expected interface needed to prune expired buy and sell orders.
type Keeper interface {
	core.ProposalKeeper
	PruneOrders(ctx sdk.Context) error
}
