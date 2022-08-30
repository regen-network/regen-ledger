package core

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// AlloweClassCreators queries credit class allowlist setting.
func (k Keeper) CreditClassAllowlistEnabled(ctx context.Context, request *core.QueryCreditClassAllowlistEnabledRequest) (*core.QueryCreditClassAllowlistEnabledResponse, error) {
	result, err := k.stateStore.AllowListEnabledTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	return &core.QueryCreditClassAllowlistEnabledResponse{
		AllowlistEnabled: result.Enabled,
	}, nil
}
