package core

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Supply queries the supply (tradable, retired, cancelled) of a given credit batch.
func (k Keeper) Supply(ctx context.Context, request *core.QuerySupplyRequest) (*core.QuerySupplyResponse, error) {
	batch, err := k.stateStore.BatchInfoTable().GetByDenom(ctx, request.BatchDenom)
	if err != nil {
		return nil, err
	}

	supply, err := k.stateStore.BatchSupplyTable().Get(ctx, batch.Key)
	if err != nil {
		return nil, err
	}

	return &core.QuerySupplyResponse{
		TradableSupply:  supply.TradableAmount,
		RetiredSupply:   supply.RetiredAmount,
		CancelledAmount: supply.CancelledAmount,
	}, nil
}
