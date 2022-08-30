package v1

import (
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
)

// Validate performs basic validation of the CreditClass state type
func (m *Class) Validate() error {
	if m.Key == 0 {
		return ecocredit.ErrParseFailure.Wrapf("key cannot be zero")
	}

	if err := base.ValidateClassID(m.Id); err != nil {
		return errors.Wrap(err, "class id") // returns parse error
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Admin).String()); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("admin: %s", err)
	}

	if len(m.Metadata) > base.MaxMetadataLength {
		return ecocredit.ErrParseFailure.Wrap("credit class metadata cannot be more than 256 characters")
	}

	if err := base.ValidateCreditTypeAbbreviation(m.CreditTypeAbbrev); err != nil {
		return errors.Wrap(err, "credit type abbrev") // returns parse error
	}

	return nil
}
