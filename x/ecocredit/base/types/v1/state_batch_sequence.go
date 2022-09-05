package v1

import (
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs basic validation of the BatchSequence state type
func (m *BatchSequence) Validate() error {
	if m.ProjectKey == 0 {
		return ecocredit.ErrParseFailure.Wrapf("project key cannot be zero")
	}

	if m.NextSequence == 0 {
		return ecocredit.ErrParseFailure.Wrapf("next sequence cannot be zero")
	}

	return nil
}
