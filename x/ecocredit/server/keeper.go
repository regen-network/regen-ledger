package server

import (
	"context"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// In this keeper.go files we expose methods the basket.Keeper needs.

func (s serverImpl) GetCreateBasketFee(ctx context.Context) sdk.Coins {
	sdkCtx := types.UnwrapSDKContext(ctx).Context
	var params ecocredit.Params
	s.paramSpace.GetParamSet(sdkCtx, &params)
	return params.BasketCreationFee
}
