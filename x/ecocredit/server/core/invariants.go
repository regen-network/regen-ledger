package core

import (
	"context"
	"fmt"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
)

// BatchSupplyInvariant checks
// - calculated total tradable balance of each credit batch matches the total tradable supply
// - calculated total retired balance of each credit batch matches the total retired supply
func BatchSupplyInvariant(ctx context.Context, k Keeper, basketBalances map[uint64]math.Dec) (msg string, broken bool) {
	// sum of tradeable ecocredits with credits locked in baskets
	batchIdToBalanceTradable := make(map[uint64]math.Dec) // map batch id => tradable balance
	batchIdToBalanceRetired := make(map[uint64]math.Dec)  // map batch id => retired balance

	itr, err := k.stateStore.BatchBalanceTable().List(ctx, ecocreditv1.BatchBalancePrimaryKey{})
	if err != nil {
		return err.Error(), true
	}
	defer itr.Close()

	balanceFunc := func(name string, batchIdToBalance map[uint64]math.Dec, bID uint64, amount string) {
		amt, err := math.NewNonNegativeDecFromString(amount)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("error while parsing %s balance %v", name, err)
		}
		if val, ok := batchIdToBalance[bID]; ok {
			result, err := math.SafeAddBalance(val, amt)
			if err != nil {
				broken = true
				msg += fmt.Sprintf("error adding credit batch %s supply %v", name, err)
			}
			batchIdToBalance[bID] = result
		} else {
			batchIdToBalance[bID] = amt
		}
	}

	for itr.Next() {
		bBalance, err := itr.Value()
		if err != nil {
			return err.Error(), true
		}

		// tradable balance
		balanceFunc("tradable", batchIdToBalanceTradable, bBalance.BatchKey, bBalance.Tradable)

		//escrowed balance
		balanceFunc("tradable", batchIdToBalanceTradable, bBalance.BatchKey, bBalance.Escrowed)

		// retired balance
		balanceFunc("retired", batchIdToBalanceRetired, bBalance.BatchKey, bBalance.Retired)
	}

	for batchId, amt := range basketBalances {
		if amount, ok := batchIdToBalanceTradable[batchId]; ok {
			result, err := math.SafeAddBalance(amount, amt)
			if err != nil {
				broken = true
				msg += fmt.Sprintf("\tfailed to add %v amount to %v for credit batch %d\n", amt, amount, batchId)
				return msg, broken
			}
			batchIdToBalanceTradable[batchId] = result
		} else {
			msg += fmt.Sprintf("\tunknown credit batch %d in basket", batchId)
			return msg, true
		}
	}

	supplyFunc := func(name string, batchIdToBalance map[uint64]math.Dec, batchKey uint64, amount string) {
		expected, err := math.NewNonNegativeDecFromString(amount)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("\terror while parsing %s supply for batch id: %d\n", name, batchKey)
		}
		if actual, ok := batchIdToBalance[batchKey]; ok {
			if expected.Cmp(actual) != math.EqualTo {
				broken = true
				msg += fmt.Sprintf("\t%s supply is incorrect for %d credit batch, expected %v, got %v\n", name, batchKey, expected, actual)
			}
		} else {
			broken = true
			msg += fmt.Sprintf("\t%s supply is not found for %d credit batch\n", name, batchKey)
		}
	}

	sItr, err := k.stateStore.BatchSupplyTable().List(ctx, ecocreditv1.BatchSupplyPrimaryKey{})
	if err != nil {
		return msg + err.Error(), true
	}
	defer sItr.Close()

	for sItr.Next() {
		supply, err := sItr.Value()
		if err != nil {
			return err.Error(), true
		}

		// tradable supply invariant check
		supplyFunc("tradable", batchIdToBalanceTradable, supply.BatchKey, supply.TradableAmount)

		// retired supply invariant check
		supplyFunc("retired", batchIdToBalanceRetired, supply.BatchKey, supply.RetiredAmount)
	}

	return msg, broken
}
