package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgUpdateClassIssuers{}

func (m MsgUpdateClassIssuers) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassIssuers) Type() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassIssuers) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

func (m *MsgUpdateClassIssuers) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	if err := ValidateClassID(m.ClassId); err != nil {
		return err
	}

	if len(m.AddIssuers) == 0 && len(m.RemoveIssuers) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("must specify at least one of add_issuers or remove_issuers")
	}

	validateIssuers := func(issuers []string) error {
		for _, addr := range issuers {
			if _, err := sdk.AccAddressFromBech32(addr); err != nil {
				return sdkerrors.ErrInvalidAddress.Wrap(addr)
			}
		}
		return nil
	}

	if err := validateIssuers(m.AddIssuers); err != nil {
		return err
	}

	return validateIssuers(m.RemoveIssuers)
}

func (m *MsgUpdateClassIssuers) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}
