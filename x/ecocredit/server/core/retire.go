package core

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
)

// Retire credits to the specified location.
// WARNING: retiring credits is permanent. Retired credits cannot be un-retired.
func (k Keeper) Retire(ctx context.Context, req *v1.MsgRetire) (*v1.MsgRetireResponse, error) {
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
		userBalance, err := k.stateStore.BatchBalanceStore().Get(ctx, holder, batch.Id)
		if err != nil {
			return nil, err
		}

		decs, err := getNonNegativeFixedDecs(creditType.Precision, credit.Amount, userBalance.Tradable)
		if err != nil {
			return nil, err
		}
		amtToRetire, userTradableBalance := decs[0], decs[1]

		userTradableBalance, err = math.SafeSubBalance(userTradableBalance, amtToRetire)
		if err != nil {
			return nil, err
		}
		userRetiredBalance, err := math.NewNonNegativeFixedDecFromString(userBalance.Retired, creditType.Precision)
		if err != nil {
			return nil, err
		}
		userRetiredBalance, err = userRetiredBalance.Add(amtToRetire)
		if err != nil {
			return nil, err
		}
		batchSupply, err := k.stateStore.BatchSupplyStore().Get(ctx, batch.Id)
		if err != nil {
			return nil, err
		}
		decs, err = getNonNegativeFixedDecs(creditType.Precision, batchSupply.RetiredAmount, batchSupply.TradableAmount)
		if err != nil {
			return nil, err
		}
		supplyRetired, supplyTradable := decs[0], decs[1]
		supplyRetired, err = supplyRetired.Add(amtToRetire)
		if err != nil {
			return nil, err
		}
		supplyTradable, err = math.SafeSubBalance(supplyTradable, amtToRetire)
		if err != nil {
			return nil, err
		}

		if err = k.stateStore.BatchBalanceStore().Update(ctx, &ecocreditv1.BatchBalance{
			Address:  holder,
			BatchId:  batch.Id,
			Tradable: userTradableBalance.String(),
			Retired:  userRetiredBalance.String(),
		}); err != nil {
			return nil, err
		}
		err = k.stateStore.BatchSupplyStore().Update(ctx, &ecocreditv1.BatchSupply{
			BatchId:         batch.Id,
			TradableAmount:  supplyTradable.String(),
			RetiredAmount:   supplyRetired.String(),
			CancelledAmount: batchSupply.CancelledAmount,
		})
		if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventRetire{
			Retirer:    req.Holder,
			BatchDenom: credit.BatchDenom,
			Amount:     credit.Amount,
			Location:   req.Location,
		}); err != nil {
			return nil, err
		}
		sdkCtx.GasMeter().ConsumeGas(gasCostPerIteration, "retire ecocredits")
	}
	return &v1.MsgRetireResponse{}, nil
}
