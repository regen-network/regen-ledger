package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func (k Keeper) BalancesByBatch(ctx context.Context, req *types.QueryBalancesByBatchRequest) (*types.QueryBalancesByBatchResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}
	batch, err := k.stateStore.BatchTable().GetByDenom(ctx, req.BatchDenom)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrapf("could not get batch with denom %s: %s", req.BatchDenom, err.Error())
	}
	it, err := k.stateStore.BatchBalanceTable().List(ctx, api.BatchBalanceBatchKeyAddressIndexKey{}.WithBatchKey(batch.Key), ormlist.Paginate(pg))
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
	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}
	return &types.QueryBalancesByBatchResponse{
		Balances:   balances,
		Pagination: pr,
	}, nil
}
