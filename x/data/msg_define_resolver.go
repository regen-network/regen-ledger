package data

import (
	"net/url"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgDefineResolver{}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgDefineResolver) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Manager); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}

	if _, err := url.ParseRequestURI(m.ResolverUrl); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap("invalid resolver url")
	}

	return nil
}

// GetSigners returns the expected signers for MsgDefineResolver.
func (m *MsgDefineResolver) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Manager)
	return []sdk.AccAddress{addr}
}

// LegacyMsg.Type implementations
func (msg MsgDefineResolver) Route() string { return "" }
func (msg MsgDefineResolver) Type() string  { return sdk.MsgTypeURL(&msg) }
func (msg *MsgDefineResolver) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
