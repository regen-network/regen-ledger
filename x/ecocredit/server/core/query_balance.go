package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
)

// Balance queries the balance (both tradable and retired) of a given credit
// batch for a given account.
func (k Keeper) Balance(ctx context.Context, req *v1.QueryBalanceRequest) (*v1.QueryBalanceResponse, error) {
	batch, err := k.stateStore.BatchInfoStore().GetByBatchDenom(ctx, req.BatchDenom)
	if err != nil {
		return nil, err
	}
	addr, err := sdk.AccAddressFromBech32(req.Account)
	if err != nil {
		return nil, err
	}

	balance, err := k.stateStore.BatchBalanceStore().Get(ctx, addr, batch.Id)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return &v1.QueryBalanceResponse{
				TradableAmount: "0",
				RetiredAmount:  "0",
			}, nil
		}
		return nil, err
	}
	return &v1.QueryBalanceResponse{
		TradableAmount: balance.Tradable,
		RetiredAmount:  balance.Retired,
	}, nil
}
