package core

import (
	sdkerrors "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs basic validation of the Batch state type
func (m *Batch) Validate() error {
	if m.Key == 0 {
		return ecocredit.ErrParseFailure.Wrapf("key cannot be zero")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Issuer).String()); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("issuer: %s", err)
	}

	if m.ProjectKey == 0 {
		return ecocredit.ErrParseFailure.Wrapf("project key cannot be zero")
	}

	if err := ValidateBatchDenom(m.Denom); err != nil {
		return sdkerrors.Wrap(err, "denom") // returns parse error
	}

	if len(m.Metadata) > MaxMetadataLength {
		return ecocredit.ErrParseFailure.Wrap("metadata cannot be more than 256 characters")
	}

	if m.StartDate == nil {
		return ecocredit.ErrParseFailure.Wrapf("must provide a start date for the credit batch")
	}

	if m.EndDate == nil {
		return ecocredit.ErrParseFailure.Wrapf("must provide an end date for the credit batch")
	}

	if m.EndDate.Compare(*m.StartDate) != 1 {
		return ecocredit.ErrParseFailure.Wrapf(
			"the batch end date (%s) must be the same as or after the batch start date (%s)",
			m.EndDate, m.StartDate,
		)
	}

	if m.IssuanceDate == nil {
		return ecocredit.ErrParseFailure.Wrapf("must provide an issuance date for the credit batch")
	}

	return nil
}
