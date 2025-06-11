package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

// Batches queries for all batches in the given credit class.
func (k Keeper) Batches(ctx context.Context, req *types.QueryBatchesRequest) (*types.QueryBatchesResponse, error) {
	pg := ormutil.PageReqToOrmPaginate(req.Pagination)
	it, err := k.stateStore.BatchTable().List(ctx, api.BatchPrimaryKey{}, pg)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	batches := make([]*types.BatchInfo, 0)
	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return nil, err
		}

		issuer := sdk.AccAddress(batch.Issuer)
		project, err := k.stateStore.ProjectTable().Get(ctx, batch.ProjectKey)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("failed to get project by key: %d: %s",
				batch.ProjectKey, err.Error())
		}
		info := types.BatchInfo{
			Issuer:       issuer.String(),
			ProjectId:    project.Id,
			Denom:        batch.Denom,
			Metadata:     batch.Metadata,
			StartDate:    regentypes.ProtobufToGogoTimestamp(batch.StartDate),
			EndDate:      regentypes.ProtobufToGogoTimestamp(batch.EndDate),
			IssuanceDate: regentypes.ProtobufToGogoTimestamp(batch.IssuanceDate),
			Open:         batch.Open,
		}
		batches = append(batches, &info)
	}

	return &types.QueryBatchesResponse{
		Batches:    batches,
		Pagination: ormutil.PageResToCosmosTypes(it.PageResponse()),
	}, nil
}
