package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Batch queries for information on a credit batch.
func (k Keeper) Batch(ctx context.Context, request *core.QueryBatchRequest) (*core.QueryBatchResponse, error) {
	if err := core.ValidateBatchDenom(request.BatchDenom); err != nil {
		return nil, err
	}

	batch, err := k.stateStore.BatchTable().GetByDenom(ctx, request.BatchDenom)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get batch with denom %s: %s", request.BatchDenom, err.Error())
	}

	issuer := sdk.AccAddress(batch.Issuer)

	project, err := k.stateStore.ProjectTable().Get(ctx, batch.ProjectKey)
	if err != nil {
		return nil, err
	}

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

	return &core.QueryBatchResponse{Batch: &info}, nil
}
