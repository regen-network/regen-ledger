package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgCreateOrUpdateApplication{}

// Route implements the LegacyMsg interface.
func (m *MsgCreateOrUpdateApplication) Route() string { return sdk.MsgTypeURL(m) }

// Type implements the LegacyMsg interface.
func (m *MsgCreateOrUpdateApplication) Type() string { return sdk.MsgTypeURL(m) }

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateOrUpdateApplication) ValidateBasic() error {
	panic("implement me")
}

// GetSigners returns the expected signers for MsgCreateOrUpdateApplication.
func (m *MsgCreateOrUpdateApplication) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.ProjectAdmin)}
}
