package v1

import (
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
)

// Validate performs a basic validation of project
func (m *Project) Validate() error {
	if m.Key == 0 {
		return ecocredit.ErrParseFailure.Wrap("key cannot be zero")
	}

	if err := base.ValidateProjectID(m.Id); err != nil {
		return errors.Wrap(err, "project id") // returns parse error
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Admin).String()); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("admin: %s", err)
	}

	if m.ClassKey == 0 {
		return ecocredit.ErrParseFailure.Wrap("class key cannot be zero")
	}

	if err := base.ValidateJurisdiction(m.Jurisdiction); err != nil {
		return errors.Wrap(err, "jurisdiction") // returns parse error
	}

	if len(m.Metadata) > base.MaxMetadataLength {
		return ecocredit.ErrParseFailure.Wrapf("metadata exceeds 256 characters")
	}

	return nil
}
