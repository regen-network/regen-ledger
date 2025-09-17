package keeper

import (
	"context"

	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

func (k Keeper) SetClassCreatorAllowlist(ctx context.Context, req *types.MsgSetClassCreatorAllowlist) (*types.MsgSetClassCreatorAllowlistResponse, error) {
	authorityBz, err := k.ac.StringToBytes(req.Authority)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid authority address")
	}
	authority := sdk.AccAddress(authorityBz)

	if !authority.Equals(k.authority) {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	if err := k.stateStore.ClassCreatorAllowlistTable().Save(ctx, &api.ClassCreatorAllowlist{
		Enabled: req.Enabled,
	}); err != nil {
		return nil, err
	}

	return &types.MsgSetClassCreatorAllowlistResponse{}, nil
}
