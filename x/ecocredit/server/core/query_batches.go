package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types"

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
	project, err := k.stateStore.ProjectInfoTable().GetByName(ctx, request.ProjectId)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.BatchInfoTable().List(ctx, api.BatchInfoProjectIdIndexKey{}.WithProjectId(project.Id), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	batches := make([]*core.BatchInfoEntry, 0)
	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return nil, err
		}

		issuer := sdk.AccAddress(batch.Issuer)

		entry := core.BatchInfoEntry{
			Issuer:       issuer.String(),
			ProjectId:    project.Name,
			BatchDenom:   batch.BatchDenom,
			Metadata:     batch.Metadata,
			StartDate:    types.ProtobufToGogoTimestamp(batch.StartDate),
			EndDate:      types.ProtobufToGogoTimestamp(batch.EndDate),
			IssuanceDate: types.ProtobufToGogoTimestamp(batch.IssuanceDate),
			Open:         batch.Open,
		}

		batches = append(batches, &entry)
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
