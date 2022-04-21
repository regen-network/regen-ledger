package basket

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/math"
)

// GetBasketBalanceMap calculates credit balance of each batch within the basket
func (k Keeper) GetBasketBalanceMap(ctx context.Context) (map[uint64]math.Dec, error) {
	batchDenomToKey := make(map[string]uint64)     // map of a batch denom to batch key
	batchKeyToBalance := make(map[uint64]math.Dec) // map of a basket batch key to balance

	itr, err := k.stateStore.BasketBalanceTable().List(ctx, api.BasketBalancePrimaryKey{})
	if err != nil {
		return nil, err
	}
	defer itr.Close()

	for itr.Next() {
		bb, err := itr.Value()
		if err != nil {
			return nil, err
		}

		amount, err := math.NewDecFromString(bb.Balance)
		if err != nil {
			return nil, err
		}

		var batchKey uint64
		if _, ok := batchDenomToKey[bb.BatchDenom]; !ok {
			bInfo, err := k.coreStore.BatchInfoTable().GetByBatchDenom(ctx, bb.BatchDenom)
			if err != nil {
				return nil, err
			}

			batchDenomToKey[bb.BatchDenom] = bInfo.Key
			batchKey = bInfo.Key
		} else {
			batchKey = batchDenomToKey[bb.BatchDenom]
		}

		if existingBal, ok := batchKeyToBalance[batchKey]; ok {
			existingBal, err = existingBal.Add(amount)
			if err != nil {
				return nil, err
			}

			batchKeyToBalance[batchKey] = existingBal
		} else {
			batchKeyToBalance[batchKey] = amount
		}
	}

	return batchKeyToBalance, nil
}
