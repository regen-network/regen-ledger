package core

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// RegisterInvariants registers the ecocredit module invariants.
func (s serverImpl) RegisterInvariants(ir sdk.InvariantRegistry) {
	ir.RegisterRoute(ecocredit.ModuleName, "tradable-supply", s.tradableSupplyInvariant())
	ir.RegisterRoute(ecocredit.ModuleName, "retired-supply", s.retiredSupplyInvariant())
}

func (s serverImpl) tradableSupplyInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		supplyStore := s.batchSupplyStore
		batchStore := s.batchInfoStore
		balanceStore := s.batchBalanceStore
		return tradableSupplyInvariant(supplyStore, balanceStore, batchStore, ctx)
	}
}

func tradableSupplyInvariant(batchSupplyStore ecocreditv1beta1.BatchSupplyStore, balanceStore ecocreditv1beta1.BatchBalanceStore, batchInfoStore ecocreditv1beta1.BatchInfoStore, ctx sdk.Context) (string, bool) {
	var (
		msg    string
		broken bool
	)
	calTradableSupplies := make(map[string]math.Dec)

	err := iterateBalances(balanceStore, batchInfoStore, true, ctx, func(denom, supply string) bool {
		balance, err := math.NewNonNegativeDecFromString(supply)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("error while parsing tradable balance %v", err)
		}
		if supply, ok := calTradableSupplies[denom]; ok {
			supply, err := math.SafeAddBalance(supply, balance)
			if err != nil {
				broken = true
				msg += fmt.Sprintf("error adding credit batch tradable supply %v", err)
			}
			calTradableSupplies[denom] = supply
		} else {
			calTradableSupplies[denom] = balance
		}

		return false
	})
	if err != nil {
		msg = fmt.Sprintf("error querying credit balances tradable supply %v", err)
	}

	if err := iterateSupplies(batchSupplyStore, batchInfoStore, true, ctx, func(denom string, s string) bool {
		supply, err := math.NewNonNegativeDecFromString(s)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("error while parsing tradable supply for denom: %s", denom)
		}
		if s1, ok := calTradableSupplies[denom]; ok {
			if supply.Cmp(s1) != 0 {
				broken = true
				msg += fmt.Sprintf("tradable supply is incorrect for %s credit batch, expected %v, got %v", denom, supply, s1)
			}
		} else {
			broken = true
			msg += fmt.Sprintf("tradable supply is not found for %s credit batch", denom)
		}
		return false
	}); err != nil {
		msg = fmt.Sprintf("error querying credit tradable supply %v", err)
	}

	return msg, broken
}

func (s serverImpl) retiredSupplyInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		supplyStore := s.batchSupplyStore
		balanceStore := s.batchBalanceStore
		batchStore := s.batchInfoStore
		return retiredSupplyInvariant(supplyStore, balanceStore, batchStore, ctx)
	}
}

func retiredSupplyInvariant(supplyStore ecocreditv1beta1.BatchSupplyStore, balanceStore ecocreditv1beta1.BatchBalanceStore, batchInfoStore ecocreditv1beta1.BatchInfoStore, ctx sdk.Context) (string, bool) {
	var (
		msg    string
		broken bool
	)
	calRetiredSupplies := make(map[string]math.Dec)
	err := iterateBalances(balanceStore, batchInfoStore, false, ctx, func(denom, supply string) bool {
		balance, err := math.NewNonNegativeDecFromString(supply)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("error while parsing retired balance %v", err)
		}
		if supply, ok := calRetiredSupplies[denom]; ok {
			supply, err := math.SafeAddBalance(balance, supply)
			if err != nil {
				broken = true
				msg += fmt.Sprintf("error adding credit batch retired supply %v", err)
			}
			calRetiredSupplies[denom] = supply
		} else {
			calRetiredSupplies[denom] = balance
		}
		return false
	})
	if err != nil {
		msg = fmt.Sprintf("error querying credit retired balances %v", err)
	}

	if err := iterateSupplies(supplyStore, batchInfoStore, false, ctx, func(denom, s string) bool {
		supply, err := math.NewNonNegativeDecFromString(s)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("error while parsing reired supply for denom: %s", denom)
		}
		if s1, ok := calRetiredSupplies[denom]; ok {
			if supply.Cmp(s1) != 0 {
				broken = true
				msg += fmt.Sprintf("retired supply is incorrect for %s credit batch, expected %v, got %v", denom, supply, s1)
			}
		} else {
			broken = true
			msg += fmt.Sprintf("retired supply is not found for %s credit batch", denom)
		}

		return false
	}); err != nil {
		msg = fmt.Sprintf("error querying credit retired supply %v", err)
	}

	return msg, broken
}

func iterateBalances(store ecocreditv1beta1.BatchBalanceStore, batchStore ecocreditv1beta1.BatchInfoStore, tradable bool, sdkCtx sdk.Context, cb func(denom, balance string) bool) error {
	ctx := sdkCtx.Context()
	it, err := store.List(ctx, &ecocreditv1beta1.BatchBalanceBatchIdAddressIndexKey{})
	if err != nil {
		return err
	}
	// we cache denoms here so we don't have to ORM query each time
	batchIdToDenom := make(map[uint64]string)
	for it.Next() {
		val, err := it.Value()
		if err != nil {
			return err
		}
		var batchDenom string
		if d, ok := batchIdToDenom[val.BatchId]; ok {
			batchDenom = d
		} else {
			batch, err := batchStore.Get(ctx, val.BatchId)
			if err != nil {
				return err
			}
			if batch == nil {
				panic(fmt.Sprintf("batch was nil with batch id: %d and denom %s", val.BatchId, batchDenom))
			}
			batchDenom = batch.BatchDenom
			batchIdToDenom[val.BatchId] = batchDenom
		}

		var balance string
		if tradable {
			balance = val.Tradable
		} else {
			balance = val.Retired
		}

		if cb(batchDenom, balance) {
			break
		}
	}
	return nil
}

func iterateSupplies(supplyStore ecocreditv1beta1.BatchSupplyStore, batchStore ecocreditv1beta1.BatchInfoStore, tradable bool, sdkCtx sdk.Context, cb func(denom, supply string) bool) error {
	ctx := sdkCtx.Context()
	it, err := supplyStore.List(ctx, &ecocreditv1beta1.BatchSupplyBatchIdIndexKey{})
	if err != nil {
		return err
	}
	// we cache denoms here so we don't have to ORM query each time
	batchIdToDenom := make(map[uint64]string)
	for it.Next() {
		val, err := it.Value()
		if err != nil {
			return err
		}
		var batchDenom string
		if d, ok := batchIdToDenom[val.BatchId]; ok {
			batchDenom = d
		} else {
			batch, err := batchStore.Get(ctx, val.BatchId)
			if err != nil {
				return err
			}
			batchDenom = batch.BatchDenom
			batchIdToDenom[val.BatchId] = batchDenom
		}
		var supply string
		if tradable {
			supply = val.TradableAmount
		} else {
			supply = val.RetiredAmount
		}
		if cb(batchDenom, supply) {
			break
		}
	}
	return nil
}
