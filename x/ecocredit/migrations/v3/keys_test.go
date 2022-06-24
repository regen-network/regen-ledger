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

func TestParseBatchDenom(t *testing.T) {
	testCases := []struct {
		name       string
		denom      string
		expErr     bool
		errMessage string
	}{
		{
			"invalid denom",
			"invalid-denom",
			true,
			"invalid batch denom",
		},
		{
			"invalid date",
			"C01-1234-1234-001",
			true,
			"parsing time",
		},
		{
			"valid test",
			"C01-20220509-20230509-001",
			false,
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sd, ed, err := ParseBatchDenom(tc.denom)
			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errMessage)
			} else {
				require.NoError(t, err)
				require.NotNil(t, sd)
				require.NotNil(t, ed)
			}
		})
	}
}
