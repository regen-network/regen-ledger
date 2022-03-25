package marketplace

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// assertHasBalance checks that the account has `qty` credits from the given batch id.
func assertHasBalance(ctx context.Context, store ecoApi.StateStore, acc sdk.AccAddress, batchId uint64, qty math.Dec) error {
	res, err := store.BatchBalanceTable().Get(ctx, acc, batchId)
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
func isDenomAllowed(ctx context.Context, store api.StateStore, denom string) (bool, error) {
	return store.AllowedDenomTable().Has(ctx, denom)
}
