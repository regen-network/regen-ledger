package basket

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func (k Keeper) Baskets(ctx context.Context, request *baskettypes.QueryBasketsRequest) (*baskettypes.QueryBasketsResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	pulsarPageReq, err := GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.BasketStore().List(ctx, basketv1.BasketPrimaryKey{},
		ormlist.Paginate(pulsarPageReq),
	)

	if err != nil {
		return nil, err
	}

	res := &baskettypes.QueryBasketsResponse{}
	for it.Next() {
		basket, err := it.Value()
		if err != nil {
			return nil, err
		}

		basketGogo := &baskettypes.Basket{}
		err = PulsarToGogoSlow(basket, basketGogo)
		if err != nil {
			return nil, err
		}

		res.Baskets = append(res.Baskets, basketGogo)
	}

	it.Close()

	res.Pagination, err = PulsarPageResToGogoPageRes(it.PageResponse())
	return res, err
}
