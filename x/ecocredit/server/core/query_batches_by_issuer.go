package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// BatchesByIssuer queries all batches issued from the given issuer address
func (k Keeper) BatchesByIssuer(ctx context.Context, req *core.QueryBatchesByIssuerRequest) (*core.QueryBatchesByIssuerResponse, error) {
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

	batches := make([]*core.BatchInfo, 0, 8)

	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return nil, err
		}

		project, err := k.stateStore.ProjectTable().Get(ctx, batch.ProjectKey)
		if err != nil {
			return nil, err
		}

		info := core.BatchInfo{
			Issuer:       req.Issuer,
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

	return &core.QueryBatchesByIssuerResponse{Batches: batches, Pagination: pr}, nil
}
