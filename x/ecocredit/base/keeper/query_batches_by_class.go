package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// BatchesByClass queries all batches issued under a given credit class.
func (k Keeper) BatchesByClass(ctx context.Context, req *types.QueryBatchesByClassRequest) (*types.QueryBatchesByClassResponse, error) {
	class, err := k.stateStore.ClassTable().GetById(ctx, req.ClassId)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("could not get class with id %s: %s", req.ClassId, err.Error())
	}

	// we put a "-" after the class name to avoid including class names outside of the query (i.e. a query for C01 could technically include C011 otherwise).
	pg := ormutil.PageReqToOrmPaginate(req.Pagination)
	it, err := k.stateStore.BatchTable().List(ctx, api.BatchDenomIndexKey{}.WithDenom(class.Id+"-"), pg)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	batches := make([]*types.BatchInfo, 0, 10)
	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return nil, err
		}

		issuer := sdk.AccAddress(batch.Issuer)

		project, err := k.stateStore.ProjectTable().Get(ctx, batch.ProjectKey)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("unable to get project with key: %d", batch.ProjectKey)
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

	return &types.QueryBatchesByClassResponse{
		Batches:    batches,
		Pagination: ormutil.PageResToCosmosTypes(it.PageResponse()),
	}, nil
}
