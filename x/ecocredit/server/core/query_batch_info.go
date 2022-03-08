package core

import (
	"context"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
)

// BatchInfo queries for information on a credit batch.
func (k Keeper) BatchInfo(ctx context.Context, request *v1.QueryBatchInfoRequest) (*v1.QueryBatchInfoResponse, error) {
	if err := ecocredit.ValidateDenom(request.BatchDenom); err != nil {
		return nil, err
	}

	batch, err := k.stateStore.BatchInfoStore().GetByBatchDenom(ctx, request.BatchDenom)
	if err != nil {
		return nil, err
	}

	var bi v1.BatchInfo
	if err = PulsarToGogoSlow(batch, &bi); err != nil {
		return nil, err
	}
	return &v1.QueryBatchInfoResponse{Info: &bi}, nil
}
