package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

// Cancel credits, removing them from the supply and balance of the owner
func (k Keeper) Cancel(ctx context.Context, req *types.MsgCancel) (*types.MsgCancelResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	owner, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	for _, credit := range req.Credits {
		batch, err := k.stateStore.BatchTable().GetByDenom(ctx, credit.BatchDenom)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get batch with denom %s: %s", credit.BatchDenom, err.Error())
		}
		creditType, err := utils.GetCreditTypeFromBatchDenom(ctx, k.stateStore, batch.Denom)
		if err != nil {
			return nil, err
		}
		userBalance, err := k.stateStore.BatchBalanceTable().Get(ctx, owner, batch.Key)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get %s balance for %s: %s", batch.Denom, owner.String(), err.Error())
		}
		batchSupply, err := k.stateStore.BatchSupplyTable().Get(ctx, batch.Key)
		if err != nil {
			return nil, err
		}
		decs, err := utils.GetNonNegativeFixedDecs(creditType.Precision, credit.Amount, batchSupply.TradableAmount, userBalance.TradableAmount, batchSupply.CancelledAmount)
		if err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrap(err.Error())
		}
		amtToCancelDec, supplyTradable, userBalTradable, cancelledDec := decs[0], decs[1], decs[2], decs[3]
		userBalTradable, err = math.SafeSubBalance(userBalTradable, amtToCancelDec)
		if err != nil {
			return nil, ecocredit.ErrInsufficientCredits.Wrapf(
				"tradable balance: %s, cancel amount %s", decs[2], amtToCancelDec,
			)
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
			Address:        owner,
			TradableAmount: userBalTradable.String(),
			RetiredAmount:  userBalance.RetiredAmount,
			EscrowedAmount: userBalance.EscrowedAmount,
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

		if err = sdkCtx.EventManager().EmitTypedEvent(&types.EventCancel{
			Owner:      owner.String(),
			BatchDenom: credit.BatchDenom,
			Amount:     credit.Amount,
			Reason:     req.Reason,
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/core/MsgCancel credit iteration")
	}
	return &types.MsgCancelResponse{}, nil
}
