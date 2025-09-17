package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgSetClassCreatorAllowlist{}

// Route implements the LegacyMsg interface.
func (m MsgSetClassCreatorAllowlist) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgSetClassCreatorAllowlist) Type() string { return sdk.MsgTypeURL(&m) }

// GetSigners returns the expected signers for MsgSetClassCreatorAllowlist.
func (m *MsgSetClassCreatorAllowlist) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
