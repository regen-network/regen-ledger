package core

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Retire credits to the specified jurisdiction.
// WARNING: retiring credits is permanent. Retired credits cannot be un-retired.
func (k Keeper) Retire(ctx context.Context, req *core.MsgRetire) (*core.MsgRetireResponse, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	holder, _ := sdk.AccAddressFromBech32(req.Holder)

	for _, credit := range req.Credits {
		batch, err := k.stateStore.BatchInfoTable().GetByBatchDenom(ctx, credit.BatchDenom)
		if err != nil {
			return nil, err
		}
		creditType, err := utils.GetCreditTypeFromBatchDenom(ctx, k.stateStore, k.paramsKeeper, batch.BatchDenom)
		if err != nil {
			return nil, err
		}
		userBalance, err := k.stateStore.BatchBalanceTable().Get(ctx, holder, batch.Key)
		if err != nil {
			return nil, err
		}

		decs, err := utils.GetNonNegativeFixedDecs(creditType.Precision, credit.Amount, userBalance.Tradable)
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
		batchSupply, err := k.stateStore.BatchSupplyTable().Get(ctx, batch.Key)
		if err != nil {
			return nil, err
		}
		decs, err = utils.GetNonNegativeFixedDecs(creditType.Precision, batchSupply.RetiredAmount, batchSupply.TradableAmount)
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

		if err = k.stateStore.BatchBalanceTable().Update(ctx, &api.BatchBalance{
			BatchKey: batch.Key,
			Address:  holder,
			Tradable: userTradableBalance.String(),
			Retired:  userRetiredBalance.String(),
		}); err != nil {
			return nil, err
		}

		err = k.stateStore.BatchSupplyTable().Update(ctx, &api.BatchSupply{
			BatchKey:        batch.Key,
			TradableAmount:  supplyTradable.String(),
			RetiredAmount:   supplyRetired.String(),
			CancelledAmount: batchSupply.CancelledAmount,
		})

		if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventRetire{
			Retirer:      req.Holder,
			BatchDenom:   credit.BatchDenom,
			Amount:       credit.Amount,
			Jurisdiction: req.Jurisdiction,
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/core/MsgRetire credit iteration")
	}
	return &core.MsgRetireResponse{}, nil
}
