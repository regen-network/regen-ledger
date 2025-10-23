package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgAddClassCreator{}

// Route implements the LegacyMsg interface.
func (m MsgAddClassCreator) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgAddClassCreator) Type() string { return sdk.MsgTypeURL(&m) }

// GetSigners returns the expected signers for MsgAddClassCreator.
func (m *MsgAddClassCreator) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
