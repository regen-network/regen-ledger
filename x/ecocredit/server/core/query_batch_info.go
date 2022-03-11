package core

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// BatchInfo queries for information on a credit batch.
func (k Keeper) BatchInfo(ctx context.Context, request *core.QueryBatchInfoRequest) (*core.QueryBatchInfoResponse, error) {
	if err := ecocredit.ValidateDenom(request.BatchDenom); err != nil {
		return nil, err
	}

	batch, err := k.stateStore.BatchInfoStore().GetByBatchDenom(ctx, request.BatchDenom)
	if err != nil {
		return nil, err
	}

	var bi core.BatchInfo
	if err = PulsarToGogoSlow(batch, &bi); err != nil {
		return nil, err
	}
	return &core.QueryBatchInfoResponse{Info: &bi}, nil
}
