package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
)

// Projects queries all projects from a given credit class.
func (k Keeper) Projects(ctx context.Context, request *v1.QueryProjectsRequest) (*v1.QueryProjectsResponse, error) {
	pg, err := GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}
	cInfo, err := k.stateStore.ClassInfoStore().GetByName(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.ProjectInfoStore().List(ctx, ecocreditv1.ProjectInfoClassIdNameIndexKey{}.WithClassId(cInfo.Id), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	projectInfos := make([]*v1.ProjectInfo, 0)
	for it.Next() {
		info, err := it.Value()
		if err != nil {
			return nil, err
		}
		var pi v1.ProjectInfo
		if err = PulsarToGogoSlow(info, &pi); err != nil {
			return nil, err
		}
		projectInfos = append(projectInfos, &pi)
	}
	pr, err := PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &v1.QueryProjectsResponse{
		Projects:   projectInfos,
		Pagination: pr,
	}, nil
}
