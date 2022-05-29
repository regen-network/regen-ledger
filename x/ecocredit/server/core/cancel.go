package core

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit"

	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Cancel credits, removing them from the supply and balance of the holder
func (k Keeper) Cancel(ctx context.Context, req *core.MsgCancel) (*core.MsgCancelResponse, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	holder, err := sdk.AccAddressFromBech32(req.Holder)
	if err != nil {
		return nil, err
	}

	for _, credit := range req.Credits {
		batch, err := k.stateStore.BatchTable().GetByDenom(ctx, credit.BatchDenom)
		if err != nil {
			return nil, err
		}
		creditType, err := utils.GetCreditTypeFromBatchDenom(ctx, k.stateStore, batch.Denom)
		if err != nil {
			return nil, err
		}
		userBalance, err := k.stateStore.BatchBalanceTable().Get(ctx, holder, batch.Key)
		if err != nil {
			return nil, err
		}
		batchSupply, err := k.stateStore.BatchSupplyTable().Get(ctx, batch.Key)
		if err != nil {
			return nil, err
		}
		decs, err := utils.GetNonNegativeFixedDecs(creditType.Precision, credit.Amount, batchSupply.TradableAmount, userBalance.TradableAmount, batchSupply.CancelledAmount)
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

		if err = k.stateStore.BatchBalanceTable().Update(ctx, &api.BatchBalance{
			BatchKey:       batch.Key,
			Address:        holder,
			TradableAmount: userBalTradable.String(),
			RetiredAmount:  userBalance.RetiredAmount,
		}); err != nil {
			return nil, err
		}

		if err = k.stateStore.BatchSupplyTable().Update(ctx, &api.BatchSupply{
			BatchKey:        batch.Key,
			TradableAmount:  supplyTradable.String(),
			RetiredAmount:   batchSupply.RetiredAmount,
			CancelledAmount: cancelledDec.String(),
		}); err != nil {
			return nil, err
		}

		if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventCancel{
			Owner:      holder.String(),
			BatchDenom: credit.BatchDenom,
			Amount:     credit.Amount,
			Reason:     req.Reason,
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/core/MsgCancel credit iteration")
	}
	return &core.MsgCancelResponse{}, nil
}
