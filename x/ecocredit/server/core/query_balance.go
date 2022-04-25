package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Balance queries the balance (both tradable and retired) of a given credit
// batch for a given account.
func (k Keeper) Balance(ctx context.Context, req *core.QueryBalanceRequest) (*core.QueryBalanceResponse, error) {
	batch, err := k.stateStore.BatchTable().GetByDenom(ctx, req.BatchDenom)
	if err != nil {
		return nil, err
	}

	addr, err := sdk.AccAddressFromBech32(req.Account)
	if err != nil {
		return nil, err
	}

	balance, err := k.stateStore.BatchBalanceTable().Get(ctx, addr, batch.Key)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return &core.QueryBalanceResponse{
				Balance: &core.BatchBalanceInfo{
					Address:    addr.String(),
					BatchDenom: batch.Denom,
					Tradable:   "0",
					Retired:    "0",
					Escrowed:   "0",
				},
			}, nil
		}
		return nil, err
	}

	info := core.BatchBalanceInfo{
		Address:    addr.String(),
		BatchDenom: batch.Denom,
		Tradable:   balance.Tradable,
		Retired:    balance.Retired,
		Escrowed:   balance.Escrowed,
	}

	return &core.QueryBalanceResponse{Balance: &info}, nil
}
