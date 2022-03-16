package basket

import (
	"context"
	"fmt"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (k Keeper) RegisterInvariants(ir sdk.InvariantRegistry) {
	ir.RegisterRoute(ecocredit.ModuleName, "basket-supply", k.basketSupplyInvariant())
}

func (k Keeper) basketSupplyInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		goCtx := sdk.WrapSDKContext(ctx)

		bals, err := k.computeBasketBalances(goCtx)
		if err != nil {
			return err.Error(), true
		}
		return BasketSupplyInvariant(ctx, k.stateStore.BasketTable(), k.bankKeeper, bals)
	}
}

type bankSupplyStore interface {
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
}

// BasketSupplyInvariant cross check the balance of baskets and bank
func BasketSupplyInvariant(ctx sdk.Context, store api.BasketTable, bank bankSupplyStore, basketBalances map[uint64]math.Dec) (string, bool) {
	goCtx := sdk.WrapSDKContext(ctx)

	bids := make([]uint64, len(basketBalances))
	i := 0
	for bid := range basketBalances {
		bids[i] = bid
		i++
	}
	sort.Slice(bids, func(i, j int) bool { return bids[i] < bids[j] })

	var inbalances []string
	for _, bid := range bids {
		b, err := store.Get(goCtx, bid)
		if err != nil {
			return fmt.Sprintf("Can't get basket %v: %v", bid, err), true
		}
		bal := basketBalances[bid]
		exp := math.NewDecFinite(1, int32(b.Exponent))
		mul, err := bal.Mul(exp)
		if err != nil {
			return fmt.Sprintf("Can't multiply balance by exponent, %v", err), true
		}
		balInt, err := mul.BigInt()
		if err != nil {
			return fmt.Sprintf("Can't convert Dec to big.Int, %v", err), true
		}
		c := bank.GetSupply(ctx, b.BasketDenom)
		balSdkInt := sdk.NewIntFromBigInt(balInt)
		if !c.Amount.Equal(balSdkInt) {
			inbalances = append(inbalances, fmt.Sprintf("Basket denom %s is imbalanced, expected: %v, got %v",
				b.BasketDenom, balSdkInt, c.Amount))
		}
	}
	if len(inbalances) != 0 {
		return strings.Join(inbalances, "\n"), true
	}
	return "", false
}

// computeBasketBalances returns a map from basket id to the total number of eco credits
func (k Keeper) computeBasketBalances(ctx context.Context) (map[uint64]math.Dec, error) {
	it, err := k.stateStore.BasketBalanceTable().List(ctx, &api.BasketBalancePrimaryKey{})
	if err != nil {
		return nil, fmt.Errorf("can't create basket balance iterator, %w", err)
	}
	balances := map[uint64]math.Dec{}
	for it.Next() {
		b, err := it.Value()
		if err != nil {
			return nil, fmt.Errorf("Can't get basket balance %w", err)
		}
		bal, err := math.NewDecFromString(b.Balance)
		if err != nil {
			return nil, fmt.Errorf("Can't decode balance %s as math.Dec: %w", b.Balance, err)
		}
		if a, ok := balances[b.BasketId]; ok {
			if a, err = a.Add(bal); err != nil {
				return nil, fmt.Errorf("Can't add balances: %w", err)
			}
			balances[b.BasketId] = a
		} else {
			balances[b.BasketId] = bal
		}
	}
	return balances, nil
}
