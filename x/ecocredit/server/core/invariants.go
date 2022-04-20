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

	for itr.Next() {
		bBalance, err := itr.Value()
		if err != nil {
			return err.Error(), true
		}

		// tradable balance
		tBalance, err := math.NewNonNegativeDecFromString(bBalance.Tradable)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("error while parsing tradable balance %v", err)
		}
		if val, ok := batchIdToBalanceTradable[bBalance.BatchId]; ok {
			result, err := math.SafeAddBalance(val, tBalance)
			if err != nil {
				broken = true
				msg += fmt.Sprintf("error adding credit batch tradable supply %v", err)
			}
			batchIdToBalanceTradable[bBalance.BatchId] = result
		} else {
			batchIdToBalanceTradable[bBalance.BatchId] = tBalance
		}

		//escrowed balance
		eBalance, err := math.NewNonNegativeDecFromString(bBalance.Escrowed)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("error while parsing escrowed balance %v", err)
		}
		if val, ok := batchIdToBalanceTradable[bBalance.BatchId]; ok {
			result, err := math.SafeAddBalance(val, eBalance)
			if err != nil {
				broken = true
				msg += fmt.Sprintf("error adding credit batch tradable supply %v", err)
			}
			batchIdToBalanceTradable[bBalance.BatchId] = result
		} else {
			batchIdToBalanceTradable[bBalance.BatchId] = eBalance
		}

		// retired balance
		rBalance, err := math.NewNonNegativeDecFromString(bBalance.Retired)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("error while parsing retired balance %v", err)
		}
		if val, ok := batchIdToBalanceRetired[bBalance.BatchId]; ok {
			result, err := math.SafeAddBalance(val, rBalance)
			if err != nil {
				broken = true
				msg += fmt.Sprintf("error adding credit batch retired supply %v", err)
			}
			batchIdToBalanceRetired[bBalance.BatchId] = result
		} else {
			batchIdToBalanceRetired[bBalance.BatchId] = rBalance
		}
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
		tSupply, err := math.NewNonNegativeDecFromString(supply.TradableAmount)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("\terror while parsing tradable supply for batch id: %d\n", supply.BatchId)
		}
		if s1, ok := batchIdToBalanceTradable[supply.BatchId]; ok {
			if tSupply.Cmp(s1) != math.EqualTo {
				broken = true
				msg += fmt.Sprintf("\ttradable supply is incorrect for %d credit batch, expected %v, got %v\n", supply.BatchId, tSupply, s1)
			}
		} else {
			broken = true
			msg += fmt.Sprintf("\ttradable supply is not found for %d credit batch\n", supply.BatchId)
		}

		// retired supply invariant check
		retired, err := math.NewNonNegativeDecFromString(supply.RetiredAmount)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("\nerror while parsing reired supply for denom: %d\n", supply.BatchId)
		}
		if s1, ok := batchIdToBalanceRetired[supply.BatchId]; ok {
			if retired.Cmp(s1) != math.EqualTo {
				broken = true
				msg += fmt.Sprintf("\tretired supply is incorrect for %d credit batch, expected %v, got %v\n", supply.BatchId, retired, s1)
			}
		} else {
			broken = true
			msg += fmt.Sprintf("\tretired supply is not found for %d credit batch\n", supply.BatchId)
		}

	}

	return msg, broken
}
