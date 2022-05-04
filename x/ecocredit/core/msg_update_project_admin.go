package core

import (
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgUpdateProjectAdmin{}

func (m MsgUpdateProjectAdmin) ValidateBasic() error {
	addr1, err := types.AccAddressFromBech32(m.Admin)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}
	addr2, err := types.AccAddressFromBech32(m.NewAdmin)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}
	if addr1.Equals(addr2) {
		return sdkerrors.ErrInvalidRequest.Wrap("new_admin and admin addresses cannot be the same")
	}
	if err := ValidateProjectId(m.ProjectId); err != nil {
		return err
	}
	return nil
}

func (m MsgUpdateProjectAdmin) GetSigners() []types.AccAddress {
	addr, _ := types.AccAddressFromBech32(m.Admin)
	return []types.AccAddress{addr}
}

func (m MsgUpdateProjectAdmin) GetSignBytes() []byte {
	return types.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgUpdateProjectAdmin) Route() string { return types.MsgTypeURL(&m) }

func (m MsgUpdateProjectAdmin) Type() string { return types.MsgTypeURL(&m) }
