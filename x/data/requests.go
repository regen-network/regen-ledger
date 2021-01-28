package data

import sdk "github.com/cosmos/cosmos-sdk/types"

var (
	_, _, _ sdk.MsgRequest = &MsgAnchorDataRequest{}, &MsgSignDataRequest{}, &MsgStoreRawDataRequest{}
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

	for i, signer := range m.Signers {
		addr, err := sdk.AccAddressFromBech32(signer)
		if err != nil {
			panic(err)
		}
		addrs[i] = addr
	}

	return addrs
}

func (m *MsgStoreRawDataRequest) ValidateBasic() error {
	return nil
}

func (m *MsgStoreRawDataRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}
