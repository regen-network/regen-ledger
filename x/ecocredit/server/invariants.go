package server

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		store := ctx.KVStore(s.storeKey)
		return tradableSupplyInvariant(store)
	}
}

func tradableSupplyInvariant(store types.KVStore) (string, bool) {
	var (
		msg    string
		broken bool
	)
	calTradableSupplies := make(map[string]math.Dec)

	iterateBalances(store, TradableBalancePrefix, func(_, denom, b string) bool {
		balance, err := math.NewNonNegativeDecFromString(b)
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

	if err := iterateSupplies(store, TradableSupplyPrefix, func(denom string, s string) (bool, error) {
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
		return false, nil
	}); err != nil {
		msg = fmt.Sprintf("error querying credit batch tradable supply %v", err)
	}

	return msg, broken
}

func (s serverImpl) retiredSupplyInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		store := ctx.KVStore(s.storeKey)
		return retiredSupplyInvariant(store)
	}
}

func retiredSupplyInvariant(store types.KVStore) (string, bool) {
	var (
		msg    string
		broken bool
	)
	calRetiredSupplies := make(map[string]math.Dec)
	iterateBalances(store, RetiredBalancePrefix, func(_, denom, b string) bool {
		balance, err := math.NewNonNegativeDecFromString(b)
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

	if err := iterateSupplies(store, RetiredSupplyPrefix, func(denom, s string) (bool, error) {
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

		return false, nil
	}); err != nil {
		msg = fmt.Sprintf("error querying credit batch supply %v", err)
	}

	return msg, broken
}
