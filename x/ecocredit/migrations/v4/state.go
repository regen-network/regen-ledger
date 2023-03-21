package v4

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	basketv1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
)

// MigrateState performs in-place store migrations from ConsensusVersion 3 to 4.
func MigrateState(sdkCtx sdk.Context, basketStore basketv1.StateStore) error {
	if sdkCtx.ChainID() == "regen-1" {
		if err := migrateBasket(sdkCtx, basketStore); err != nil {
			return err
		}
	}

	return nil
}

func migrateBasket(ctx sdk.Context, basketStore basketv1.StateStore) error {
	b, err := basketStore.BasketTable().GetByBasketDenom(ctx, "eco.uC.NCT")
	if err != nil {
		return err
	}

	b.DisableAutoRetire = true

	b.DateCriteria = nil // TODO

	return basketStore.BasketTable().Update(ctx, b)
}
