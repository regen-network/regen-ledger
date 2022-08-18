package core

import (
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs basic validation of the BatchSupply state type
func (b BatchSupply) Validate() error {
	if b.BatchKey == 0 {
		return ecocredit.ErrParseFailure.Wrapf("batch key cannot be zero")
	}

	if _, err := math.NewNonNegativeDecFromString(b.TradableAmount); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("tradable amount: %s", err)
	}

	if _, err := math.NewNonNegativeDecFromString(b.RetiredAmount); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("retired amount: %s", err)
	}

	if _, err := math.NewNonNegativeDecFromString(b.CancelledAmount); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("cancelled amount: %s", err)
	}

	return nil
}
