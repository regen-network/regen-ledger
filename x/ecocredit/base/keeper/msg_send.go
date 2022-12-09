package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/v2/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/server/utils"
)

// Send sends credits to a recipient.
// Send also retires credits if the amount to retire is specified in the request.
// NOTE: This method will return an error if both sender and recipient are same.
func (k Keeper) Send(ctx context.Context, req *types.MsgSend) (*types.MsgSendResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sender, _ := sdk.AccAddressFromBech32(req.Sender)
	recipient, _ := sdk.AccAddressFromBech32(req.Recipient)

	if sender.Equals(recipient) {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("sender and recipient cannot be the same")
	}

	for _, credit := range req.Credits {
		err := k.sendEcocredits(sdkCtx, credit, recipient, sender)
		if err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/MsgSend credit iteration")
	}
	return &types.MsgSendResponse{}, nil
}

func (k Keeper) sendEcocredits(sdkCtx sdk.Context, credit *types.MsgSend_SendCredits, to, from sdk.AccAddress) error {
	ctx := sdk.WrapSDKContext(sdkCtx)
	batch, err := k.stateStore.BatchTable().GetByDenom(ctx, credit.BatchDenom)
	if err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("could not get batch with denom %s: %s", credit.BatchDenom, err.Error())
	}
	creditType, err := utils.GetCreditTypeFromBatchDenom(ctx, k.stateStore, batch.Denom)
	if err != nil {
		return err
	}
	precision := creditType.Precision

	batchSupply, err := k.stateStore.BatchSupplyTable().Get(ctx, batch.Key)
	if err != nil {
		return err
	}
	fromBalance, err := k.stateStore.BatchBalanceTable().Get(ctx, from, batch.Key)
	if err != nil {
		if err == ormerrors.NotFound {
			return ecocredit.ErrInsufficientCredits.Wrapf("you do not have any credits from batch %s", batch.Denom)
		}
		return err
	}

	toBalance, err := k.stateStore.BatchBalanceTable().Get(ctx, to, batch.Key)
	if err != nil {
		if err == ormerrors.NotFound {
			toBalance = &api.BatchBalance{
				BatchKey:       batch.Key,
				Address:        to,
				TradableAmount: "0",
				RetiredAmount:  "0",
				EscrowedAmount: "0",
			}
		} else {
			return err
		}
	}
	decs, err := utils.GetNonNegativeFixedDecs(precision, toBalance.TradableAmount, toBalance.RetiredAmount, fromBalance.TradableAmount, fromBalance.RetiredAmount, credit.TradableAmount, credit.RetiredAmount, batchSupply.TradableAmount, batchSupply.RetiredAmount)
	if err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}
	toTradableBalance, toRetiredBalance,
		fromTradableBalance, fromRetiredBalance,
		sendAmtTradable, sendAmtRetired,
		batchSupplyTradable, batchSupplyRetired := decs[0], decs[1], decs[2], decs[3], decs[4], decs[5], decs[6], decs[7]

	if !sendAmtTradable.IsZero() {
		fromTradableBalance, err = math.SafeSubBalance(fromTradableBalance, sendAmtTradable)
		if err != nil {
			return ecocredit.ErrInsufficientCredits.Wrapf(
				"tradable balance: %s, send tradable amount %s", decs[2], sendAmtTradable,
			)
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
			return ecocredit.ErrInsufficientCredits.Wrapf(
				"tradable balance: %s, send retired amount %s", decs[2], sendAmtRetired,
			)
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
	if err := k.stateStore.BatchBalanceTable().Save(ctx, &api.BatchBalance{
		BatchKey:       batch.Key,
		Address:        to,
		TradableAmount: toTradableBalance.String(),
		RetiredAmount:  toRetiredBalance.String(),
		EscrowedAmount: toBalance.EscrowedAmount,
	}); err != nil {
		return err
	}

	// update the "from" balance
	if err := k.stateStore.BatchBalanceTable().Update(ctx, &api.BatchBalance{
		BatchKey:       batch.Key,
		Address:        from,
		TradableAmount: fromTradableBalance.String(),
		RetiredAmount:  fromRetiredBalance.String(),
		EscrowedAmount: fromBalance.EscrowedAmount,
	}); err != nil {
		return err
	}
	// update the "retired" supply only if credits were retired
	if didRetire {
		if err := k.stateStore.BatchSupplyTable().Update(ctx, &api.BatchSupply{
			BatchKey:        batch.Key,
			TradableAmount:  batchSupplyTradable.String(),
			RetiredAmount:   batchSupplyRetired.String(),
			CancelledAmount: batchSupply.CancelledAmount,
		}); err != nil {
			return err
		}
		if err = sdkCtx.EventManager().EmitTypedEvent(&types.EventRetire{
			Owner:        to.String(),
			BatchDenom:   credit.BatchDenom,
			Amount:       sendAmtRetired.String(),
			Jurisdiction: credit.RetirementJurisdiction,
			Reason:       credit.RetirementReason,
		}); err != nil {
			return err
		}
	}
	if err = sdkCtx.EventManager().EmitTypedEvent(&types.EventTransfer{
		Sender:         from.String(),
		Recipient:      to.String(),
		BatchDenom:     credit.BatchDenom,
		TradableAmount: sendAmtTradable.String(),
		RetiredAmount:  sendAmtRetired.String(),
	}); err != nil {
		return err
	}
	return nil
}
