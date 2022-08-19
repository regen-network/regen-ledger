package core

import (
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs basic validation of the BatchSupply state type
func (m *BatchSupply) Validate() error {
	if m.BatchKey == 0 {
		return ecocredit.ErrParseFailure.Wrapf("batch key cannot be zero")
	}

	if _, err := math.NewNonNegativeDecFromString(m.TradableAmount); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("tradable amount: %s", err)
	}

	if _, err := math.NewNonNegativeDecFromString(m.RetiredAmount); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("retired amount: %s", err)
	}

	if _, err := math.NewNonNegativeDecFromString(m.CancelledAmount); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("cancelled amount: %s", err)
	}

	return nil
}
