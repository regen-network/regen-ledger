package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/errors"
	"github.com/regen-network/regen-ledger/types/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
)

func (k Keeper) BasketBalances(ctx context.Context, request *types.QueryBasketBalancesRequest) (*types.QueryBasketBalancesResponse, error) {
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

	pulsarPageReq, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	it, err := k.stateStore.BasketBalanceTable().List(ctx, api.BasketBalancePrimaryKey{}.WithBasketId(basket.Id),
		ormlist.Paginate(pulsarPageReq),
	)
	if err != nil {
		return nil, err
	}

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
	defer it.Close()

	res.Pagination, err = ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}

	return res, nil
}
