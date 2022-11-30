package v3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	basketapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	baseapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// MigrateState performs in-place store migrations from ConsensusVersion 2 to 3.
func MigrateState(sdkCtx sdk.Context, baseStore baseapi.StateStore, basketStore basketapi.StateStore, subspace paramtypes.Subspace) error {
	var params basetypes.Params
	subspace.GetParamSet(sdkCtx, &params)

	// validate credit class fee
	if err := params.CreditClassFee.Validate(); err != nil {
		return err
	}

	// migrate credit class fees
	classFees := regentypes.CoinsToProtoCoins(params.CreditClassFee)
	if err := baseStore.ClassFeeTable().Save(sdkCtx, &baseapi.ClassFee{
		Fee: classFees[0], // we assume there is one fee at the time of the upgrade
	}); err != nil {
		return err
	}

	// migrate credit class allow list
	if err := baseStore.ClassCreatorAllowlistTable().Save(sdkCtx, &baseapi.ClassCreatorAllowlist{
		Enabled: params.AllowlistEnabled,
	}); err != nil {
		return err
	}

	// migrate allowed class creators to orm table
	for _, creator := range params.AllowedClassCreators {
		addr, err := sdk.AccAddressFromBech32(creator)
		if err != nil {
			return err
		}

		if err := baseStore.AllowedClassCreatorTable().Save(sdkCtx, &baseapi.AllowedClassCreator{
			Address: addr,
		}); err != nil {
			return err
		}
	}

	// verify basket fee is valid
	if err := params.BasketFee.Validate(); err != nil {
		return err
	}

	// migrate basket params
	basketFees := regentypes.CoinsToProtoCoins(params.BasketFee)
	if err := basketStore.BasketFeeTable().Save(sdkCtx, &basketapi.BasketFee{
		Fee: basketFees[0], // we assume there is one fee at the time of the upgrade
	}); err != nil {
		return err
	}

	return nil
}
