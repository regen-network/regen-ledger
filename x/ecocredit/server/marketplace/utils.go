package marketplace

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"

	"github.com/cosmos/cosmos-sdk/types"
)

// isDenomAllowed checks if the denom is allowed to be used in orders.
func isDenomAllowed(ctx context.Context, denom string, pk ecocredit.ParamKeeper) bool {
	sdkCtx := types.UnwrapSDKContext(ctx)
	var params core.Params
	pk.GetParamSet(sdkCtx, &params)
	for _, askDenom := range params.AllowedAskDenoms {
		if askDenom.Denom == denom {
			return true
		}
	}
	return false
}
