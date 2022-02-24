package basket

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func (k Keeper) Basket(ctx context.Context, request *baskettypes.QueryBasketRequest) (*baskettypes.QueryBasketResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	basket, err := k.stateStore.BasketStore().GetByBasketDenom(ctx, request.BasketDenom)
	if err != nil {
		return nil, err
	}

	classes, err := k.stateStore.BasketBalanceStore().Get(ctx, basket.Id, basket.BasketDenom)
	if err != nil {
		return nil, err
	}

	basketGogo := &baskettypes.Basket{}
	err = PulsarToGogoSlow(basket, basketGogo)
	if err != nil {
		return nil, err
	}

	return &baskettypes.QueryBasketResponse{Basket: basketGogo}, nil
}
