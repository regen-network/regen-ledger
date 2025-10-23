package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/v4/base"
)

var _ sdk.Msg = &MsgUpdateClassAdmin{}

func (m MsgUpdateClassAdmin) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassAdmin) Type() string { return sdk.MsgTypeURL(&m) }

func (m *MsgUpdateClassAdmin) ValidateBasic() error {

	if err := base.ValidateClassID(m.ClassId); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("class id: %s", err)
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
