package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/basket/types/v1"
)

func (k Keeper) BasketBalances(ctx context.Context, req *types.QueryBasketBalancesRequest) (*types.QueryBasketBalancesResponse, error) {
	if req == nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap("empty request")
	}

	basket, err := k.stateStore.BasketTable().GetByBasketDenom(ctx, req.BasketDenom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return nil, regenerrors.ErrNotFound.Wrapf("basket %s not found", req.BasketDenom)
		}
		return nil, regenerrors.ErrInternal.Wrapf("failed to get basket %s", req.BasketDenom)
	}

	pg := ormutil.PageReqToOrmPaginate(req.Pagination)
	it, err := k.stateStore.BasketBalanceTable().List(
		ctx, api.BasketBalancePrimaryKey{}.WithBasketId(basket.Id), pg)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	res := &types.QueryBasketBalancesResponse{}
	for it.Next() {
		bal, err := it.Value()
		if err != nil {
			return nil, err
		}

		balanceGogo := &types.BasketBalance{}
		if err = ormutil.PulsarToGogoSlow(bal, balanceGogo); err != nil {
			return nil, err
		}
		res.Balances = append(res.Balances, balanceGogo)
		res.BalancesInfo = append(res.BalancesInfo, &types.BasketBalanceInfo{
			BatchDenom: bal.BatchDenom,
			Balance:    bal.Balance,
		})
	}

	res.Pagination = ormutil.PageResToCosmosTypes(it.PageResponse())
	return res, nil
}
