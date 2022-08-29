package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// problem 1: cannot list by exact match.

// credit type abbrev
// credit class id
// batch denom

// return entries straight from ORM

func (k Keeper) AllBalances(ctx context.Context, req *core.QueryAllBalancesRequest) (*core.QueryAllBalancesResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.BatchBalanceTable().List(ctx, api.BatchBalancePrimaryKey{}, ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	var res core.QueryAllBalancesResponse
	for it.Next() {
		balance, err := it.Value()
		if err != nil {
			return nil, err
		}

		batch, err := k.stateStore.BatchTable().Get(ctx, balance.BatchKey)
		if err != nil {
			return nil, err
		}

		res.Balances = append(res.Balances, &core.BatchBalanceInfo{
			Address:        sdk.AccAddress(balance.Address).String(),
			BatchDenom:     batch.Denom,
			TradableAmount: balance.TradableAmount,
			RetiredAmount:  balance.RetiredAmount,
			EscrowedAmount: balance.EscrowedAmount,
		})
	}
	res.Pagination, err = ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &res, nil
}
