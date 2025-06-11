package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/basket/types/v1"
)

func (k Keeper) Baskets(ctx context.Context, req *types.QueryBasketsRequest) (*types.QueryBasketsResponse, error) {
	if req == nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap("empty request")
	}

	pg := ormutil.PageReqToOrmPaginate(req.Pagination)
	it, err := k.stateStore.BasketTable().List(ctx, api.BasketPrimaryKey{}, pg)
	if err != nil {
		return nil, err
	}
	defer it.Close()

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
	res.Pagination = ormutil.PageResToCosmosTypes(it.PageResponse())

	return res, nil
}
