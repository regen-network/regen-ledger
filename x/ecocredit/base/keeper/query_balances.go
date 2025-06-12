package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

func (k Keeper) Balances(ctx context.Context, req *types.QueryBalancesRequest) (*types.QueryBalancesResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	pg := ormutil.PageReqToOrmPaginate(req.Pagination)
	it, err := k.stateStore.BatchBalanceTable().List(ctx, api.BatchBalanceAddressBatchKeyIndexKey{}.WithAddress(addr), pg)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	balances := make([]*types.BatchBalanceInfo, 0, 8) // pre-allocate some cap space
	for it.Next() {
		balance, err := it.Value()
		if err != nil {
			return nil, err
		}

		batch, err := k.stateStore.BatchTable().Get(ctx, balance.BatchKey)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("balance with key: %d", balance.BatchKey)
		}

		info := types.BatchBalanceInfo{
			Address:        addr.String(),
			BatchDenom:     batch.Denom,
			TradableAmount: balance.TradableAmount,
			RetiredAmount:  balance.RetiredAmount,
			EscrowedAmount: balance.EscrowedAmount,
		}

		balances = append(balances, &info)
	}

	pr := ormutil.PageResToCosmosTypes(it.PageResponse())
	return &types.QueryBalancesResponse{Balances: balances, Pagination: pr}, nil
}
