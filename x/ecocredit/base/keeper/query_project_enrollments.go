package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	"github.com/cosmos/cosmos-sdk/types/query"

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

	var it ecocreditv1.ProjectEnrollmentIterator
	if request.ClassId != "" && request.ProjectId != "" {
		res, err := k.ProjectEnrollment(ctx, &types.QueryProjectEnrollmentRequest{
			ProjectId: request.ProjectId,
			ClassId:   request.ClassId,
		})
		if err != nil {
			return nil, err
		}

		return &types.QueryProjectEnrollmentsResponse{
			Enrollments: []*types.EnrollmentInfo{res.Enrollment},
			Pagination:  &query.PageResponse{Total: 1},
		}, nil
	} else if request.ProjectId != "" {
		project, err := k.stateStore.ProjectTable().GetById(ctx, request.ProjectId)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("could not get project with id %s: %s", request.ProjectId, err.Error())
		}

		it, err = k.stateStore.ProjectEnrollmentTable().List(ctx,
			ecocreditv1.ProjectEnrollmentProjectKeyClassKeyIndexKey{}.WithProjectKey(project.Key),
			ormlist.Paginate(pg))
	} else if request.ClassId != "" {
		cls, err := k.stateStore.ClassTable().GetById(ctx, request.ClassId)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("could not get class with id %s: %s", request.ClassId, err.Error())
		}

		it, err = k.stateStore.ProjectEnrollmentTable().List(ctx,
			ecocreditv1.ProjectEnrollmentClassKeyIndexKey{}.WithClassKey(cls.Key),
			ormlist.Paginate(pg))
	} else {
		it, err = k.stateStore.ProjectEnrollmentTable().List(ctx,
			ecocreditv1.ProjectEnrollmentPrimaryKey{},
			ormlist.Paginate(pg))
	}
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

		project, err := k.stateStore.ProjectTable().Get(ctx, enrollment.ProjectKey)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("could not get project with key: %d", enrollment.ProjectKey)
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
