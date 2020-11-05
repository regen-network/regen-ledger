package data

import sdk "github.com/cosmos/cosmos-sdk/types"

var (
	_, _, _ sdk.MsgRequest = &MsgAnchorDataRequest{}, &MsgSignDataRequest{}, &MsgStoreDataRequest{}
)

func (m *MsgAnchorDataRequest) ValidateBasic() error {
	return nil
}

func (m *MsgAnchorDataRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgSignDataRequest) ValidateBasic() error {
	return nil
}

func (m *MsgSignDataRequest) GetSigners() []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, len(m.Signers))

	for _, signer := range m.Signers {
		addr, err := sdk.AccAddressFromBech32(signer)
		if err != nil {
			panic(err)
		}
		addrs = append(addrs, addr)
	}

	return addrs
}

func (m *MsgStoreDataRequest) ValidateBasic() error {
	return nil
}

func (m *MsgStoreDataRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}
