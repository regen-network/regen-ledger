package basket

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func (k Keeper) Baskets(ctx context.Context, request *baskettypes.QueryBasketsRequest) (*baskettypes.QueryBasketsResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	pulsarPageReq, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.BasketTable().List(ctx, api.BasketPrimaryKey{},
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

		var criteria *baskettypes.DateCriteria
		if basket.DateCriteria != nil {
			criteria = &baskettypes.DateCriteria{}
			if err := ormutil.PulsarToGogoSlow(basket.DateCriteria, criteria); err != nil {
				return nil, err
			}
		}

		res.Baskets = append(res.Baskets, &baskettypes.BasketInfo{
			BasketDenom:       basket.BasketDenom,
			Name:              basket.Name,
			DisableAutoRetire: basket.DisableAutoRetire,
			CreditTypeAbbrev:  basket.CreditTypeAbbrev,
			Exponent:          basket.Exponent,
			Curator:           sdk.AccAddress(basket.Curator).String(),
			DateCriteria:      criteria,
		})
	}

	it.Close()

	res.Pagination, err = ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	return res, err
}
