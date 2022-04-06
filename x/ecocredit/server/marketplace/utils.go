package marketplace

import (
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// isDenomAllowed checks if the denom is allowed to be used in orders.
func isDenomAllowed(ctx sdk.Context, denom string, pk ecocredit.ParamKeeper) bool {
	var params core.Params
	pk.GetParamSet(ctx, &params)
	for _, askDenom := range params.AllowedAskDenoms {
		if askDenom.Denom == denom {
			return true
		}
	}
	return false
}
