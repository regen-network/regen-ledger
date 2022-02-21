package server

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

// RegisterInvariants registers the ecocredit module invariants.
func (s serverImpl) RegisterInvariants(ir sdk.InvariantRegistry) {
	ir.RegisterRoute(ecocredit.ModuleName, "tradable-supply", s.tradableSupplyInvariant())
	ir.RegisterRoute(ecocredit.ModuleName, "retired-supply", s.retiredSupplyInvariant())
	s.basketKeeper.RegisterInvariants(ir)
}

func (s serverImpl) tradableSupplyInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		store := ctx.KVStore(s.storeKey)
		goCtx := sdk.WrapSDKContext(ctx)
		basketBalances := s.getBasketBalanceMap(goCtx)
		return tradableSupplyInvariant(store, basketBalances)
	}
}

func (s serverImpl) getBasketBalanceMap(ctx context.Context) map[string]math.Dec {
	res, err := s.basketKeeper.Baskets(ctx, &baskettypes.QueryBasketsRequest{})
	if err != nil {
		panic(err)
	}
	batchBalances := make(map[string]math.Dec) // map of a basket batch_denom to balance
	for _, basket := range res.Baskets {
		res, err := s.basketKeeper.BasketBalances(ctx, &baskettypes.QueryBasketBalancesRequest{BasketDenom: basket.BasketDenom})
		if err != nil {
			panic(err)
		}
		for _, bal := range res.Balances {
			amount, err := math.NewDecFromString(bal.Balance)
			if err != nil {
				panic(err)
			}
			if existingBal, ok := batchBalances[bal.BatchDenom]; ok {
				existingBal, err = existingBal.Add(amount)
				if err != nil {
					panic(err)
				}
				batchBalances[bal.BatchDenom] = existingBal
			} else {
				batchBalances[bal.BatchDenom] = amount
			}
		}
	}
	return batchBalances
}

func tradableSupplyInvariant(store types.KVStore, basketBalances map[string]math.Dec) (string, bool) {
	var (
		msg    string
		broken bool
	)
	// sum of tradeable eco credits with credits locked in baskets
	sumBatchSupplies := make(map[string]math.Dec) // map batch denom => balance

	ecocredit.IterateBalances(store, ecocredit.TradableBalancePrefix, func(_, denom, b string) bool {
		balance, err := math.NewNonNegativeDecFromString(b)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("error while parsing tradable balance %v", err)
		}
		if supply, ok := sumBatchSupplies[denom]; ok {
			supply, err := math.SafeAddBalance(supply, balance)
			if err != nil {
				broken = true
				msg += fmt.Sprintf("error adding credit batch tradable supply %v", err)
			}
			sumBatchSupplies[denom] = supply
		} else {
			sumBatchSupplies[denom] = balance
		}

		return false
	})

	for denom, amt := range basketBalances {
		if amount, ok := sumBatchSupplies[denom]; ok {
			amount, err := math.SafeAddBalance(amount, amt)
			if err != nil {
				panic(err)
			}
			sumBatchSupplies[denom] = amount
		} else {
			panic("unknown denom in basket")
		}
	}

	if err := ecocredit.IterateSupplies(store, ecocredit.TradableSupplyPrefix, func(denom string, s string) (bool, error) {
		supply, err := math.NewNonNegativeDecFromString(s)
		if err != nil {
			broken = true
			msg += fmt.Sprintf("error while parsing tradable supply for denom: %s", denom)
		}
		if s1, ok := sumBatchSupplies[denom]; ok {
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
	ecocredit.IterateBalances(store, ecocredit.RetiredBalancePrefix, func(_, denom, b string) bool {
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

	if err := ecocredit.IterateSupplies(store, ecocredit.RetiredSupplyPrefix, func(denom, s string) (bool, error) {
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
