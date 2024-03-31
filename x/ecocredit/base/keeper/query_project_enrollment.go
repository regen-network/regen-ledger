package keeper

import (
	"context"

	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func (k Keeper) ProjectEnrollment(ctx context.Context, request *types.QueryProjectEnrollmentRequest) (*types.QueryProjectEnrollmentResponse, error) {
	project, err := k.stateStore.ProjectTable().GetById(ctx, request.ProjectId)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("could not get project with id %s: %s", request.ProjectId, err.Error())
	}

	class, err := k.stateStore.ClassTable().GetById(ctx, request.ClassId)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("could not get class with id %s: %s", request.ClassId, err.Error())
	}

	enrollment, err := k.stateStore.ProjectEnrollmentTable().Get(ctx, project.Key, class.Key)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("could not get enrollment for project %s and class %s: %s", project.Id, class.Id, err.Error())
	}

	info := types.EnrollmentInfo{
		ProjectId:           project.Id,
		ClassId:             class.Id,
		Status:              types.ProjectEnrollmentStatus(enrollment.Status),
		ApplicationMetadata: enrollment.ApplicationMetadata,
		EnrollmentMetadata:  enrollment.EnrollmentMetadata,
	}

	return &types.QueryProjectEnrollmentResponse{Enrollment: &info}, nil
}
