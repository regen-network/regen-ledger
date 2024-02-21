package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func (k Keeper) UpdateProjectEnrollment(ctx context.Context, msg *types.MsgUpdateProjectEnrollment) (*types.MsgUpdateProjectEnrollmentResponse, error) {
	issuer, err := sdk.AccAddressFromBech32(msg.Issuer)
	if err != nil {
		return nil, err
	}

	class, err := k.stateStore.ClassTable().GetById(ctx, msg.ClassId)
	if err != nil {
		return nil, err
	}

	err = k.assertClassIssuer(ctx, class.Key, issuer)
	if err != nil {
		return nil, err
	}

	proj, err := k.stateStore.ProjectTable().GetById(ctx, msg.ProjectId)
	if err != nil {
		return nil, err
	}

	enrollment, err := k.stateStore.ProjectEnrollmentTable().Get(ctx, proj.Key, class.Key)
	if err != nil {
		return nil, err
	}

	existingStatus := enrollment.Status
	newStatus := ecocreditv1.ProjectEnrollmentStatus(msg.NewStatus)
	delete := false
	switch existingStatus {
	case ecocreditv1.ProjectEnrollmentStatus_PROJECT_ENROLLMENT_STATUS_UNSPECIFIED,
		ecocreditv1.ProjectEnrollmentStatus_PROJECT_ENROLLMENT_STATUS_CHANGES_REQUESTED:
		switch newStatus {
		case ecocreditv1.ProjectEnrollmentStatus_PROJECT_ENROLLMENT_STATUS_CHANGES_REQUESTED,
			ecocreditv1.ProjectEnrollmentStatus_PROJECT_ENROLLMENT_STATUS_ACCEPTED:
			// Valid case
			break
		case ecocreditv1.ProjectEnrollmentStatus_PROJECT_ENROLLMENT_STATUS_REJECTED:
			delete = true
			break
		default:
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("invalid status transition from %s to %s", existingStatus, newStatus)
		}
	case ecocreditv1.ProjectEnrollmentStatus_PROJECT_ENROLLMENT_STATUS_ACCEPTED:
		switch newStatus {
		case ecocreditv1.ProjectEnrollmentStatus_PROJECT_ENROLLMENT_STATUS_TERMINATED:
			delete = true
			break
		default:
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("invalid status transition from %s to %s", existingStatus, newStatus)
		}
	default:
		return nil, sdkerrors.ErrLogic.Wrapf("invalid existing status: %s", existingStatus)
	}

	enrollment.Status = newStatus
	enrollment.EnrollmentMetadata = msg.Metadata

	if delete {
		if err := k.stateStore.ProjectEnrollmentTable().Delete(ctx, enrollment); err != nil {
			return nil, err
		}

	} else {
		if err := k.stateStore.ProjectEnrollmentTable().Save(ctx, enrollment); err != nil {
			return nil, err
		}
	}

	if err := sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&types.EventUpdateProjectEnrollment{
		ProjectId:             proj.Id,
		ClassId:               class.Id,
		NewStatus:             types.ProjectEnrollmentStatus(newStatus),
		NewEnrollmentMetadata: msg.Metadata,
	}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateProjectEnrollmentResponse{}, nil
}
