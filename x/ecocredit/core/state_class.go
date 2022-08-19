package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs basic validation of the CreditClass state type
func (c Class) Validate() error {
	if c.Key == 0 {
		return ecocredit.ErrParseFailure.Wrapf("key cannot be zero")
	}

	if err := ValidateClassId(c.Id); err != nil {
		return err // returns parse error
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(c.Admin).String()); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("admin: %s", err)
	}

	if len(c.Metadata) > MaxMetadataLength {
		return ecocredit.ErrParseFailure.Wrap("credit class metadata cannot be more than 256 characters")
	}

	if err := ValidateCreditTypeAbbreviation(c.CreditTypeAbbrev); err != nil {
		return err // returns parse error
	}

	return nil
}
