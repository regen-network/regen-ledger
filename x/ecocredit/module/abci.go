package module

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/server"
)

// BeginBlocker checks if there are any expired sell orders and removes them from state.
func BeginBlocker(ctx sdk.Context, k server.Keeper) error {
	defer telemetry.ModuleMeasureSince(ecocredit.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	return k.PruneOrders(ctx)
}
