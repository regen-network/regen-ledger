package basket

import (
	"context"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/math"

	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func (k Keeper) Take(ctx context.Context, msg *baskettypes.MsgTake) (*baskettypes.MsgTakeResponse, error) {
	basket, err := k.stateStore.BasketStore().GetByBasketDenom(ctx, msg.BasketDenom)
	if err != nil {
		return nil, err
	}

	retire := msg.RetireOnTake
	if !basket.DisableAutoRetire && !retire {
		return nil, ErrCantDisableRetire
	}

	it, err := k.stateStore.BasketBalanceStore().List(ctx, basketv1.BasketBalanceBatchStartDateIndexKey{})
	if err != nil {
		return nil, err
	}

	amount, err := math.NewDecFromString(msg.Amount)
	if err != nil {
		return nil, err
	}

	for it.Next() {

	}
}
