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
	cInfo, err := k.stateStore.ClassInfoTable().GetByName(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.ProjectInfoTable().List(ctx, api.ProjectInfoClassIdNameIndexKey{}.WithClassId(cInfo.Id), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	projectInfos := make([]*core.ProjectInfoEntry, 0)
	for it.Next() {
		info, err := it.Value()
		if err != nil {
			return nil, err
		}

		admin := sdk.AccAddress(info.Admin)

		class, err := k.stateStore.ClassInfoTable().Get(ctx, info.ClassId)
		if err != nil {
			return nil, err
		}

		project := core.ProjectInfoEntry{
			Id:              info.Name,
			Admin:           admin.String(),
			ClassId:         class.Name,
			ProjectLocation: info.ProjectLocation,
			Metadata:        info.Metadata,
		}

		projectInfos = append(projectInfos, &project)
	}
	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &core.QueryProjectsResponse{
		Projects:   projectInfos,
		Pagination: pr,
	}, nil
}
