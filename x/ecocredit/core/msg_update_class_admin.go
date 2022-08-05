package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgUpdateClassAdmin{}

func (m MsgUpdateClassAdmin) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassAdmin) Type() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassAdmin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

func (m *MsgUpdateClassAdmin) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("admin: %s", err)
	}

	if err := ValidateClassId(m.ClassId); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	if _, err := sdk.AccAddressFromBech32(m.NewAdmin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("new admin: %s", err)
	}

	if m.Admin == m.NewAdmin {
		return sdkerrors.ErrInvalidRequest.Wrap("admin and new admin cannot be the same")
	}

	return nil
}

func (m *MsgUpdateClassAdmin) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}
