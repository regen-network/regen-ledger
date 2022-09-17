package keeper

import (
	"context"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func (k Keeper) SetClassCreatorAllowlist(ctx context.Context, req *types.MsgSetClassCreatorAllowlist) (*types.MsgSetClassCreatorAllowlistResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	if err := k.stateStore.ClassCreatorAllowlistTable().Save(ctx, &api.ClassCreatorAllowlist{
		Enabled: req.Enabled,
	}); err != nil {
		return nil, err
	}

	return &types.MsgSetClassCreatorAllowlistResponse{}, nil
}
