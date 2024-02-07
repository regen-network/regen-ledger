package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgCreateUnregisteredProject{}

// Route implements the LegacyMsg interface.
func (m *MsgCreateUnregisteredProject) Route() string { return sdk.MsgTypeURL(m) }

// Type implements the LegacyMsg interface.
func (m *MsgCreateUnregisteredProject) Type() string { return sdk.MsgTypeURL(m) }

// GetSignBytes implements the LegacyMsg interface.
func (m *MsgCreateUnregisteredProject) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateUnregisteredProject) ValidateBasic() error {
	panic("implement me")
}

// GetSigners returns the expected signers for MsgCreateUnregisteredProject.
func (m *MsgCreateUnregisteredProject) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Admin)}
}
