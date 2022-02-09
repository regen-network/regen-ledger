package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// In this keeper.go files we expose methods the basket.Keeper needs.

func (s serverImpl) GetCreateBasketFee(ctx context.Context) sdk.Coins {
	//TODO implement me
	panic("implement me")
}
