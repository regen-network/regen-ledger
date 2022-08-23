package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// UpdateClassAdmin updates the admin address for a class.
// WARNING: this method will forfeit control of the entire class to the provided address.
// double check your inputs to ensure you do not lose control of the class.
func (k Keeper) UpdateClassAdmin(ctx context.Context, req *core.MsgUpdateClassAdmin) (*core.MsgUpdateClassAdminResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	reqAddr, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}
	newAdmin, err := sdk.AccAddressFromBech32(req.NewAdmin)
	if err != nil {
		return nil, err
	}

	classInfo, err := k.stateStore.ClassTable().GetById(ctx, req.ClassId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf(
			"could not get credit class with id %s: %s", req.ClassId, err,
		)
	}

	classAdmin := sdk.AccAddress(classInfo.Admin)
	if !classAdmin.Equals(reqAddr) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"%s is not the admin of credit class %s", req.Admin, req.ClassId,
		)
	}
	classInfo.Admin = newAdmin
	if err = k.stateStore.ClassTable().Update(ctx, classInfo); err != nil {
		return nil, err
	}

	if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventUpdateClassAdmin{
		ClassId: req.ClassId,
	}); err != nil {
		return nil, err
	}

	return &core.MsgUpdateClassAdminResponse{}, err
}
