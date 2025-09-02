package v1

import (
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgSetClassCreatorAllowlist{}

// Route implements the LegacyMsg interface.
func (m MsgSetClassCreatorAllowlist) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgSetClassCreatorAllowlist) Type() string { return sdk.MsgTypeURL(&m) }

// ValidateBasic does a sanity check on the provided data.
func (m *MsgSetClassCreatorAllowlist) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(err, "invalid authority address")
	}

	return nil
}

// GetSigners returns the expected signers for MsgSetClassCreatorAllowlist.
func (m *MsgSetClassCreatorAllowlist) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
