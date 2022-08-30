package keeper

import (
	"context"

	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// AlloweClassCreators queries credit class allowlist setting.
func (k Keeper) CreditClassAllowlistEnabled(ctx context.Context, request *types.QueryCreditClassAllowlistEnabledRequest) (*types.QueryCreditClassAllowlistEnabledResponse, error) {
	result, err := k.stateStore.AllowListEnabledTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryCreditClassAllowlistEnabledResponse{
		AllowlistEnabled: result.Enabled,
	}, nil
}
