package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs basic validation of the Batch state type
func (b Batch) Validate() error {
	if err := ValidateBatchDenom(b.Denom); err != nil {
		return err // returns parse error
	}

	if b.ProjectKey == 0 {
		return ecocredit.ErrParseFailure.Wrapf("project key cannot be zero")
	}

	if b.StartDate == nil {
		return ecocredit.ErrParseFailure.Wrapf("must provide a start date for the credit batch")
	}

	if b.EndDate == nil {
		return ecocredit.ErrParseFailure.Wrapf("must provide an end date for the credit batch")
	}

	if b.EndDate.Compare(*b.StartDate) != 1 {
		return ecocredit.ErrParseFailure.Wrapf(
			"the batch end date (%s) must be the same as or after the batch start date (%s)",
			b.EndDate.String(),
			b.StartDate.String(),
		)
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(b.Issuer).String()); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("issuer: %s", err)
	}

	return nil
}
