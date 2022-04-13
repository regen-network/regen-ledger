package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgCreateBatch{}

// Route implements the LegacyMsg interface.
func (m MsgCreateBatch) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCreateBatch) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCreateBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateBatch) ValidateBasic() error {

	if len(m.Metadata) > MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrap("credit batch metadata")
	}

	if _, err := sdk.AccAddressFromBech32(m.Issuer); err != nil {
		return sdkerrors.Wrap(err, "issuer")
	}

	if m.StartDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("must provide a start date for the credit batch")
	}
	if m.EndDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("must provide an end date for the credit batch")
	}
	if m.EndDate.Before(*m.StartDate) {
		return sdkerrors.ErrInvalidRequest.Wrapf("the batch end date (%s) must be the same as or after the batch start date (%s)", m.EndDate.Format("2006-01-02"), m.StartDate.Format("2006-01-02"))
	}

	if err := ValidateProjectID(m.ProjectId); err != nil {
		return err
	}

	if err := validateBatchIssuances(m.Issuance); err != nil {
		return err
	}
	if err := validateOriginTx(m.OriginTx, false); err != nil {
		return err
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateBatch.
func (m *MsgCreateBatch) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Issuer)
	return []sdk.AccAddress{addr}
}
