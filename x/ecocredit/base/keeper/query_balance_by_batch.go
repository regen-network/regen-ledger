package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

func (k Keeper) BalancesByBatch(ctx context.Context, req *types.QueryBalancesByBatchRequest) (*types.QueryBalancesByBatchResponse, error) {
	batch, err := k.stateStore.BatchTable().GetByDenom(ctx, req.BatchDenom)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrapf("could not get batch with denom %s: %s", req.BatchDenom, err.Error())
	}
	pg := ormutil.PageReqToOrmPaginate(req.Pagination)
	it, err := k.stateStore.BatchBalanceTable().List(ctx, api.BatchBalanceBatchKeyAddressIndexKey{}.WithBatchKey(batch.Key), pg)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	balances := make([]*types.BatchBalanceInfo, 0, 10) // preallocate
	for it.Next() {
		bal, err := it.Value()
		if err != nil {
			return nil, err
		}
		balances = append(balances, &types.BatchBalanceInfo{
			Address:        sdk.AccAddress(bal.Address).String(),
			BatchDenom:     batch.Denom,
			TradableAmount: bal.TradableAmount,
			RetiredAmount:  bal.RetiredAmount,
			EscrowedAmount: bal.EscrowedAmount,
		})
	}
	return &types.QueryBalancesByBatchResponse{
		Balances:   balances,
		Pagination: ormutil.PageResToCosmosTypes(it.PageResponse()),
	}, nil
}
