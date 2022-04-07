package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

var _ legacytx.LegacyMsg = &MsgMintBatchCredits{}

// Route implements the LegacyMsg interface.
func (m MsgMintBatchCredits) Route() string {
	panic("implement me")
}

// Type implements the LegacyMsg interface.
func (m MsgMintBatchCredits) Type() string {
	panic("implement me")
}

// GetSignBytes implements the LegacyMsg interface.
func (m MsgMintBatchCredits) GetSignBytes() []byte {
	panic("implement me")
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgMintBatchCredits) ValidateBasic() error {
	panic("implement me")
}

// GetSigners returns the expected signers for MsgMintBatchCredits.
func (m *MsgMintBatchCredits) GetSigners() []sdk.AccAddress {
	panic("implement me")
}
