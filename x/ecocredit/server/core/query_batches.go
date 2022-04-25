package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Batches queries for all batches in the given credit class.
func (k Keeper) Batches(ctx context.Context, request *core.QueryBatchesRequest) (*core.QueryBatchesResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	project, err := k.stateStore.ProjectTable().GetById(ctx, request.ProjectId)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.BatchTable().List(ctx, api.BatchProjectKeyIndexKey{}.WithProjectKey(project.Key), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	batches := make([]*core.BatchInfo, 0)
	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return nil, err
		}

		issuer := sdk.AccAddress(batch.Issuer)

		info := core.BatchInfo{
			Issuer:       issuer.String(),
			ProjectId:    project.Id,
			Denom:        batch.Denom,
			Metadata:     batch.Metadata,
			StartDate:    types.ProtobufToGogoTimestamp(batch.StartDate),
			EndDate:      types.ProtobufToGogoTimestamp(batch.EndDate),
			IssuanceDate: types.ProtobufToGogoTimestamp(batch.IssuanceDate),
			Open:         batch.Open,
		}

		batches = append(batches, &info)
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
