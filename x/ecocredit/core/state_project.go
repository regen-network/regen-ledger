package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs a basic validation of project
func (m *Project) Validate() error {
	if m.Key == 0 {
		return ecocredit.ErrParseFailure.Wrap("key cannot be zero")
	}

	if err := ValidateProjectId(m.Id); err != nil {
		return err // returns parse error
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Admin).String()); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("admin: %s", err)
	}

	if m.ClassKey == 0 {
		return ecocredit.ErrParseFailure.Wrap("class key cannot be zero")
	}

	if err := ValidateJurisdiction(m.Jurisdiction); err != nil {
		return err // returns parse error
	}

	if len(m.Metadata) > MaxMetadataLength {
		return ecocredit.ErrParseFailure.Wrapf("metadata exceeds 256 characters")
	}

	return nil
}
