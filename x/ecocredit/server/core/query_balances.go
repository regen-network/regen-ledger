package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func (k Keeper) Balances(ctx context.Context, req *core.QueryBalancesRequest) (*core.QueryBalancesResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.BatchBalanceTable().List(ctx, api.BatchBalanceAddressBatchKeyIndexKey{}.WithAddress(addr), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	balances := make([]*core.BatchBalanceInfo, 0, 8) // pre-allocate some cap space
	for it.Next() {
		balance, err := it.Value()
		if err != nil {
			return nil, err
		}

		batch, err := k.stateStore.BatchTable().Get(ctx, balance.BatchKey)

		info := core.BatchBalanceInfo{
			Address:        addr.String(),
			BatchDenom:     batch.Denom,
			TradableAmount: balance.TradableAmount,
			RetiredAmount:  balance.RetiredAmount,
			EscrowedAmount: balance.EscrowedAmount,
		}

		balances = append(balances, &info)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &core.QueryBalancesResponse{Balances: balances, Pagination: pr}, nil
}
