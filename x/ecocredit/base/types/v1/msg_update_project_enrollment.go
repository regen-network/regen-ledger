package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgUpdateProjectEnrollment{}

// Route implements the LegacyMsg interface.
func (m *MsgUpdateProjectEnrollment) Route() string { return sdk.MsgTypeURL(m) }

// Type implements the LegacyMsg interface.
func (m *MsgUpdateProjectEnrollment) Type() string { return sdk.MsgTypeURL(m) }

// ValidateBasic does a sanity check on the provided data.
func (m *MsgUpdateProjectEnrollment) ValidateBasic() error {
	panic("implement me")
}

// GetSigners returns the expected signers for MsgUpdateProjectEnrollment.
func (m *MsgUpdateProjectEnrollment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Issuer)}
}
