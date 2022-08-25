package v3

import (
	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// MigrateState performs in-place store migrations from ConsensusVersion 2 to 3.
func MigrateState(sdkCtx sdk.Context, ss api.StateStore, basketStore basketapi.StateStore, subspace paramtypes.Subspace) error {
	// TODO: migrate core params

	// migrate basket params
	var params core.Params
	subspace.GetParamSet(sdkCtx, &params)

	// verify basket fee is valid
	if err := params.BasketFee.Validate(); err != nil {
		return err
	}

	basketFees := []*basev1beta1.Coin{}
	for _, coin := range params.BasketFee {
		basketFees = append(basketFees, &basev1beta1.Coin{
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
