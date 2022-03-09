package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
)

// Send sends credits to a recipient.
// Send also retires credits if the amount to retire is specified in the request.
func (k Keeper) Send(ctx context.Context, req *v1.MsgSend) (*v1.MsgSendResponse, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	sender, _ := sdk.AccAddressFromBech32(req.Sender)
	recipient, _ := sdk.AccAddressFromBech32(req.Recipient)

	for _, credit := range req.Credits {
		err := k.sendEcocredits(ctx, credit, recipient, sender)
		if err != nil {
			return nil, err
		}
		if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventReceive{
			Sender:         req.Sender,
			Recipient:      req.Recipient,
			BatchDenom:     credit.BatchDenom,
			TradableAmount: credit.TradableAmount,
			RetiredAmount:  credit.RetiredAmount,
		}); err != nil {
			return nil, err
		}
		sdkCtx.GasMeter().ConsumeGas(gasCostPerIteration, "send ecocredits")
	}
	return &v1.MsgSendResponse{}, nil
}

func (k Keeper) sendEcocredits(ctx context.Context, credit *v1.MsgSend_SendCredits, to, from sdk.AccAddress) error {
	batch, err := k.stateStore.BatchInfoStore().GetByBatchDenom(ctx, credit.BatchDenom)
	if err != nil {
		return err
	}
	creditType, err := k.getCreditTypeFromBatchDenom(ctx, batch.BatchDenom)
	if err != nil {
		return err
	}
	precision := creditType.Precision

	batchSupply, err := k.stateStore.BatchSupplyStore().Get(ctx, batch.Id)
	if err != nil {
		return err
	}
	fromBalance, err := k.stateStore.BatchBalanceStore().Get(ctx, from, batch.Id)
	if err != nil {
		if err == ormerrors.NotFound {
			return ecocredit.ErrInsufficientFunds.Wrapf("you do not have any credits from batch %s", batch.BatchDenom)
		}
		return err
	}

	toBalance, err := k.stateStore.BatchBalanceStore().Get(ctx, to, batch.Id)
	if err != nil {
		if err == ormerrors.NotFound {
			toBalance = &ecocreditv1.BatchBalance{
				Address:  to,
				BatchId:  batch.Id,
				Tradable: "0",
				Retired:  "0",
			}
		} else {
			return err
		}
	}
	decs, err := getNonNegativeFixedDecs(precision, toBalance.Tradable, toBalance.Retired, fromBalance.Tradable, fromBalance.Retired, credit.TradableAmount, credit.RetiredAmount, batchSupply.TradableAmount, batchSupply.RetiredAmount)
	if err != nil {
		return err
	}
	toTradableBalance, toRetiredBalance,
		fromTradableBalance, fromRetiredBalance,
		sendAmtTradable, sendAmtRetired,
		batchSupplyTradable, batchSupplyRetired := decs[0], decs[1], decs[2], decs[3], decs[4], decs[5], decs[6], decs[7]

	if !sendAmtTradable.IsZero() {
		fromTradableBalance, err = math.SafeSubBalance(fromTradableBalance, sendAmtTradable)
		if err != nil {
			return err
		}
		toTradableBalance, err = toTradableBalance.Add(sendAmtTradable)
		if err != nil {
			return err
		}
	}

	didRetire := false
	if !sendAmtRetired.IsZero() {
		didRetire = true
		fromTradableBalance, err = math.SafeSubBalance(fromTradableBalance, sendAmtRetired)
		if err != nil {
			return err
		}
		toRetiredBalance, err = toRetiredBalance.Add(sendAmtRetired)
		if err != nil {
			return err
		}
		batchSupplyRetired, err = batchSupplyRetired.Add(sendAmtRetired)
		if err != nil {
			return err
		}
		batchSupplyTradable, err = batchSupplyTradable.Sub(sendAmtRetired)
		if err != nil {
			return err
		}
	}
	// update the "to" balance
	if err := k.stateStore.BatchBalanceStore().Save(ctx, &ecocreditv1.BatchBalance{
		Address:  to,
		BatchId:  batch.Id,
		Tradable: toTradableBalance.String(),
		Retired:  toRetiredBalance.String(),
	}); err != nil {
		return err
	}

	// update the "from" balance
	if err := k.stateStore.BatchBalanceStore().Update(ctx, &ecocreditv1.BatchBalance{
		Address:  from,
		BatchId:  batch.Id,
		Tradable: fromTradableBalance.String(),
		Retired:  fromRetiredBalance.String(),
	}); err != nil {
		return err
	}
	// update the "retired" supply only if credits were retired
	if didRetire {
		if err := k.stateStore.BatchSupplyStore().Update(ctx, &ecocreditv1.BatchSupply{
			BatchId:         batch.Id,
			TradableAmount:  batchSupplyTradable.String(),
			RetiredAmount:   batchSupplyRetired.String(),
			CancelledAmount: batchSupply.CancelledAmount,
		}); err != nil {
			return err
		}
		if err = sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&ecocredit.EventRetire{
			Retirer:    to.String(),
			BatchDenom: credit.BatchDenom,
			Amount:     sendAmtRetired.String(),
			Location:   credit.RetirementLocation,
		}); err != nil {
			return err
		}
	}
	return nil
}
