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
func (k Keeper) Send(ctx context.Context, req *types.MsgSend) (*types.MsgSendResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sender, _ := sdk.AccAddressFromBech32(req.Sender)
	recipient, _ := sdk.AccAddressFromBech32(req.Recipient)

	for _, credit := range req.Credits {

		// get credit batch
		batch, err := k.stateStore.BatchTable().GetByDenom(ctx, credit.BatchDenom)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf(
				"could not get batch with denom %s: %s", credit.BatchDenom, err,
			)
		}

		// get credit type precision
		creditType, err := utils.GetCreditTypeFromBatchDenom(ctx, k.stateStore, batch.Denom)
		if err != nil {
			return nil, err
		}
		precision := creditType.Precision

		// get decimal values of credits to send
		creditDecs, err := utils.GetNonNegativeFixedDecs(precision, credit.TradableAmount, credit.RetiredAmount)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrap(err.Error())
		}
		sendAmtTradable, sendAmtRetired := creditDecs[0], creditDecs[1]

		// send tradable credits
		if !sendAmtTradable.IsZero() {
			err := k.sendTradable(sdkCtx, sendParams{
				precision:  precision,
				batchKey:   batch.Key,
				batchDenom: batch.Denom,
				sender:     sender,
				recipient:  recipient,
				amount:     sendAmtTradable,
			})
			if err != nil {
				return nil, err
			}
		}

		// send retired credits
		if !sendAmtRetired.IsZero() {
			err := k.sendRetired(sdkCtx, sendParams{
				precision:  precision,
				batchKey:   batch.Key,
				batchDenom: batch.Denom,
				sender:     sender,
				recipient:  recipient,
				amount:     sendAmtRetired,
			})
			if err != nil {
				return nil, err
			}

			// emit retire event
			if err = sdkCtx.EventManager().EmitTypedEvent(&types.EventRetire{
				Owner:        req.Recipient,
				BatchDenom:   credit.BatchDenom,
				Amount:       credit.RetiredAmount,
				Jurisdiction: credit.RetirementJurisdiction,
				Reason:       credit.RetirementReason,
			}); err != nil {
				return nil, err
			}
		}

		// check and set zero-value
		if credit.TradableAmount == "" {
			credit.TradableAmount = "0"
		}

		// check and set zero-value
		if credit.RetiredAmount == "" {
			credit.RetiredAmount = "0"
		}

		// emit transfer event
		if err = sdkCtx.EventManager().EmitTypedEvent(&types.EventTransfer{
			Sender:         req.Sender,
			Recipient:      req.Recipient,
			BatchDenom:     credit.BatchDenom,
			TradableAmount: credit.TradableAmount,
			RetiredAmount:  credit.RetiredAmount,
		}); err != nil {
			return nil, err
		}

		// increase gas consumption for each iteration
		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/MsgSend credit iteration")
	}

	return &types.MsgSendResponse{}, nil
}

type sendParams struct {
	precision  uint32
	batchKey   uint64
	batchDenom string
	sender     sdk.AccAddress
	recipient  sdk.AccAddress
	amount     math.Dec
}

func (k Keeper) sendTradable(ctx context.Context, params sendParams) error {

	// get sender balance and return error if balance does not exist
	senderBalance, err := k.stateStore.BatchBalanceTable().Get(ctx, params.sender, params.batchKey)
	if err != nil {
		if err == ormerrors.NotFound {
			return ecocredit.ErrInsufficientCredits.Wrapf(
				"you do not have any credits from batch %s", params.batchDenom,
			)
		}
		return err
	}

	// get decimal value of sender tradable balance
	senderTradable, err := math.NewNonNegativeFixedDecFromString(senderBalance.TradableAmount, params.precision)
	if err != nil {
		return err
	}

	// subtract send amount from sender tradable balance
	newSenderTradable, err := math.SafeSubBalance(senderTradable, params.amount)
	if err != nil {
		return ecocredit.ErrInsufficientCredits.Wrapf(
			"tradable balance: %s, send tradable amount %s", senderTradable, params.amount,
		)
	}

	// update sender balance with new tradable amount
	if err := k.stateStore.BatchBalanceTable().Update(ctx, &api.BatchBalance{
		BatchKey:       params.batchKey,
		Address:        params.sender,
		TradableAmount: newSenderTradable.String(),
		RetiredAmount:  senderBalance.RetiredAmount,
		EscrowedAmount: senderBalance.EscrowedAmount,
	}); err != nil {
		return err
	}

	// get recipient balance and create empty balance if balance does not exist
	recipientBalance, err := k.stateStore.BatchBalanceTable().Get(ctx, params.recipient, params.batchKey)
	if err != nil {
		if err == ormerrors.NotFound {
			recipientBalance = &api.BatchBalance{
				BatchKey:       params.batchKey,
				Address:        params.recipient,
				TradableAmount: "0",
				RetiredAmount:  "0",
				EscrowedAmount: "0",
			}
		} else {
			return err
		}
	}

	// get decimal value of recipient tradable balance
	recipientTradable, err := math.NewNonNegativeFixedDecFromString(recipientBalance.TradableAmount, params.precision)
	if err != nil {
		return err
	}

	// add send amount to recipient tradable balance
	newRecipientTradable, err := recipientTradable.Add(params.amount)
	if err != nil {
		return err
	}

	// update recipient balance with new tradable amount
	return k.stateStore.BatchBalanceTable().Save(ctx, &api.BatchBalance{
		BatchKey:       params.batchKey,
		Address:        params.recipient,
		TradableAmount: newRecipientTradable.String(),
		RetiredAmount:  recipientBalance.RetiredAmount,
		EscrowedAmount: recipientBalance.EscrowedAmount,
	})
}

