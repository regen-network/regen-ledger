package marketplace

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// assertHasBalance checks that the account has `qty` credits from the given batch id.
func assertHasBalance(ctx context.Context, store ecocreditv1.StateStore, acc sdk.AccAddress, batchId uint64, qty math.Dec) error {
	res, err := store.BatchBalanceStore().Get(ctx, acc, batchId)
	if err != nil {
		return err
	}
	tradableBalance, err := math.NewDecFromString(res.Tradable)
	if err != nil {
		return err
	}
	if tradableBalance.Cmp(qty) == -1 {
		return ecocredit.ErrInsufficientFunds.Wrapf("cannot create a sell order of %s credits with a balance of %s", qty.String(), res.Tradable)
	}
	return nil
}

// isDenomAllowed checks if the denom is allowed to be used in orders.
func isDenomAllowed(ctx context.Context, store marketplacev1.StateStore, denom string) (bool, error) {
	return store.AllowedDenomStore().Has(ctx, denom)
}
