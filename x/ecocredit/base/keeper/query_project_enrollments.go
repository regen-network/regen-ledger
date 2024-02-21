package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func (k Keeper) ProjectEnrollments(ctx context.Context, request *types.QueryProjectEnrollmentsRequest) (*types.QueryProjectEnrollmentsResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	project, err := k.stateStore.ProjectTable().GetById(ctx, request.ProjectId)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("could not get project with id %s: %s", request.ProjectId, err.Error())
	}

	it, err := k.stateStore.ProjectEnrollmentTable().List(ctx,
		ecocreditv1.ProjectEnrollmentProjectKeyClassKeyIndexKey{}.WithProjectKey(project.Key),
		ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	enrollments := make([]*types.EnrollmentInfo, 0)
	for it.Next() {
		enrollment, err := it.Value()
		if err != nil {
			return nil, err
		}

		class, err := k.stateStore.ClassTable().Get(ctx, enrollment.ClassKey)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("could not get class with key: %d", enrollment.ClassKey)
		}

		info := &types.EnrollmentInfo{
			ProjectId:           project.Id,
			ClassId:             class.Id,
			Status:              types.ProjectEnrollmentStatus(enrollment.Status),
			ApplicationMetadata: enrollment.ApplicationMetadata,
			EnrollmentMetadata:  enrollment.EnrollmentMetadata,
		}

		enrollments = append(enrollments, info)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}

	return &types.QueryProjectEnrollmentsResponse{
		Enrollments: enrollments,
		Pagination:  pr,
	}, nil
}
