package server

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/core"
)

// RegisterInvariants registers the ecocredit module invariants.
func (s serverImpl) RegisterInvariants(ir sdk.InvariantRegistry) {
	ir.RegisterRoute(ecocredit.ModuleName, "batch-supply", s.batchSupplyInvariant())
	s.basketKeeper.RegisterInvariants(ir)
}

func (s serverImpl) batchSupplyInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		goCtx := sdk.WrapSDKContext(ctx)
		basketBalances, err := s.basketKeeper.GetBasketBalanceMap(goCtx)
		if err != nil {
			return err.Error(), true
		}

		return BatchSupplyInvariant(goCtx, s.coreKeeper, basketBalances)
	}
}

func BatchSupplyInvariant(ctx context.Context, k core.Keeper, basketBalances map[uint64]math.Dec) (string, bool) {
	var (
		msg    string
		broken bool
	)
	// sum of tradeable eco credits with credits locked in baskets
	sumBatchSupplies := make(map[uint64]math.Dec) // map batch id => balance
	calRetiredSupplies := make(map[uint64]math.Dec)
	itr, err := k.BatchBalanceIterator(ctx)
	if err != nil {
		return err.Error(), true
	}
	defer itr.Close()

	for itr.Next() {
		bb, err := itr.Value()
		if err != nil {
			return err.Error(), true
		}

		// tradable balance
		balance, err := math.NewNonNegativeDecFromString(bb.Tradable)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("error while parsing tradable balance %v", err)
		}
		if supply, ok := sumBatchSupplies[bb.BatchId]; ok {
			supply, err := math.SafeAddBalance(supply, balance)
			if err != nil {
				broken = true
				msg += fmt.Sprintf("error adding credit batch tradable supply %v", err)
			}
			sumBatchSupplies[bb.BatchId] = supply
		} else {
			sumBatchSupplies[bb.BatchId] = balance
		}

		// retired balance
		balance, err = math.NewNonNegativeDecFromString(bb.Retired)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("error while parsing retired balance %v", err)
		}
		if supply, ok := calRetiredSupplies[bb.BatchId]; ok {
			supply, err := math.SafeAddBalance(balance, supply)
			if err != nil {
				broken = true
				msg += fmt.Sprintf("error adding credit batch retired supply %v", err)
			}
			calRetiredSupplies[bb.BatchId] = supply
		} else {
			calRetiredSupplies[bb.BatchId] = balance
		}
	}

	for id, amt := range basketBalances {
		if amount, ok := sumBatchSupplies[id]; ok {
			amount, err := math.SafeAddBalance(amount, amt)
			if err != nil {
				panic(err)
			}
			sumBatchSupplies[id] = amount
		} else {
			return "unknown denom in basket", true
		}
	}

	sItr, err := k.BatchSupplyIterator(ctx)
	if err != nil {
		return msg + err.Error(), true
	}
	defer sItr.Close()

	for sItr.Next() {
		bs, err := sItr.Value()
		if err != nil {
			return msg + err.Error(), true
		}

		// tradable supply invariant check
		tSupply, err := math.NewNonNegativeDecFromString(bs.TradableAmount)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("error while parsing tradable supply for denom: %d", bs.BatchId)
		}
		if s1, ok := sumBatchSupplies[bs.BatchId]; ok {
			if tSupply.Cmp(s1) != 0 {
				broken = true
				msg += fmt.Sprintf("tradable supply is incorrect for %d credit batch, expected %v, got %v", bs.BatchId, tSupply, s1)
			}
		} else {
			broken = true
			msg += fmt.Sprintf("tradable supply is not found for %d credit batch", bs.BatchId)
		}

		// retired supply invariant check
		supply, err := math.NewNonNegativeDecFromString(bs.RetiredAmount)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("error while parsing reired supply for denom: %d", bs.BatchId)
		}
		if s1, ok := calRetiredSupplies[bs.BatchId]; ok {
			if supply.Cmp(s1) != 0 {
				broken = true
				msg += fmt.Sprintf("retired supply is incorrect for %d credit batch, expected %v, got %v", bs.BatchId, supply, s1)
			}
		} else {
			broken = true
			msg += fmt.Sprintf("retired supply is not found for %d credit batch", bs.BatchId)
		}

	}

	return msg, broken
}
