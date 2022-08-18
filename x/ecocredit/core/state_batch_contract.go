package core

import (
	"github.com/regen-network/regen-ledger/types/eth"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs basic validation of the BatchContract state type
func (b BatchContract) Validate() error {
	if b.BatchKey == 0 {
		return ecocredit.ErrParseFailure.Wrapf("batch key cannot be zero")
	}

	if b.ClassKey == 0 {
		return ecocredit.ErrParseFailure.Wrapf("class key cannot be zero")
	}

	if !eth.IsValidAddress(b.Contract) {
		return ecocredit.ErrParseFailure.Wrapf("contract must be a valid ethereum address")
	}

	return nil
}
