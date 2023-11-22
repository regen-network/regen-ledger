package v1

import (
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
)

// Validate performs basic validation of the ProjectSequence state type
func (m *ProjectSequence) Validate() error {
	if m.ClassKey == 0 {
		return ecocredit.ErrParseFailure.Wrap("class key cannot be zero")
	}

	if m.NextSequence == 0 {
		return ecocredit.ErrParseFailure.Wrap("next sequence cannot be zero")
	}

	return nil
}
