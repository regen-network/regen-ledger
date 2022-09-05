package v1alpha1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgCreateClass{}

// Route implements the LegacyMsg interface.
func (m MsgCreateClass) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCreateClass) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCreateClass) GetSignBytes() []byte {
	return nil
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateClass) ValidateBasic() error {
	return nil
}

// GetSigners returns the expected signers for MsgCreateClass.
func (m *MsgCreateClass) GetSigners() []sdk.AccAddress {
	return nil
}

var _ legacytx.LegacyMsg = &MsgCreateBatch{}

// Route implements the LegacyMsg interface.
func (m MsgCreateBatch) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCreateBatch) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCreateBatch) GetSignBytes() []byte {
	return nil
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateBatch) ValidateBasic() error {
	return nil
}

// GetSigners returns the expected signers for MsgCreateBatch.
func (m *MsgCreateBatch) GetSigners() []sdk.AccAddress {
	return nil
}

var _ legacytx.LegacyMsg = &MsgSend{}

// Route implements the LegacyMsg interface.
func (m MsgSend) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgSend) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgSend) GetSignBytes() []byte {
	return nil
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgSend) ValidateBasic() error {
	return nil
}

// GetSigners returns the expected signers for MsgSend.
func (m *MsgSend) GetSigners() []sdk.AccAddress {
	return nil
}

var _ legacytx.LegacyMsg = &MsgRetire{}

// Route implements the LegacyMsg interface.
func (m MsgRetire) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgRetire) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgRetire) GetSignBytes() []byte {
	return nil
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgRetire) ValidateBasic() error {
	return nil
}

// GetSigners returns the expected signers for MsgRetire.
func (m *MsgRetire) GetSigners() []sdk.AccAddress {
	return nil
}

var _ legacytx.LegacyMsg = &MsgCancel{}

// Route implements the LegacyMsg interface.
func (m MsgCancel) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCancel) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCancel) GetSignBytes() []byte {
	return nil
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCancel) ValidateBasic() error {
	return nil
}

// GetSigners returns the expected signers for MsgCancel.
func (m *MsgCancel) GetSigners() []sdk.AccAddress {
	return nil
}

var _ legacytx.LegacyMsg = &MsgUpdateClassAdmin{}

// Route implements the LegacyMsg interface.
func (m MsgUpdateClassAdmin) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgUpdateClassAdmin) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgUpdateClassAdmin) GetSignBytes() []byte {
	return nil
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgUpdateClassAdmin) ValidateBasic() error {
	return nil
}

// GetSigners returns the expected signers for MsgUpdateClassAdmin.
func (m *MsgUpdateClassAdmin) GetSigners() []sdk.AccAddress {
	return nil
}

var _ legacytx.LegacyMsg = &MsgUpdateClassIssuers{}

// Route implements the LegacyMsg interface.
func (m MsgUpdateClassIssuers) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgUpdateClassIssuers) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgUpdateClassIssuers) GetSignBytes() []byte {
	return nil
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgUpdateClassIssuers) ValidateBasic() error {
	return nil
}

// GetSigners returns the expected signers for MsgUpdateClassIssuers.
func (m *MsgUpdateClassIssuers) GetSigners() []sdk.AccAddress {
	return nil
}

var _ legacytx.LegacyMsg = &MsgUpdateClassMetadata{}

// Route implements the LegacyMsg interface.
func (m MsgUpdateClassMetadata) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgUpdateClassMetadata) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgUpdateClassMetadata) GetSignBytes() []byte {
	return nil
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgUpdateClassMetadata) ValidateBasic() error {
	return nil
}

// GetSigners returns the expected signers for MsgUpdateClassMetadata.
func (m *MsgUpdateClassMetadata) GetSigners() []sdk.AccAddress {
	return nil
}
