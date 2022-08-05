package data

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgRegisterResolver{}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgRegisterResolver) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Manager); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}

	if m.ResolverId == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("resolver id cannot be empty")
	}

	if len(m.ContentHashes) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("content hashes cannot be empty")
	}

	for _, ch := range m.ContentHashes {
		if err := ch.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgRegisterResolver.
func (m *MsgRegisterResolver) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Manager)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgRegisterResolver) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgRegisterResolver) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgRegisterResolver) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
