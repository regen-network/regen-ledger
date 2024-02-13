package keeper

import (
	"bytes"
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func (k Keeper) CreateOrUpdateApplication(ctx context.Context, msg *types.MsgCreateOrUpdateApplication) (*types.MsgCreateOrUpdateApplicationResponse, error) {
	admin, err := sdk.AccAddressFromBech32(msg.ProjectAdmin)
	if err != nil {
		return nil, err
	}

	proj, err := k.stateStore.ProjectTable().GetById(ctx, msg.ProjectId)
	if err != nil {
		return nil, err
	}

	if bytes.Equal(proj.Admin, admin) {
		return nil, sdkerrors.ErrUnauthorized
	}

	class, err := k.stateStore.ClassTable().GetById(ctx, msg.ClassId)
	if err != nil {
		return nil, err
	}

	enrollment, err := k.stateStore.ProjectEnrollmentTable().Get(ctx, proj.Key, class.Key)
	if ormerrors.IsNotFound(err) {
		enrollment = &ecocreditv1.ProjectEnrollment{
			ProjectKey: proj.Key,
			ClassKey:   class.Key,
			Status:     ecocreditv1.ProjectEnrollmentStatus_PROJECT_ENROLLMENT_STATUS_UNSPECIFIED,
		}
	} else if err != nil {
		return nil, err
	}

	enrollment.ApplicationMetadata = msg.Metadata

	if err := k.stateStore.ProjectEnrollmentTable().Save(ctx, enrollment); err != nil {
		return nil, err
	}

	return &types.MsgCreateOrUpdateApplicationResponse{}, nil
}
