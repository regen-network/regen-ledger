package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// ProjectsByClass queries all projects from a given credit class.
func (k Keeper) ProjectsByClass(ctx context.Context, request *types.QueryProjectsByClassRequest) (*types.QueryProjectsByClassResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	cInfo, err := k.stateStore.ClassTable().GetById(ctx, request.ClassId)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("could not get class with id %s: %s", request.ClassId, err.Error())
	}

	it, err := k.stateStore.ProjectEnrollmentTable().List(ctx, api.ProjectEnrollmentClassKeyIndexKey{}.WithClassKey(cInfo.Key), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	projects := make([]*types.ProjectInfo, 0)
	for it.Next() {
		enrollment, err := it.Value()
		if err != nil {
			return nil, err
		}

		project, err := k.stateStore.ProjectTable().Get(ctx, enrollment.ProjectKey)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("could not get project with key: %d", enrollment.ProjectKey)
		}

		admin := sdk.AccAddress(project.Admin)

		info := types.ProjectInfo{
			Id:           project.Id,
			Admin:        admin.String(),
			Jurisdiction: project.Jurisdiction,
			Metadata:     project.Metadata,
			ReferenceId:  project.ReferenceId,
		}

		projects = append(projects, &info)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}

	return &types.QueryProjectsByClassResponse{
		Projects:   projects,
		Pagination: pr,
	}, nil
}
