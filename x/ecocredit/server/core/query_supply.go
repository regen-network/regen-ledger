package core

import (
	"context"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
)

// Supply queries the supply (tradable, retired, cancelled) of a given credit batch.
func (k Keeper) Supply(ctx context.Context, request *v1.QuerySupplyRequest) (*v1.QuerySupplyResponse, error) {
	batch, err := k.stateStore.BatchInfoStore().GetByBatchDenom(ctx, request.BatchDenom)
	if err != nil {
		return nil, err
	}

	supply, err := k.stateStore.BatchSupplyStore().Get(ctx, batch.Id)
	if err != nil {
		return nil, err
	}

	return &v1.QuerySupplyResponse{
		TradableSupply:  supply.TradableAmount,
		RetiredSupply:   supply.RetiredAmount,
		CancelledAmount: supply.CancelledAmount,
	}, nil
}
