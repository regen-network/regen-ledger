package basket_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
)

type BasketWithSupply struct {
	supply int64
	b      api.Basket
}

type BankSupplyMock map[string]sdk.Coin

func (bs BankSupplyMock) GetSupply(_ sdk.Context, denom string) sdk.Coin {
	if c, ok := bs[denom]; ok {
		return c
	}
	return sdk.NewInt64Coin(denom, 0)
}

func TestBasketSupplyInvarint(t *testing.T) {
	require := require.New(t)
	s := setupBase(t)

	baskets := []BasketWithSupply{
		{10, api.Basket{BasketDenom: "bb1", Name: "b1"}},
		{20, api.Basket{BasketDenom: "bb2", Name: "b2"}},
	}
	store := s.stateStore.BasketStore()
	basketBalances := map[uint64]math.Dec{}
	correctBalances := BankSupplyMock{}
	for _, b := range baskets {
		id, err := store.InsertReturningID(s.ctx, &b.b)
		require.NoError(err)
		basketBalances[id] = math.NewDecFromInt64(b.supply)
		correctBalances[b.b.BasketDenom] = sdk.NewInt64Coin(b.b.BasketDenom, b.supply)
	}

	tcs := []struct {
		name string
		bank BankSupplyMock
		msg  string
	}{
		{"no bank supply",
			BankSupplyMock{}, "imbalanced"},
		{"partial bank supply",
			BankSupplyMock{"bb1": newCoin("bb1", 10)}, "bb2 is imbalanced"},
		{"smaller bank supply",
			BankSupplyMock{"bb1": newCoin("bb1", 8)}, "bb1 is imbalanced"},
		{"smaller bank supply2",
			BankSupplyMock{"bb1": newCoin("bb1", 10), "bb2": newCoin("bb2", 10)}, "bb2 is imbalanced"},
		{"bigger bank supply",
			BankSupplyMock{"bb1": newCoin("bb1", 10), "bb2": newCoin("bb2", 30)}, "bb2 is imbalanced"},

		{"all good",
			correctBalances, ""},
		{"more denoms",
			BankSupplyMock{"bb1": newCoin("bb1", 10), "bb2": newCoin("bb2", 20), "other": newCoin("other", 100)}, ""},
	}

	for _, tc := range tcs {
		tc.bank.GetSupply(s.sdkCtx, "abc")

		msg, _ := basket.BasketSupplyInvariant(s.sdkCtx, store, tc.bank, basketBalances)
		if tc.msg != "" {
			require.Contains(msg, tc.msg, tc.name)
		} else {
			require.Empty(msg, tc.name)
		}
	}

	t.Log(baskets)
}

func newCoin(denom string, a int64) sdk.Coin {
	return sdk.NewInt64Coin(denom, a)
}
