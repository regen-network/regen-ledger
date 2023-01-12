package keeper

import (
	"context"

	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// Supply queries the supply (tradable, retired, cancelled) of a given credit batch.
func (k Keeper) Supply(ctx context.Context, request *types.QuerySupplyRequest) (*types.QuerySupplyResponse, error) {
	batch, err := k.stateStore.BatchTable().GetByDenom(ctx, request.BatchDenom)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrapf("could not get batch with denom %s: %s", request.BatchDenom, err.Error())
	}

	supply, err := k.stateStore.BatchSupplyTable().Get(ctx, batch.Key)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrapf("unable to get batch supply for batch: %s", batch.Denom)
	}

	return &types.QuerySupplyResponse{
		TradableAmount:  supply.TradableAmount,
		RetiredAmount:   supply.RetiredAmount,
		CancelledAmount: supply.CancelledAmount,
	}, nil
}
