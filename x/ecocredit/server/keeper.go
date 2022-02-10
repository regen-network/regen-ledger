package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types/math"

	"github.com/regen-network/regen-ledger/types"
)

// In this keeper.go files we expose methods the basket.Keeper needs.

func (s serverImpl) GetCreateBasketFee(ctx context.Context) sdk.Coins {
	//TODO implement me
	panic("implement me")
}

func (s serverImpl) AddCreditBalance(ctx context.Context, owner sdk.AccAddress, batchDenom string, amount math.Dec, doRetire bool, retirementLocation string) error {
	sdkCtx := types.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(s.storeKey)
	if !doRetire {
		err := addAndSetDecimal(store, TradableBalanceKey(owner, batchDenomT(batchDenom)), amount)
		if err != nil {
			return err
		}
	} else {
		err := retire(sdkCtx, store, owner, batchDenomT(batchDenom), amount, retirementLocation)
		if err != nil {
			return err
		}

		err = addAndSetDecimal(store, RetiredSupplyKey(batchDenomT(batchDenom)), amount)
		if err != nil {
			return err
		}
	}

	return nil
}