func (k Keeper) sendRetired(ctx sdk.Context, params sendParams) error {

	// get sender balance and return error if balance does not exist
	senderBalance, err := k.stateStore.BatchBalanceTable().Get(ctx, params.sender, params.batchKey)
	if err != nil {
		if err == ormerrors.NotFound {
			return ecocredit.ErrInsufficientCredits.Wrapf(
				"you do not have any credits from batch %s", params.batchDenom,
			)
		}
		return err
	}

	// get decimal value of sender tradable balance
	senderTradable, err := math.NewNonNegativeFixedDecFromString(senderBalance.TradableAmount, params.precision)
	if err != nil {
		return err
	}

	// subtract send amount from sender tradable balance
	newSenderTradable, err := math.SafeSubBalance(senderTradable, params.amount)
	if err != nil {
		return ecocredit.ErrInsufficientCredits.Wrapf(
			"tradable balance: %s, send retired amount %s", senderTradable, params.amount,
		)
	}

	// update sender balance with new tradable amount
	if err := k.stateStore.BatchBalanceTable().Update(ctx, &api.BatchBalance{
		BatchKey:       params.batchKey,
		Address:        params.sender,
		TradableAmount: newSenderTradable.String(),
		RetiredAmount:  senderBalance.RetiredAmount,
		EscrowedAmount: senderBalance.EscrowedAmount,
	}); err != nil {
		return err
	}

	// get recipient balance and create empty balance if balance does not exist
	recipientBalance, err := k.stateStore.BatchBalanceTable().Get(ctx, params.recipient, params.batchKey)
	if err != nil {
		if err == ormerrors.NotFound {
			recipientBalance = &api.BatchBalance{
				BatchKey:       params.batchKey,
				Address:        params.recipient,
				TradableAmount: "0",
				RetiredAmount:  "0",
				EscrowedAmount: "0",
			}
		} else {
			return err
		}
	}

	// get decimal value of recipient retired balance
	recipientRetired, err := math.NewNonNegativeFixedDecFromString(recipientBalance.RetiredAmount, params.precision)
	if err != nil {
		return err
	}

	// add send amount to recipient retired balance
	newRecipientRetired, err := recipientRetired.Add(params.amount)
	if err != nil {
		return err
	}

	// update recipient balance with new retired amount
	if err := k.stateStore.BatchBalanceTable().Save(ctx, &api.BatchBalance{
		BatchKey:       params.batchKey,
		Address:        params.recipient,
		TradableAmount: recipientBalance.TradableAmount,
		RetiredAmount:  newRecipientRetired.String(),
		EscrowedAmount: recipientBalance.EscrowedAmount,
	}); err != nil {
		return err
	}

	// get batch supply
	batchSupply, err := k.stateStore.BatchSupplyTable().Get(ctx, params.batchKey)
	if err != nil {
		return err
	}

	// get decimal values of batch supply
	supplyDecs, err := utils.GetNonNegativeFixedDecs(params.precision, batchSupply.TradableAmount, batchSupply.RetiredAmount)
	if err != nil {
		return err
	}
	supplyTradable, supplyRetired := supplyDecs[0], supplyDecs[1]

	// subtract send amount from tradable supply
	newSupplyTradable, err := supplyTradable.Sub(params.amount)
	if err != nil {
		return err
	}

	// add send amount to retired supply
	newSupplyRetired, err := supplyRetired.Add(params.amount)
	if err != nil {
		return err
	}

	// update batch supply with new tradable and retired amounts
	return k.stateStore.BatchSupplyTable().Update(ctx, &api.BatchSupply{
		BatchKey:        params.batchKey,
		TradableAmount:  newSupplyTradable.String(),
		RetiredAmount:   newSupplyRetired.String(),
		CancelledAmount: batchSupply.CancelledAmount,
	})
}
