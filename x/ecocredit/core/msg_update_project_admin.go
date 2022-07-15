package core

import (
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgUpdateProjectAdmin{}

func (m MsgUpdateProjectAdmin) Route() string { return types.MsgTypeURL(&m) }

func (m MsgUpdateProjectAdmin) Type() string { return types.MsgTypeURL(&m) }

func (m MsgUpdateProjectAdmin) GetSignBytes() []byte {
	return types.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgUpdateProjectAdmin) ValidateBasic() error {
	if _, err := types.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("admin: %s", err)
	}

	if m.ProjectId == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("project id cannot be empty")
	}

	if err := ValidateProjectId(m.ProjectId); err != nil {
		return err
	}

	if _, err := types.AccAddressFromBech32(m.NewAdmin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("new admin: %s", err)
	}

	if m.Admin == m.NewAdmin {
		return sdkerrors.ErrInvalidRequest.Wrap("admin and new admin cannot be the same")
	}

	return nil
}

func (m MsgUpdateProjectAdmin) GetSigners() []types.AccAddress {
	addr, _ := types.AccAddressFromBech32(m.Admin)
	return []types.AccAddress{addr}
}
