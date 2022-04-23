package axelarbridge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRecordBridgeEvent{}

func (m *MsgRecordBridgeEvent) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.BridgeAccount)
	return []sdk.AccAddress{addr}
}

func (m *MsgRecordBridgeEvent) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.BridgeAccount); err != nil {
		return sdkerrors.Wrap(err, "malformed signer account")
	}

	return nil
}
