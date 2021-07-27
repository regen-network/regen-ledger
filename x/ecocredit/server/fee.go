package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s serverImpl) getCreditClassFee(ctx sdk.Context) sdk.Coins {
	var params ecocredit.Params
	s.paramSpace.GetParamSet(ctx, &params)
	return params.CreditClassFee
}

func (s serverImpl) chargeCreditClassFee(ctx sdk.Context, designerAddr sdk.AccAddress) error {
	creditClassFee := s.getCreditClassFee(ctx)

	// Move the fee to the ecocredit module's account
	err := s.bankKeeper.SendCoinsFromAccountToModule(ctx, designerAddr, ecocredit.ModuleName, creditClassFee)
	if err != nil {
		return err
	}

	// Burn the coins
	// TODO: Update this implementation based on the discussion at
	// https://github.com/regen-network/regen-ledger/issues/351
	err = s.bankKeeper.BurnCoins(ctx, ecocredit.ModuleName, creditClassFee)
	if err != nil {
		return err
	}

	return nil
}
