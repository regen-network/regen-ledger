package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func UpgradeEco(params types.Subspace, ctx sdk.Context) error {
	params.Set(ctx, ecocredit.KeyAllowlistEnabled, true)
	return nil
}
