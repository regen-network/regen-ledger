package data

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_, _ sdk.Msg = &MsgAnchorData{}, &MsgSignData{}
)

func (m *MsgAnchorData) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}
	if m.Hash == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("hash cannot be empty")
	}
	return m.Hash.Validate()
}

func (m *MsgAnchorData) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m *MsgSignData) ValidateBasic() error {
	for _, addr := range m.Signers {
		if _, err := sdk.AccAddressFromBech32(addr); err != nil {
			return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
		}
	}
	if m.Hash == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("hash cannot be empty")
	}
	return m.Hash.Validate()
}

func (m *MsgSignData) GetSigners() []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, len(m.Signers))

	for i, signer := range m.Signers {
		addr, _ := sdk.AccAddressFromBech32(signer)
		addrs[i] = addr
	}

	return addrs
}
