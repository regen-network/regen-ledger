package core

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Balances(ctx context.Context, req *core.QueryBalancesRequest) (*core.QueryBalancesResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Account)
	if err != nil {
		return nil, err
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.BatchBalanceTable().List(ctx, api.BatchBalanceAddressBatchIdIndexKey{}.WithAddress(addr), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	balances := make([]*core.BatchBalanceEntry, 0, 8) // pre-allocate some cap space
	for it.Next() {
		balance, err := it.Value()
		if err != nil {
			return nil, err
		}

		batch, err := k.stateStore.BatchInfoTable().Get(ctx, balance.BatchId)

		entry := core.BatchBalanceEntry{
			Address:    addr.String(),
			BatchDenom: batch.BatchDenom,
			Tradable:   balance.Tradable,
			Retired:    balance.Retired,
			Escrowed:   balance.Escrowed,
		}

		balances = append(balances, &entry)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &core.QueryBalancesResponse{Balances: balances, Pagination: pr}, nil
}
