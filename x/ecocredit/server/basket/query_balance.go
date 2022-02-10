package basket

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func (k Keeper) BasketBalance(ctx context.Context, request *baskettypes.QueryBasketBalanceRequest) (*baskettypes.QueryBasketBalanceResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	basket, err := k.stateStore.BasketStore().GetByBasketDenom(ctx, request.BasketDenom)
	if err != nil {
		return nil, err
	}

	balance, err := k.stateStore.BasketBalanceStore().Get(ctx, basket.Id, request.BatchDenom)
	if err != nil {
		return nil, err
	}

	return &baskettypes.QueryBasketBalanceResponse{Balance: balance.Balance}, nil
}
