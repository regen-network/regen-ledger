package core

import (
	"context"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	it, err := k.stateStore.BatchInfoTable().List(ctx, ecocreditv1.BatchInfoIssuerIndexKey{}.WithIssuer(issuer), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	batches := make([]*core.BatchInfoEntry, 0, 8)

	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return nil, err
		}

		project, err := k.stateStore.ProjectInfoTable().Get(ctx, batch.ProjectId)
		if err != nil {
			return nil, err
		}

		entry := core.BatchInfoEntry{
			Issuer:       req.Issuer,
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

	return &core.QueryBatchesByIssuerResponse{Batches: batches, Pagination: pr}, nil
}
