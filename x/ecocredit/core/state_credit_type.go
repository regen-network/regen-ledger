package core

import (
	"cosmossdk.io/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

const (
	maxCreditTypeNameLength = 75
)

// Validate performs basic validation of the ClassIssuer state type
func (m *CreditType) Validate() error {
	if err := ValidateCreditTypeAbbreviation(m.Abbreviation); err != nil {
		return errors.Wrapf(err, "abbreviation") // returns parse error
	}

	if len(m.Name) == 0 {
		return ecocredit.ErrParseFailure.Wrap("name cannot be empty")
	}

	if len(m.Name) > maxCreditTypeNameLength {
		return ecocredit.ErrParseFailure.Wrapf("credit type name cannot exceed %d characters", maxCreditTypeNameLength)
	}

	if len(m.Unit) == 0 {
		return ecocredit.ErrParseFailure.Wrap("unit cannot be empty")
	}

	if m.Precision != PRECISION {
		return ecocredit.ErrParseFailure.Wrapf("precision is currently locked to %d", PRECISION)
	}

	return nil
}
