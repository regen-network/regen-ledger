package keeper

import (
	"context"
	"fmt"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/v2/math"
)

// BatchSupplyInvariant checks
// - calculated total tradable balance of each credit batch matches the total tradable supply
// - calculated total retired balance of each credit batch matches the total retired supply
func BatchSupplyInvariant(ctx context.Context, k Keeper, basketBalances map[uint64]math.Dec) (msg string, broken bool) {
	// sum of tradeable ecocredits with credits locked in baskets
	batchIDToBalanceTradable := make(map[uint64]math.Dec) // map batch id => tradable balance
	batchIDToBalanceRetired := make(map[uint64]math.Dec)  // map batch id => retired balance

	itr, err := k.stateStore.BatchBalanceTable().List(ctx, api.BatchBalancePrimaryKey{})
	if err != nil {
		return err.Error(), true
	}
	defer itr.Close()

	balanceFunc := func(name string, batchIDToBalance map[uint64]math.Dec, bID uint64, amount string) {
		amt, err := math.NewNonNegativeDecFromString(amount)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("error while parsing %s balance %v", name, err)
		}
		if val, ok := batchIDToBalance[bID]; ok {
			result, err := math.SafeAddBalance(val, amt)
			if err != nil {
				broken = true
				msg += fmt.Sprintf("error adding credit batch %s supply %v", name, err)
			}
			batchIDToBalance[bID] = result
		} else {
			batchIDToBalance[bID] = amt
		}
	}

	for itr.Next() {
		balance, err := itr.Value()
		if err != nil {
			return err.Error(), true
		}

		// tradable balance
		balanceFunc("tradable", batchIDToBalanceTradable, balance.BatchKey, balance.TradableAmount)

		// escrowed balance
		balanceFunc("tradable", batchIDToBalanceTradable, balance.BatchKey, balance.EscrowedAmount)

		// retired balance
		balanceFunc("retired", batchIDToBalanceRetired, balance.BatchKey, balance.RetiredAmount)
	}

	for batchID, amt := range basketBalances {
		if amount, ok := batchIDToBalanceTradable[batchID]; ok {
			result, err := math.SafeAddBalance(amount, amt)
			if err != nil {
				broken = true
				msg += fmt.Sprintf("\tfailed to add %v amount to %v for credit batch %d\n", amt, amount, batchID)
				return msg, broken
			}
			batchIDToBalanceTradable[batchID] = result
		} else {
			msg += fmt.Sprintf("\tunknown credit batch %d in basket", batchID)
			return msg, true
		}
	}

	supplyFunc := func(name string, batchIDToBalance map[uint64]math.Dec, batchKey uint64, amount string) {
		expected, err := math.NewNonNegativeDecFromString(amount)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("\terror while parsing %s supply for batch id: %d\n", name, batchKey)
		}
		if actual, ok := batchIDToBalance[batchKey]; ok {
			if expected.Cmp(actual) != math.EqualTo {
				broken = true
				msg += fmt.Sprintf("\t%s supply is incorrect for %d credit batch, expected %v, got %v\n", name, batchKey, expected, actual)
			}
		} else {
			broken = true
			msg += fmt.Sprintf("\t%s supply is not found for %d credit batch\n", name, batchKey)
		}
	}

	sItr, err := k.stateStore.BatchSupplyTable().List(ctx, api.BatchSupplyPrimaryKey{})
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
		supplyFunc("tradable", batchIDToBalanceTradable, supply.BatchKey, supply.TradableAmount)

		// retired supply invariant check
		supplyFunc("retired", batchIDToBalanceRetired, supply.BatchKey, supply.RetiredAmount)
	}

	return msg, broken
}
