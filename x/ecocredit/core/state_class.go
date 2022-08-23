package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs basic validation of the CreditClass state type
func (m *Class) Validate() error {
	if m.Key == 0 {
		return ecocredit.ErrParseFailure.Wrapf("key cannot be zero")
	}

	if err := ValidateClassID(m.Id); err != nil {
		return err // returns parse error
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Admin).String()); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("admin: %s", err)
	}

	if len(m.Metadata) > MaxMetadataLength {
		return ecocredit.ErrParseFailure.Wrap("credit class metadata cannot be more than 256 characters")
	}

	if err := ValidateCreditTypeAbbreviation(m.CreditTypeAbbrev); err != nil {
		return err // returns parse error
	}

	return nil
}
