package basket

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/math"
)

func (k Keeper) GetBasketBalanceMap(ctx context.Context) (map[uint64]math.Dec, error) {
	itr, err := k.stateStore.BasketTable().List(ctx, api.BasketPrimaryKey{})
	if err != nil {
		return nil, err
	}
	defer itr.Close()

	batchDenomToId := make(map[string]uint64)     // map of a batch denom to batch id
	batchIdToBalance := make(map[uint64]math.Dec) // map of a basket batch_id to balance
	for itr.Next() {
		basket, err := itr.Value()
		if err != nil {
			return nil, err
		}

		bb, err := k.stateStore.BasketBalanceTable().Get(ctx, basket.Id, basket.BasketDenom)
		if err != nil {
			return nil, err
		}

		amount, err := math.NewDecFromString(bb.Balance)
		if err != nil {
			return nil, err
		}

		var batchID uint64
		if _, ok := batchDenomToId[basket.BasketDenom]; !ok {
			bInfo, err := k.coreStore.BatchInfoTable().GetByBatchDenom(ctx, basket.BasketDenom)
			if err != nil {
				return nil, err
			}
			batchDenomToId[basket.BasketDenom] = bInfo.Id
		} else {
			batchID = batchDenomToId[basket.BasketDenom]
		}

		if existingBal, ok := batchIdToBalance[batchID]; ok {
			existingBal, err = existingBal.Add(amount)
			if err != nil {
				return nil, err
			}

			batchIdToBalance[batchID] = existingBal
		} else {
			batchIdToBalance[batchID] = amount
		}
	}

	return batchIdToBalance, nil
}
