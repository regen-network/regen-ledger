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

	creditDecs, err := utils.GetNonNegativeFixedDecs(precision, credit.TradableAmount, credit.RetiredAmount)
	if err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	sendAmtTradable, sendAmtRetired := creditDecs[0], creditDecs[1]
	if !sendAmtTradable.IsZero() {
		fromBalance, err := k.stateStore.BatchBalanceTable().Get(ctx, from, batch.Key)
		if err != nil {
			if err == ormerrors.NotFound {
				return ecocredit.ErrInsufficientCredits.Wrapf("you do not have any credits from batch %s", batch.Denom)
			}
			return err
		}

		fromDec, err := math.NewNonNegativeFixedDecFromString(fromBalance.TradableAmount, precision)
		if err != nil {
			return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
		}

		fromTradableBalance, err := math.SafeSubBalance(fromDec, sendAmtTradable)
		if err != nil {
			return ecocredit.ErrInsufficientCredits.Wrapf(
				"tradable balance: %s, send tradable amount %s", fromDec, sendAmtTradable,
			)
		}

		// update the "from" balance
		if err := k.stateStore.BatchBalanceTable().Update(ctx, &api.BatchBalance{
			BatchKey:       batch.Key,
			Address:        from,
			TradableAmount: fromTradableBalance.String(),
			RetiredAmount:  fromBalance.RetiredAmount,
			EscrowedAmount: fromBalance.EscrowedAmount,
		}); err != nil {
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

		toTradableBalance, err := math.NewNonNegativeFixedDecFromString(toBalance.TradableAmount, precision)
		if err != nil {
			return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
		}

		toTradableBalance, err = toTradableBalance.Add(sendAmtTradable)
		if err != nil {
			return err
		}

		// update the "to" balance
		if err := k.stateStore.BatchBalanceTable().Save(ctx, &api.BatchBalance{
			BatchKey:       batch.Key,
			Address:        to,
			TradableAmount: toTradableBalance.String(),
			RetiredAmount:  toBalance.RetiredAmount,
			EscrowedAmount: toBalance.EscrowedAmount,
		}); err != nil {
			return err
		}
	}

	if !sendAmtRetired.IsZero() {
		if err := retireCredit(sdkCtx, k.stateStore, batch, credit, precision, from, to, sendAmtRetired); err != nil {
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

func retireCredit(ctx sdk.Context, stateStore api.StateStore, batch *api.Batch, credit *types.MsgSend_SendCredits,
	precision uint32, from, to sdk.AccAddress, sendAmtRetired math.Dec) error {
	batchKey := batch.Key
	batchDenom := batch.Denom

	fromBalance, err := stateStore.BatchBalanceTable().Get(ctx, from, batchKey)
	if err != nil {
		if err == ormerrors.NotFound {
			return ecocredit.ErrInsufficientCredits.Wrapf("you do not have any credits from batch %s", batchDenom)
		}
		return err
	}

	tradableDec, err := math.NewNonNegativeFixedDecFromString(fromBalance.TradableAmount, precision)
	if err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	fromTradableBalance, err := math.SafeSubBalance(tradableDec, sendAmtRetired)
	if err != nil {
		return ecocredit.ErrInsufficientCredits.Wrapf(
			"tradable balance: %s, send retired amount %s", tradableDec, sendAmtRetired,
		)
	}

	// update the "from" balance
	if err := stateStore.BatchBalanceTable().Update(ctx, &api.BatchBalance{
		BatchKey:       batchKey,
		Address:        from,
		TradableAmount: fromTradableBalance.String(),
		RetiredAmount:  fromBalance.RetiredAmount,
		EscrowedAmount: fromBalance.EscrowedAmount,
	}); err != nil {
		return err
	}

	toBalance, err := stateStore.BatchBalanceTable().Get(ctx, to, batchKey)
	if err != nil {
		if err == ormerrors.NotFound {
			toBalance = &api.BatchBalance{
				BatchKey:       batchKey,
				Address:        to,
				TradableAmount: "0",
				RetiredAmount:  "0",
				EscrowedAmount: "0",
			}
		} else {
			return err
		}
	}

	toRetiredBalance, err := math.NewNonNegativeFixedDecFromString(toBalance.RetiredAmount, precision)
	if err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	toRetiredBalance, err = toRetiredBalance.Add(sendAmtRetired)
	if err != nil {
		return err
	}

	// update the "to" balance
	if err := stateStore.BatchBalanceTable().Save(ctx, &api.BatchBalance{
		BatchKey:       batchKey,
		Address:        to,
		TradableAmount: toBalance.TradableAmount,
		RetiredAmount:  toRetiredBalance.String(),
		EscrowedAmount: toBalance.EscrowedAmount,
	}); err != nil {
		return err
	}

	batchSupply, err := stateStore.BatchSupplyTable().Get(ctx, batchKey)
	if err != nil {
		return err
	}

	supplyDecs, err := utils.GetNonNegativeFixedDecs(precision, batchSupply.TradableAmount, batchSupply.RetiredAmount)
	if err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	batchSupplyTradable, batchSupplyRetired := supplyDecs[0], supplyDecs[1]

	batchSupplyRetired, err = batchSupplyRetired.Add(sendAmtRetired)
	if err != nil {
		return err
	}

	batchSupplyTradable, err = batchSupplyTradable.Sub(sendAmtRetired)
	if err != nil {
		return err
	}

	if err := stateStore.BatchSupplyTable().Update(ctx, &api.BatchSupply{
		BatchKey:        batchKey,
		TradableAmount:  batchSupplyTradable.String(),
		RetiredAmount:   batchSupplyRetired.String(),
		CancelledAmount: batchSupply.CancelledAmount,
	}); err != nil {
		return err
	}

	if err = ctx.EventManager().EmitTypedEvent(&types.EventRetire{
		Owner:        to.String(),
		BatchDenom:   credit.BatchDenom,
		Amount:       sendAmtRetired.String(),
		Jurisdiction: credit.RetirementJurisdiction,
		Reason:       credit.RetirementReason,
	}); err != nil {
		return err
	}

	return nil
}
