package keeper

import (
	"context"

	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// ClassCreatorAllowlist queries credit class allowlist setting.
func (k Keeper) ClassCreatorAllowlist(ctx context.Context, request *types.QueryClassCreatorAllowlistRequest) (*types.QueryClassCreatorAllowlistResponse, error) {
	result, err := k.stateStore.ClassCreatorAllowlistTable().Get(ctx)
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrapf("failed to get class creator allowlist: %s", err.Error())
	}

	return &types.QueryClassCreatorAllowlistResponse{
		Enabled: result.Enabled,
	}, nil
}
