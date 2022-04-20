package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Batches queries for all batches in the given credit class.
func (k Keeper) Batches(ctx context.Context, request *core.QueryBatchesRequest) (*core.QueryBatchesResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}
	project, err := k.stateStore.ProjectInfoTable().GetById(ctx, request.ProjectId)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.BatchInfoTable().List(ctx, api.BatchInfoProjectKeyIndexKey{}.WithProjectKey(project.Key), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	batches := make([]*core.BatchInfo, 0)
	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return nil, err
		}

		var bi core.BatchInfo
		if err = ormutil.PulsarToGogoSlow(batch, &bi); err != nil {
			return nil, err
		}
		batches = append(batches, &bi)
	}
	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &core.QueryBatchesResponse{
		Batches:    batches,
		Pagination: pr,
	}, nil
}
