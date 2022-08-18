package core

import (
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs basic validation of the BatchSequence state type
func (b BatchSequence) Validate() error {
	if b.ProjectKey == 0 {
		return ecocredit.ErrParseFailure.Wrapf("project key cannot be zero")
	}

	if b.NextSequence == 0 {
		return ecocredit.ErrParseFailure.Wrapf("next sequence cannot be zero")
	}

	return nil
}
