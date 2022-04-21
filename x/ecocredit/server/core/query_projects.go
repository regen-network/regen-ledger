package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Projects queries all projects from a given credit class.
func (k Keeper) Projects(ctx context.Context, request *core.QueryProjectsRequest) (*core.QueryProjectsResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	cInfo, err := k.stateStore.ClassInfoTable().GetById(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.ProjectInfoTable().List(ctx, api.ProjectInfoClassKeyIdIndexKey{}.WithClassKey(cInfo.Key), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	projects := make([]*core.ProjectDetails, 0)
	for it.Next() {
		project, err := it.Value()
		if err != nil {
			return nil, err
		}

		admin := sdk.AccAddress(project.Admin)

		class, err := k.stateStore.ClassInfoTable().Get(ctx, project.ClassKey)
		if err != nil {
			return nil, err
		}

		info := core.ProjectDetails{
			Id:              project.Id,
			Admin:           admin.String(),
			ClassId:         class.Id,
			ProjectLocation: project.ProjectJurisdiction,
			Metadata:        project.Metadata,
		}

		projects = append(projects, &info)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &core.QueryProjectsResponse{
		Projects:   projects,
		Pagination: pr,
	}, nil
}
