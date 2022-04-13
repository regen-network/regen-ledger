package data

import (
	"net/url"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

var (
	_, _, _, _ legacytx.LegacyMsg = &MsgAnchor{}, &MsgAttest{}, &MsgDefineResolver{}, &MsgRegisterResolver{}
)

func (m *MsgAnchor) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}
	if m.ContentHash == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
	}
	return m.ContentHash.Validate()
}

func (m *MsgAnchor) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgAnchor) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgAnchor) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgAnchor) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m *MsgAttest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Attestor); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}

	if len(m.ContentHashes) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("content hashes cannot be empty")
	}

	for _, hash := range m.ContentHashes {
		if hash == nil {
			return sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
		}
		err := hash.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MsgAttest) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Attestor)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgAttest) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgAttest) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgAttest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m *MsgDefineResolver) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Manager); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}

	if _, err := url.ParseRequestURI(m.ResolverUrl); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap("invalid resolver url")
	}

	return nil
}

func (m *MsgDefineResolver) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Manager)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgDefineResolver) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgDefineResolver) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgDefineResolver) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m *MsgRegisterResolver) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Manager); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}
	if len(m.ContentHashes) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("content hashes cannot be empty")
	}
	for _, hash := range m.ContentHashes {
		if err := hash.Validate(); err != nil {
			return err
		}
	}
	return nil
}

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
