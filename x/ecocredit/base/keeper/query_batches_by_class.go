package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/errors"
	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// BatchesByClass queries all batches issued under a given credit class.
func (k Keeper) BatchesByClass(ctx context.Context, request *types.QueryBatchesByClassRequest) (*types.QueryBatchesByClassResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	class, err := k.stateStore.ClassTable().GetById(ctx, request.ClassId)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("could not get class with id %s: %s", request.ClassId, err.Error())
	}

	// we put a "-" after the class name to avoid including class names outside of the query (i.e. a query for C01 could technically include C011 otherwise).
	it, err := k.stateStore.BatchTable().List(ctx, api.BatchDenomIndexKey{}.WithDenom(class.Id+"-"), ormlist.Paginate(pg))
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

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}

	return &types.QueryBatchesByClassResponse{
		Batches:    batches,
		Pagination: pr,
	}, nil
}
