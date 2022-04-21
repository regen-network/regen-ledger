package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Balance queries the balance (both tradable and retired) of a given credit
// batch for a given account.
func (k Keeper) Balance(ctx context.Context, req *core.QueryBalanceRequest) (*core.QueryBalanceResponse, error) {
	batch, err := k.stateStore.BatchInfoTable().GetByDenom(ctx, req.BatchDenom)
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
				Balance: &core.BatchBalance{
					BatchKey: batch.Key,
					Address:  addr,
					Tradable: "0",
					Retired:  "0",
					Escrowed: "0",
				},
			}, nil
		}
		return nil, err
	}
	var bal core.BatchBalance
	if err = ormutil.PulsarToGogoSlow(balance, &bal); err != nil {
		return nil, err
	}

	return &core.QueryBalanceResponse{Balance: &bal}, nil
}
