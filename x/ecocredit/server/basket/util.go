package basket

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/math"
)

// GetBasketBalanceMap calculates credit balance of each batch within the basket
func (k Keeper) GetBasketBalanceMap(ctx context.Context) (map[uint64]math.Dec, error) {
	batchDenomToId := make(map[string]uint64)     // map of a batch denom to batch id
	batchIdToBalance := make(map[uint64]math.Dec) // map of a basket batch_id to balance

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

		var batchID uint64
		if _, ok := batchDenomToId[bb.BatchDenom]; !ok {
			bInfo, err := k.coreStore.BatchInfoTable().GetByBatchDenom(ctx, bb.BatchDenom)
			if err != nil {
				return nil, err
			}

			batchDenomToId[bb.BatchDenom] = bInfo.Id
			batchID = bInfo.Id
		} else {
			batchID = batchDenomToId[bb.BatchDenom]
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
