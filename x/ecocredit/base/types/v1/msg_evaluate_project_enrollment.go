package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgEvaluateProjectEnrollment{}

// Route implements the LegacyMsg interface.
func (m *MsgEvaluateProjectEnrollment) Route() string { return sdk.MsgTypeURL(m) }

// Type implements the LegacyMsg interface.
func (m *MsgEvaluateProjectEnrollment) Type() string { return sdk.MsgTypeURL(m) }

// GetSignBytes implements the LegacyMsg interface.
func (m *MsgEvaluateProjectEnrollment) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgEvaluateProjectEnrollment) ValidateBasic() error {
	panic("implement me")
}

// GetSigners returns the expected signers for MsgEvaluateProjectEnrollment.
func (m *MsgEvaluateProjectEnrollment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Issuer)}
}
