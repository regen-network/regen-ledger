package data

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"net/url"
)

var (
	_, _, _, _ sdk.Msg = &MsgAnchor{}, &MsgAttest{}, &MsgDefineResolver{}, &MsgRegisterResolver{}
)

func (m *MsgAnchor) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}
	if m.Hash == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("hash cannot be empty")
	}
	return m.Hash.Validate()
}

func (m *MsgAnchor) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m *MsgAttest) ValidateBasic() error {
	for _, addr := range m.Attestors {
		if _, err := sdk.AccAddressFromBech32(addr); err != nil {
			return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
		}
	}
	if m.Hash == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("hash cannot be empty")
	}
	return m.Hash.Validate()
}

func (m *MsgAttest) GetSigners() []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, len(m.Attestors))

	for i, attestor := range m.Attestors {
		addr, _ := sdk.AccAddressFromBech32(attestor)
		addrs[i] = addr
	}

	return addrs
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

func (m *MsgRegisterResolver) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Manager); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}
	if len(m.Data) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("data cannot be empty")
	}
	for _, hash := range m.Data {
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
