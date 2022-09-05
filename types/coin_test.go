package types_test

import (
	"testing"

	"cosmossdk.io/math"
	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/regen-network/regen-ledger/types"
)

func TestCoinToProtoCoin(t *testing.T) {
	testCases := []struct {
		name     string
		input    sdk.Coin
		expected *basev1beta1.Coin
	}{
		{
			"empty coin",
			sdk.Coin{},
			&basev1beta1.Coin{
				Denom:  "",
				Amount: "",
			},
		},
		{
			"only denom",
			sdk.Coin{
				Denom: "uregen",
			},
			&basev1beta1.Coin{
				Denom:  "uregen",
				Amount: "",
			},
		},
		{
			"only amount",
			sdk.Coin{
				Amount: sdk.NewInt(10000),
			},
			&basev1beta1.Coin{
				Denom:  "",
				Amount: "10000",
			},
		},
		{
			"negative amount",
			sdk.Coin{
				Amount: sdk.NewInt(-10000),
			},
			&basev1beta1.Coin{
				Denom:  "",
				Amount: "-10000",
			},
		},
		{
			"valid coin",
			sdk.NewCoin("uregen", sdk.NewInt(10000)),
			&basev1beta1.Coin{
				Denom:  "uregen",
				Amount: "10000",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := types.CoinToProtoCoin(tc.input)
			assert.Equal(t, result, tc.expected)
		})
	}
}

func TestCoinsToProtoCoins(t *testing.T) {
	testCases := []struct {
		name     string
		input    sdk.Coins
		expected []*basev1beta1.Coin
	}{
		{
			"empty coins",
			sdk.NewCoins(),
			[]*basev1beta1.Coin{},
		},
		{
			"single coin",
			sdk.NewCoins(sdk.NewCoin("uregen", sdk.NewInt(10000))),
			[]*basev1beta1.Coin{
				{
					Denom:  "uregen",
					Amount: "10000",
				},
			},
		},
		{
			"multiple coins",
			sdk.NewCoins(sdk.NewCoin("uregen", sdk.NewInt(10000)), sdk.NewCoin("uatom", math.NewInt(2e7))),
			[]*basev1beta1.Coin{
				{
					Denom:  "uatom",
					Amount: "20000000",
				},
				{
					Denom:  "uregen",
					Amount: "10000",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := types.CoinsToProtoCoins(tc.input)
			assert.ElementsMatch(t, result, tc.expected)
		})
	}
}

func TestProtoCoinToCoin(t *testing.T) {
	testCases := []struct {
		name     string
		input    *basev1beta1.Coin
		expected sdk.Coin
		ok       bool
		panic    bool
	}{
		{
			"empty coin",
			&basev1beta1.Coin{},
			sdk.Coin{},
			false,
			false,
		},
		{
			"single coin",
			&basev1beta1.Coin{
				Denom:  "uregen",
				Amount: "10000",
			},
			sdk.NewCoin("uregen", sdk.NewInt(10000)),
			true,
			false,
		},
		{
			"should panic: negative amount",
			&basev1beta1.Coin{
				Denom:  "uregen",
				Amount: "-20000000",
			},
			sdk.Coin{},
			true,
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panic {
				assert.Panics(t, func() { types.ProtoCoinToCoin(tc.input) })
			} else {
				result, ok := types.ProtoCoinToCoin(tc.input)
				assert.Equal(t, tc.ok, ok)
				assert.Equal(t, result, tc.expected)
			}
		})
	}
}

func TestProtoCoinsToCoins(t *testing.T) {
	testCases := []struct {
		name     string
		input    []*basev1beta1.Coin
		expected sdk.Coins
		panic    bool
	}{
		{
			"empty coins",
			[]*basev1beta1.Coin{},
			sdk.NewCoins(),
			false,
		},
		{
			"single coin",
			[]*basev1beta1.Coin{
				{
					Denom:  "uregen",
					Amount: "10000",
				},
			},
			sdk.NewCoins(sdk.NewCoin("uregen", sdk.NewInt(10000))),
			false,
		},
		{
			"multiple coins",
			[]*basev1beta1.Coin{
				{
					Denom:  "uatom",
					Amount: "20000000",
				},
				{
					Denom:  "uregen",
					Amount: "10000",
				},
			},
			sdk.NewCoins(sdk.NewCoin("uregen", sdk.NewInt(10000)), sdk.NewCoin("uatom", math.NewInt(2e7))),
			false,
		},
		{
			"should panic: negative amount",
			[]*basev1beta1.Coin{
				{
					Denom:  "uatom",
					Amount: "-20000000",
				},
				{
					Denom:  "uregen",
					Amount: "10000",
				},
			},
			sdk.NewCoins(),
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panic {
				assert.Panics(t, func() { types.ProtoCoinsToCoins(tc.input) })
			} else {
				result, ok := types.ProtoCoinsToCoins(tc.input)
				assert.True(t, ok)
				assert.ElementsMatch(t, result, tc.expected)
			}
		})
	}
}
