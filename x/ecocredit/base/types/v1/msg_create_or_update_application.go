package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgCreateOrUpdateApplication{}

// Route implements the LegacyMsg interface.
func (m *MsgCreateOrUpdateApplication) Route() string { return sdk.MsgTypeURL(m) }

// Type implements the LegacyMsg interface.
func (m *MsgCreateOrUpdateApplication) Type() string { return sdk.MsgTypeURL(m) }

// GetSignBytes implements the LegacyMsg interface.
func (m *MsgCreateOrUpdateApplication) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateOrUpdateApplication) ValidateBasic() error {
	return nil
}

// GetSigners returns the expected signers for MsgCreateOrUpdateApplication.
func (m *MsgCreateOrUpdateApplication) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.ProjectAdmin)}
}
