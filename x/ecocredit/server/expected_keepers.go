package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper defines a set of methods the ecocredit module exposes.
type Keeper interface {
	PruneOrders(ctx sdk.Context) error
}
