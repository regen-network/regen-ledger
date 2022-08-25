package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// UpdateClassMetadata updates the metadata for the class.
func (k Keeper) UpdateClassMetadata(ctx context.Context, req *core.MsgUpdateClassMetadata) (*core.MsgUpdateClassMetadataResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	reqAddr, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	classInfo, err := k.stateStore.ClassTable().GetById(ctx, req.ClassId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf(
			"could not get credit class with id %s: %s", req.ClassId, err,
		)
	}

	admin := sdk.AccAddress(classInfo.Admin)
	if !reqAddr.Equals(admin) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"%s is not the admin of credit class %s", req.Admin, req.ClassId,
		)
	}

	classInfo.Metadata = req.NewMetadata
	if err = k.stateStore.ClassTable().Update(ctx, classInfo); err != nil {
		return nil, err
	}

	if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventUpdateClassMetadata{
		ClassId: req.ClassId,
	}); err != nil {
		return nil, err
	}

	return &core.MsgUpdateClassMetadataResponse{}, err
}
