package data

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgAnchor{}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgAnchor) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}

	if m.ContentHash == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
	}

	return m.ContentHash.Validate()
}

// GetSigners returns the expected signers for MsgAnchor.
func (m *MsgAnchor) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// LegacyMsg.Type implementations
func (m MsgAnchor) Route() string { return "" }
func (m MsgAnchor) Type() string  { return sdk.MsgTypeURL(&m) }
func (m *MsgAnchor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
