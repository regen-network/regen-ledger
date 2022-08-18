package core

import (
	"fmt"

	"github.com/regen-network/regen-ledger/types/eth"
)

// Validate performs basic validation of the BatchContract state type
func (b BatchContract) Validate() error {
	if b.BatchKey == 0 {
		return fmt.Errorf("batch key cannot be zero")
	}

	if b.ClassKey == 0 {
		return fmt.Errorf("class key cannot be zero")
	}

	if !eth.IsValidAddress(b.Contract) {
		return fmt.Errorf("contract must be a valid ethereum address")
	}

	return nil
}
