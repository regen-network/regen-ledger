package marketplace

import (
	"github.com/regen-network/regen-ledger/x/ecocredit/core"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) getAllowedDenomSet(ctx sdk.Context) map[string]struct{} {
	var allowedDenoms []*core.AskDenom
	k.paramsKeeper.Get(ctx, core.KeyAllowedAskDenoms, &allowedDenoms)
	set := make(map[string]struct{}, len(allowedDenoms))
	for _, askDenom := range allowedDenoms {
		set[askDenom.Denom] = struct{}{}
	}
	return set
}
