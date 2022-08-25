package types

import (
	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CoinsToProtoCoins constructs a new protobuf coin set from gogoproto coin set.
// It returns an error, if coins are not sorted, have negitive amount, invalid or duplicate denomination.
func CoinsToProtoCoins(coins sdk.Coins) []*basev1beta1.Coin {

	result := make([]*basev1beta1.Coin, 0, coins.Len())
	for _, coin := range coins {
		result = append(result, &basev1beta1.Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.String(),
		})
	}

	return result
}

// ProtoCoinsToCoins constructs a new gogoproto coin set from protobuf coin set.
// It will panic if the amount is negative or if the denomination is invalid.
func ProtoCoinsToCoins(coins []*basev1beta1.Coin) (sdk.Coins, bool) {
	result := make([]sdk.Coin, 0, len(coins))
	for _, coin := range coins {
		amount, ok := sdk.NewIntFromString(coin.Amount)
		if !ok {
			return nil, ok
		}
		result = append(result, sdk.NewCoin(coin.Denom, amount))
	}

	return result, true
}
