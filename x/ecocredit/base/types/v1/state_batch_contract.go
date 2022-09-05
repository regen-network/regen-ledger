package v1

import (
	"github.com/regen-network/regen-ledger/types/eth"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs basic validation of the BatchContract state type
func (m *BatchContract) Validate() error {
	if m.BatchKey == 0 {
		return ecocredit.ErrParseFailure.Wrapf("batch key cannot be zero")
	}

	if m.ClassKey == 0 {
		return ecocredit.ErrParseFailure.Wrapf("class key cannot be zero")
	}

	if !eth.IsValidAddress(m.Contract) {
		return ecocredit.ErrParseFailure.Wrapf("contract must be a valid ethereum address")
	}

	return nil
}
