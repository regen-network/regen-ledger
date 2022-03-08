package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
)

// Batches queries for all batches in the given credit class.
func (k Keeper) Batches(ctx context.Context, request *v1beta1.QueryBatchesRequest) (*v1beta1.QueryBatchesResponse, error) {
	pg, err := GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}
	project, err := k.stateStore.ProjectInfoStore().GetByName(ctx, request.ProjectId)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.BatchInfoStore().List(ctx, ecocreditv1beta1.BatchInfoProjectIdIndexKey{}.WithProjectId(project.Id), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	batches := make([]*v1beta1.BatchInfo, 0)
	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return nil, err
		}
		var bi v1beta1.BatchInfo
		if err = PulsarToGogoSlow(batch, &bi); err != nil {
			return nil, err
		}
		batches = append(batches, &bi)
	}
	pr, err := PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &v1beta1.QueryBatchesResponse{
		Batches:    batches,
		Pagination: pr,
	}, nil
}
