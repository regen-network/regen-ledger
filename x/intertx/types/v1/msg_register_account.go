package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var (
	_ legacytx.LegacyMsg = &MsgRegisterAccount{}
)

func (m MsgRegisterAccount) ValidateBasic() error {
	if m.Owner == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("owner cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("owner: %s", err.Error())
	}

	if m.Version == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("version cannot be empty")
	}

	if m.ConnectionId == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("connection_id cannot be empty")
	}

	return nil
}

func (m MsgRegisterAccount) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}

func (m MsgRegisterAccount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgRegisterAccount) Route() string {
	return sdk.MsgTypeURL(&m)
}

func (m MsgRegisterAccount) Type() string {
	return sdk.MsgTypeURL(&m)
}
