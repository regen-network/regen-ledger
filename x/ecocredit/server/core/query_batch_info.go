package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// BatchInfo queries for information on a credit batch.
func (k Keeper) BatchInfo(ctx context.Context, request *core.QueryBatchInfoRequest) (*core.QueryBatchInfoResponse, error) {
	if err := ecocredit.ValidateDenom(request.BatchDenom); err != nil {
		return nil, err
	}

	batch, err := k.stateStore.BatchTable().GetByDenom(ctx, request.BatchDenom)
	if err != nil {
		return nil, err
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

	return &core.QueryBatchInfoResponse{Batch: &info}, nil
}
