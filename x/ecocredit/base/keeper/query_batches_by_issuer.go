package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// BatchesByIssuer queries all batches issued from the given issuer address
func (k Keeper) BatchesByIssuer(ctx context.Context, req *types.QueryBatchesByIssuerRequest) (*types.QueryBatchesByIssuerResponse, error) {
	issuer, err := sdk.AccAddressFromBech32(req.Issuer)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.BatchTable().List(ctx, api.BatchIssuerIndexKey{}.WithIssuer(issuer), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	batches := make([]*types.BatchInfo, 0, 8)

	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return nil, err
		}

		project, err := k.stateStore.ProjectTable().Get(ctx, batch.ProjectKey)
		if err != nil {
			return nil, err
		}

		info := types.BatchInfo{
			Issuer:       req.Issuer,
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
		return nil, err
	}

	return &types.QueryBatchesByIssuerResponse{Batches: batches, Pagination: pr}, nil
}
