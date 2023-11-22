package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
)

// RemoveAllowedDenom removes denom from the allowed denoms.
func (k Keeper) RemoveAllowedDenom(ctx context.Context, req *types.MsgRemoveAllowedDenom) (*types.MsgRemoveAllowedDenomResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	allowedDenom, err := k.stateStore.AllowedDenomTable().Get(ctx, req.Denom)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("allowed denom %s", req.Denom)
		}

		return nil, err
	}

	if err := k.stateStore.AllowedDenomTable().Delete(ctx, allowedDenom); err != nil {
		return nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if err := sdkCtx.EventManager().EmitTypedEvent(&types.EventRemoveAllowedDenom{Denom: req.Denom}); err != nil {
		return nil, err
	}

	return &types.MsgRemoveAllowedDenomResponse{}, nil
}
