package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
)

func (k Keeper) BasketBalances(ctx context.Context, request *types.QueryBasketBalancesRequest) (*types.QueryBasketBalancesResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	basket, err := k.stateStore.BasketTable().GetByBasketDenom(ctx, request.BasketDenom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return nil, errors.Wrapf(err, "basket %s not found", request.BasketDenom)
		}
		return nil, errors.Wrapf(err, "failed to get basket %s", request.BasketDenom)
	}

	pulsarPageReq, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
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
	it.Close()

	res.Pagination, err = ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	return res, err
}
