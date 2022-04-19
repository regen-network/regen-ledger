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
	acc, err := sdk.AccAddressFromBech32(req.Account)
	if err != nil {
		return nil, err
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.BatchBalanceTable().List(ctx, api.BatchBalanceAddressBatchKeyIndexKey{}.WithAddress(acc), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	balances := make([]*core.BatchBalance, 0, 8) // pre-allocate some cap space
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err
		}
		var bal core.BatchBalance
		if err = ormutil.PulsarToGogoSlow(v, &bal); err != nil {
			return nil, err
		}
		balances = append(balances, &bal)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &core.QueryBalancesResponse{Balances: balances, Pagination: pr}, nil
}
