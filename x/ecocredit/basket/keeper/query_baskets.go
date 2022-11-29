package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
)

func (k Keeper) Baskets(ctx context.Context, request *types.QueryBasketsRequest) (*types.QueryBasketsResponse, error) {
	if request == nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap("empty request")
	}

	pulsarPageReq, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}
	it, err := k.stateStore.BasketTable().List(ctx, api.BasketPrimaryKey{},
		ormlist.Paginate(pulsarPageReq),
	)

	if err != nil {
		return nil, err
	}

	res := &types.QueryBasketsResponse{}
	for it.Next() {
		basket, err := it.Value()
		if err != nil {
			return nil, err
		}

		basketGogo := &types.Basket{}
		err = ormutil.PulsarToGogoSlow(basket, basketGogo)
		if err != nil {
			return nil, err
		}

		res.Baskets = append(res.Baskets, basketGogo)
		var criteria *types.DateCriteria
		if basket.DateCriteria != nil {
			criteria = &types.DateCriteria{}
			if err := ormutil.PulsarToGogoSlow(basket.DateCriteria, criteria); err != nil {
				return nil, err
			}
		}

		res.BasketsInfo = append(res.BasketsInfo, &types.BasketInfo{
			BasketDenom:       basket.BasketDenom,
			Name:              basket.Name,
			DisableAutoRetire: basket.DisableAutoRetire,
			CreditTypeAbbrev:  basket.CreditTypeAbbrev,
			Exponent:          basket.Exponent, //nolint:staticcheck
			Curator:           sdk.AccAddress(basket.Curator).String(),
			DateCriteria:      criteria,
		})
	}

	it.Close()

	res.Pagination, err = ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}

	return res, nil
}
