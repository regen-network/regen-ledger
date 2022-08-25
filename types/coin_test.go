package types_test

import (
	"testing"

	"cosmossdk.io/math"
	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/regen-network/regen-ledger/types"
)

func TestConvertCoinsToProtoCoins(t *testing.T) {
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

func TestConvertProtoCoinsToCoins(t *testing.T) {
	testCases := []struct {
		name     string
		input    []*basev1beta1.Coin
		expected sdk.Coins
		isPanics bool
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
			"should panic: negitive amount",
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
			if tc.isPanics {
				assert.Panics(t, func() { types.ProtoCoinsToCoins(tc.input) })
			} else {
				result, ok := types.ProtoCoinsToCoins(tc.input)
				assert.True(t, ok)
				assert.ElementsMatch(t, result, tc.expected)
			}
		})
	}
}
