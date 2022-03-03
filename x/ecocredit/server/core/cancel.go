package core

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
)

func (k Keeper) Cancel(ctx context.Context, req *v1beta1.MsgCancel) (*v1beta1.MsgCancelResponse, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	holder, _ := sdk.AccAddressFromBech32(req.Holder)

	for _, credit := range req.Credits {
		batch, err := k.stateStore.BatchInfoStore().GetByBatchDenom(ctx, credit.BatchDenom)
		if err != nil {
			return nil, err
		}
		creditType, err := k.getCreditTypeFromBatchDenom(ctx, batch.BatchDenom)
		if err != nil {
			return nil, err
		}
		precision := creditType.Precision

		userBalance, err := k.stateStore.BatchBalanceStore().Get(ctx, holder, batch.Id)
		if err != nil {
			return nil, err
		}
		batchSupply, err := k.stateStore.BatchSupplyStore().Get(ctx, batch.Id)
		if err != nil {
			return nil, err
		}
		decs, err := getNonNegativeFixedDecs(precision, credit.Amount, batchSupply.TradableAmount, userBalance.Tradable, batchSupply.CancelledAmount)
		if err != nil {
			return nil, err
		}
		amtToCancelDec, supplyTradable, userBalTradable, cancelledDec := decs[0], decs[1], decs[2], decs[3]
		userBalTradable, err = math.SafeSubBalance(userBalTradable, amtToCancelDec)
		if err != nil {
			return nil, err
		}
		supplyTradable, err = math.SafeSubBalance(supplyTradable, amtToCancelDec)
		if err != nil {
			return nil, err
		}
		cancelledDec, err = cancelledDec.Add(amtToCancelDec)
		if err != nil {
			return nil, err
		}
		if err = k.stateStore.BatchBalanceStore().Update(ctx, &ecocreditv1beta1.BatchBalance{
			Address:  holder,
			BatchId:  batch.Id,
			Tradable: userBalTradable.String(),
			Retired:  userBalance.Retired,
		}); err != nil {
			return nil, err
		}
		if err = k.stateStore.BatchSupplyStore().Update(ctx, &ecocreditv1beta1.BatchSupply{
			BatchId:         batch.Id,
			TradableAmount:  supplyTradable.String(),
			RetiredAmount:   batchSupply.RetiredAmount,
			CancelledAmount: cancelledDec.String(),
		}); err != nil {
			return nil, err
		}
		if err = sdkCtx.EventManager().EmitTypedEvent(&v1beta1.EventCancel{
			Canceller:  holder.String(),
			BatchDenom: credit.BatchDenom,
			Amount:     credit.Amount,
		}); err != nil {
			return nil, err
		}
		sdkCtx.GasMeter().ConsumeGas(gasCostPerIteration, "cancel ecocredits")
	}
	return &v1beta1.MsgCancelResponse{}, nil
}
