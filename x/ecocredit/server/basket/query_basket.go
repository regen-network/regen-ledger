package basket

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
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

	basketGogo := &baskettypes.Basket{}
	err = PulsarToGogoSlow(basket, basketGogo)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.BasketClassStore().List(ctx, basketv1.BasketClassPrimaryKey{}.WithBasketId(basket.Id))
	if err != nil {
		return nil, err
	}

	var classes []string
	for it.Next() {
		class, err := it.Value()
		if err != nil {
			return nil, err
		}

		classes = append(classes, class.ClassId)
	}

	it.Close()

	return &baskettypes.QueryBasketResponse{Basket: basketGogo, Classes: classes}, nil
}
