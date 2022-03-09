package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
)

// Batches queries for all batches in the given credit class.
func (k Keeper) Batches(ctx context.Context, request *v1.QueryBatchesRequest) (*v1.QueryBatchesResponse, error) {
	pg, err := GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}
	project, err := k.stateStore.ProjectInfoStore().GetByName(ctx, request.ProjectId)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.BatchInfoStore().List(ctx, ecocreditv1.BatchInfoProjectIdIndexKey{}.WithProjectId(project.Id), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	batches := make([]*v1.BatchInfo, 0)
	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return nil, err
		}

		bz, err := types.DecodeMetadata(batch.Metadata)
		if err != nil {
			return nil, err
		}
		batch.Metadata = string(bz)

		var bi v1.BatchInfo
		if err = PulsarToGogoSlow(batch, &bi); err != nil {
			return nil, err
		}
		batches = append(batches, &bi)
	}
	pr, err := PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &v1.QueryBatchesResponse{
		Batches:    batches,
		Pagination: pr,
	}, nil
}
