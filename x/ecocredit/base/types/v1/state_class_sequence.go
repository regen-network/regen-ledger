package v1

import (
	"cosmossdk.io/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
)

// Validate performs basic validation of the ClassSequence state type
func (m *ClassSequence) Validate() error {
	if err := base.ValidateCreditTypeAbbreviation(m.CreditTypeAbbrev); err != nil {
		return errors.Wrap(err, "credit type abbrev") // returns parse error
	}

	if m.NextSequence == 0 {
		return ecocredit.ErrParseFailure.Wrap("next sequence cannot be zero")
	}

	return nil
}
