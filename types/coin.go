package types

import (
	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CoinToProtoCoin constructs a new protobuf coin from gogoproto coin.
// This function does not validate the coin.
func CoinToProtoCoin(coin sdk.Coin) *basev1beta1.Coin {
	if coin.Amount.IsNil() {
		return &basev1beta1.Coin{
			Denom:  coin.Denom,
			Amount: "",
		}
	}
	return &basev1beta1.Coin{
		Denom:  coin.Denom,
		Amount: coin.Amount.String(),
	}
}

// CoinsToProtoCoins constructs a new protobuf coin set from gogoproto coin set.
// This function does not validate the coin.
func CoinsToProtoCoins(coins sdk.Coins) []*basev1beta1.Coin {
	result := make([]*basev1beta1.Coin, 0, coins.Len())
	for _, coin := range coins {
		result = append(result, CoinToProtoCoin(coin))
	}

	return result
}

// ProtoCoinToCoin constructs a new gogoproto coin from protobuf coin.
// This function will panic if the amount is negative or if the denomination is invalid.
func ProtoCoinToCoin(coin *basev1beta1.Coin) (sdk.Coin, bool) {
	amount, ok := sdk.NewIntFromString(coin.Amount)
	if !ok {
		return sdk.Coin{}, ok
	}

	return sdk.NewCoin(coin.Denom, amount), true
}

// ProtoCoinsToCoins constructs a new gogoproto coin set from protobuf coin set.
// This function will panic if the amount is negative or if the denomination is invalid.
func ProtoCoinsToCoins(coins []*basev1beta1.Coin) (sdk.Coins, bool) {
	result := make([]sdk.Coin, 0, len(coins))
	for _, coin := range coins {
		c, ok := ProtoCoinToCoin(coin)
		if !ok {
			return nil, ok
		}
		result = append(result, c)
	}

	return result, true
}
