package basket

import (
	"context"

	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func (k Keeper) Basket(ctx context.Context, request *baskettypes.QueryBasketRequest) (*baskettypes.QueryBasketResponse, error) {
	basket, err := k.stateStore.BasketStore().GetByBasketDenom(ctx, request.BasketDenom)
	if err != nil {
		return nil, err
	}

	return &baskettypes.QueryBasketResponse{Basket: basket}, nil
}
