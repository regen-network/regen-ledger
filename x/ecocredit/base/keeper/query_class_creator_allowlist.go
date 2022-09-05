package keeper

import (
	"context"

	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// ClassCreatorAllowlist queries credit class allowlist setting.
func (k Keeper) ClassCreatorAllowlist(ctx context.Context, request *types.QueryClassCreatorAllowlistRequest) (*types.QueryClassCreatorAllowlistResponse, error) {
	result, err := k.stateStore.ClassCreatorAllowlistTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryClassCreatorAllowlistResponse{
		Enabled: result.Enabled,
	}, nil
}
