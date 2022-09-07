package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var (
	_ legacytx.LegacyMsg = &MsgSubmitTx{}
)

func (m MsgSubmitTx) ValidateBasic() error {
	if m.Owner == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("owner cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("owner: %s", err.Error())
	}
	if m.ConnectionId == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("connection_id cannot be empty")
	}
	if m.Msg == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("msg is required")
	}
	return nil
}

func (m MsgSubmitTx) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}

func (m MsgSubmitTx) GetSignBytes() []byte {
	return ModuleCdc.MustMarshalJSON(&m)
}

func (m MsgSubmitTx) Route() string {
	return sdk.MsgTypeURL(&m)
}

func (m MsgSubmitTx) Type() string {
	return sdk.MsgTypeURL(&m)
}
