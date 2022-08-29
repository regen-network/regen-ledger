package v3

import (
	sdkbase "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baseapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// MigrateState performs in-place store migrations from ConsensusVersion 2 to 3.
func MigrateState(sdkCtx sdk.Context, ss baseapi.StateStore, basketStore basketapi.StateStore, subspace paramtypes.Subspace) error {
	// TODO: migrate core params

	// migrate basket params
	var params basetypes.Params
	subspace.GetParamSet(sdkCtx, &params)

	// verify basket fee is valid
	if err := params.BasketFee.Validate(); err != nil {
		return err
	}

	basketFees := []*sdkbase.Coin{}
	for _, coin := range params.BasketFee {
		basketFees = append(basketFees, &sdkbase.Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.String(),
		})
	}
	if err := basketStore.BasketFeesTable().Save(sdkCtx, &basketapi.BasketFees{
		Fees: basketFees,
	}); err != nil {
		return err
	}

	return nil
}
