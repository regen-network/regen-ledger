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
	if _, err := sdk.AccAddressFromBech32(m.Issuer); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("issuer: %s", err)
	}

	if err := ValidateProjectId(m.ProjectId); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	if len(m.Issuance) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("issuance cannot be empty")
	}

	for i, issuance := range m.Issuance {
		if err := issuance.Validate(); err != nil {
			return sdkerrors.Wrapf(err, "issuance[%d]", i)
		}
	}

	if len(m.Metadata) > MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrapf("metadata: max length %d", MaxMetadataLength)
	}

	if m.StartDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("start date cannot be empty")
	}

	if m.EndDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("end date cannot be empty")
	}

	if m.StartDate.After(*m.EndDate) {
		return sdkerrors.ErrInvalidRequest.Wrap("start date cannot be after end date")
	}

	// origin tx is not required when creating a credit batch
	if m.OriginTx != nil {
		if err := m.OriginTx.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateBatch.
func (m *MsgCreateBatch) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Issuer)
	return []sdk.AccAddress{addr}
}
