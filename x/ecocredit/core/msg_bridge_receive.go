package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgBridgeReceive{}

// Route implements the LegacyMsg interface.
func (m MsgBridgeReceive) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgBridgeReceive) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgBridgeReceive) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgBridgeReceive) ValidateBasic() error {
	// top level fields validation
	if _, err := sdk.AccAddressFromBech32(m.Issuer); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap("issuer")
	}
	if err := ValidateClassId(m.ClassId); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}
	if m.OriginTx == nil {
		sdkerrors.ErrInvalidRequest.Wrap("origin tx cannot be empty")
	}
	if err := m.OriginTx.Validate(); err != nil {
		return err
	}

	// batch validation
	if m.Batch == nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("batch cannot be empty")
	}
	batch := m.Batch
	if _, err := sdk.AccAddressFromBech32(batch.Recipient); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap("recipient")
	}
	if _, err := math.NewPositiveDecFromString(batch.Amount); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}
	if batch.StartDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("start_date is required")
	}
	if batch.EndDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("end_date is required")
	}
	if batch.StartDate.After(*batch.EndDate) {
		return sdkerrors.ErrInvalidRequest.Wrap("start_date must be a time before end_date")
	}
	if len(batch.Metadata) > MaxMetadataLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("batch metadata length (%d) exceeds max metadata length: %d", len(batch.Metadata), MaxMetadataLength)
	}

	// project validation
	if m.Project == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("project cannot be empty")
	}
	project := m.Project
	if len(project.ReferenceId) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("reference_id is required")
	}
	if err := ValidateJurisdiction(project.Jurisdiction); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}
	if len(project.Metadata) > MaxMetadataLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("project_metadata length (%d) exceeds max metadata length: %d", len(project.Metadata), MaxMetadataLength)
	}
	return nil
}

// GetSigners returns the expected signers for MsgCancel.
func (m *MsgBridgeReceive) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Issuer)
	return []sdk.AccAddress{addr}
}
