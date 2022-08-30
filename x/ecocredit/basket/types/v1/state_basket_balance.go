package v1

import (
	"cosmossdk.io/errors"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
)

// Validate performs basic validation of the BasketBalance state type
func (m *BasketBalance) Validate() error {
	if m.BasketId == 0 {
		return ecocredit.ErrParseFailure.Wrapf("basket id cannot be zero")
	}

	if err := base.ValidateBatchDenom(m.BatchDenom); err != nil {
		return errors.Wrap(err, "batch denom") // returns parse error
	}

	if _, err := math.NewNonNegativeDecFromString(m.Balance); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("balance: %s", err)
	}

	if m.BatchStartDate.GetSeconds() == 0 && m.BatchStartDate.GetNanos() == 0 {
		return ecocredit.ErrParseFailure.Wrap("batch start date cannot be empty")
	}

	return nil
}
