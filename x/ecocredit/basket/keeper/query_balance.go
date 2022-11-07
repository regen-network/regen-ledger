package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	regenerrors "github.com/regen-network/regen-ledger/errors"

	types "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
)

func (k Keeper) BasketBalance(ctx context.Context, request *types.QueryBasketBalanceRequest) (*types.QueryBasketBalanceResponse, error) {
	if request == nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap("empty request")
	}

	basket, err := k.stateStore.BasketTable().GetByBasketDenom(ctx, request.BasketDenom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return nil, regenerrors.ErrNotFound.Wrapf("basket %s not found", request.BasketDenom)
		}
		return nil, regenerrors.ErrInternal.Wrapf("failed to get basket %s", request.BasketDenom)
	}

	found, err := k.baseStore.BatchTable().HasByDenom(ctx, request.BatchDenom)
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrapf("failed to get credit batch %s", request.BatchDenom)
	}
	if !found {
		return nil, regenerrors.ErrNotFound.Wrapf("credit batch %s not found", request.BatchDenom)
	}

	balance, err := k.stateStore.BasketBalanceTable().Get(ctx, basket.Id, request.BatchDenom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return &types.QueryBasketBalanceResponse{Balance: "0"}, nil
		}
		return nil, regenerrors.ErrInternal.Wrapf("failed to get basket balance for %s", request.BasketDenom)
	}

	return &types.QueryBasketBalanceResponse{Balance: balance.Balance}, nil
}
