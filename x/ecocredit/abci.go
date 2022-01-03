package ecocredit

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker checks if there are any expired sell or buy orders and removes them from state.
func BeginBlocker(ctx sdk.Context, k Keeper, req abci.RequestBeginBlock) {
	defer telemetry.ModuleMeasureSince(ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	if err := k.PruneOrders(ctx); err != nil {
		panic(err)
	}
}
