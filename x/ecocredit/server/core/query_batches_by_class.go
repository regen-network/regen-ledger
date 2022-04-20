package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// BatchesByClass queries all batches issued under a given credit class.
func (k Keeper) BatchesByClass(ctx context.Context, request *core.QueryBatchesByClassRequest) (*core.QueryBatchesByClassResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}
	class, err := k.stateStore.ClassInfoTable().GetById(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}
	// we put a "-" after the class name to avoid including class names outside of the query (i.e. a query for C01 could technically include C011 otherwise).
	it, err := k.stateStore.BatchInfoTable().List(ctx, api.BatchInfoDenomIndexKey{}.WithDenom(class.Id+"-"), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	batches := make([]*core.BatchInfo, 0, 10)
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
	return &core.QueryBatchesByClassResponse{
		Batches:    batches,
		Pagination: pr,
	}, nil
}
