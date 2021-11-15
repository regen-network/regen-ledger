package data

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_, _ sdk.Msg = &MsgAnchorData{}, &MsgSignData{}
)

func (m *MsgAnchorData) ValidateBasic() error {
	return m.Hash.Validate()
}

func (m *MsgAnchorData) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgSignData) ValidateBasic() error {
	return m.Hash.Validate()
}

func (m *MsgSignData) GetSigners() []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, len(m.Signers))

	for i, signer := range m.Signers {
		addr, err := sdk.AccAddressFromBech32(signer)
		if err != nil {
			panic(err)
		}
		addrs[i] = addr
	}

	return addrs
}
