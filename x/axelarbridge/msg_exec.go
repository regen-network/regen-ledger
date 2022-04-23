package axelarbridge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgExecBridgeEvent{}

func (m *MsgExecBridgeEvent) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Executor)
	return []sdk.AccAddress{addr}
}

func (m *MsgExecBridgeEvent) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Executor); err != nil {
		return sdkerrors.Wrap(err, "sender")
	}

	return nil
}
