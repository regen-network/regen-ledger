package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// UpdateClassIssuers updates a class's issuers by either adding more issuers, or removing issuers from the class issuer store.
func (k Keeper) UpdateClassIssuers(ctx context.Context, req *types.MsgUpdateClassIssuers) (*types.MsgUpdateClassIssuersResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	reqAddr, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	class, err := k.stateStore.ClassTable().GetById(ctx, req.ClassId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf(
			"could not get credit class with id %s: %s", req.ClassId, err,
		)
	}

	admin := sdk.AccAddress(class.Admin)
	if !reqAddr.Equals(admin) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"%s is not the admin of credit class %s", req.Admin, req.ClassId,
		)
	}

	if err = sdkCtx.EventManager().EmitTypedEvent(&types.EventUpdateClassIssuers{
		ClassId: req.ClassId,
		Added:   req.AddIssuers,
		Removed: req.RemoveIssuers,
	}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateClassIssuersResponse{}, nil
}
