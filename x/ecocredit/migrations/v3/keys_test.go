package v3

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParseBalanceKey(t *testing.T) {
	acc := sdk.AccAddress("abcd")
	denom := BatchDenomT("xyz")

	tradableKey := TradableBalanceKey(acc, denom)
	resAddr, denom1 := ParseBalanceKey(tradableKey)
	require.Equal(t, acc, resAddr)
	require.Equal(t, denom, denom1)

	retiredKey := RetiredBalanceKey(acc, denom)
	resAddr, denom1 = ParseBalanceKey(retiredKey)
	require.Equal(t, acc, resAddr)
	require.Equal(t, denom, denom1)

	tradableSupplyKey := TradableSupplyKey(denom)
	denom1 = ParseSupplyKey(tradableSupplyKey)
	require.Equal(t, denom, denom1)

	retiredSupplyKey := RetiredSupplyKey(denom)
	denom1 = ParseSupplyKey(retiredSupplyKey)
	require.Equal(t, denom, denom1)
}
