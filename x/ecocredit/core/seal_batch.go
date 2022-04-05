package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

var _ legacytx.LegacyMsg = &MsgSealBatch{}

// Route implements the LegacyMsg interface.
func (m MsgSealBatch) Route() string {
	panic("implement me")
}

// Type implements the LegacyMsg interface.
func (m MsgSealBatch) Type() string {
	panic("implement me")
}

// GetSignBytes implements the LegacyMsg interface.
func (m MsgSealBatch) GetSignBytes() []byte {
	panic("implement me")
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgSealBatch) ValidateBasic() error {
	panic("implement me")
}

// GetSigners returns the expected signers for MsgSealBatch.
func (m *MsgSealBatch) GetSigners() []sdk.AccAddress {
	panic("implement me")
}
