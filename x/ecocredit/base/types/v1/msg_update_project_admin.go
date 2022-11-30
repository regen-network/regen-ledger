package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
)

var _ legacytx.LegacyMsg = &MsgUpdateProjectAdmin{}

func (m MsgUpdateProjectAdmin) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateProjectAdmin) Type() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateProjectAdmin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgUpdateProjectAdmin) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("admin: %s", err)
	}

	if err := base.ValidateProjectID(m.ProjectId); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("project id: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.NewAdmin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("new admin: %s", err)
	}

	if m.Admin == m.NewAdmin {
		return sdkerrors.ErrInvalidRequest.Wrap("admin and new admin cannot be the same")
	}

	return nil
}

func (m MsgUpdateProjectAdmin) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}
