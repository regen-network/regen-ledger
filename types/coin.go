package types

import (
	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
