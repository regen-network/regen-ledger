package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

// Cancel credits, removing them from the supply and balance of the holder
func (k Keeper) Cancel(ctx context.Context, req *core.MsgCancel) (*core.MsgCancelResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	holder, err := sdk.AccAddressFromBech32(req.Holder)
	if err != nil {
		return nil, err
	}

	creditTypeMap := utils.GetCreditTypeMap(sdkCtx, k.paramsKeeper)
	for _, credit := range req.Credits {
		batch, err := k.stateStore.BatchInfoTable().GetByBatchDenom(ctx, credit.BatchDenom)
		if err != nil {
			return nil, err
		}
		class, err := k.getClassFromBatchDenom(ctx, batch.BatchDenom)
		if err != nil {
			return nil, err
		}
		creditType, ok := creditTypeMap[class.CreditType]
		if !ok {
			return nil, sdkerrors.ErrNotFound.Wrapf("could not find credit type %s", class.CreditType)
		}
		userBalance, err := k.stateStore.BatchBalanceTable().Get(ctx, holder, batch.Id)
		if err != nil {
			return nil, err
		}
		batchSupply, err := k.stateStore.BatchSupplyTable().Get(ctx, batch.Id)
		if err != nil {
			return nil, err
		}
		decs, err := utils.GetNonNegativeFixedDecs(creditType.Precision, credit.Amount, batchSupply.TradableAmount, userBalance.Tradable, batchSupply.CancelledAmount)
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
			Address:  holder,
			BatchId:  batch.Id,
			Tradable: userBalTradable.String(),
			Retired:  userBalance.Retired,
		}); err != nil {
			return nil, err
		}

		if err = k.stateStore.BatchSupplyTable().Update(ctx, &api.BatchSupply{
			BatchId:         batch.Id,
			TradableAmount:  supplyTradable.String(),
			RetiredAmount:   batchSupply.RetiredAmount,
			CancelledAmount: cancelledDec.String(),
		}); err != nil {
			return nil, err
		}

		if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventCancel{
			Canceller:  holder.String(),
			BatchDenom: credit.BatchDenom,
			Amount:     credit.Amount,
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/core/MsgCancel credit iteration")
	}
	return &core.MsgCancelResponse{}, nil
}
