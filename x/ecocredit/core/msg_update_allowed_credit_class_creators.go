package core

import (
	"github.com/regen-network/regen-ledger/x/ecocredit"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

var _ legacytx.LegacyMsg = &MsgUpdateAllowedCreditClassCreators{}

func (m MsgUpdateAllowedCreditClassCreators) ValidateBasic() error {
	validateAddresses := func(addrs ...string) error {
		for _, addr := range addrs {
			if _, err := sdk.AccAddressFromBech32(addr); err != nil {
				return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
			}
		}
		return nil
	}
	if len(m.AddCreators) == 0 && len(m.RemoveCreators) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("must specify creators to add and/or remove")
	}
	if err := validateAddresses(m.RootAddress); err != nil {
		return err
	}
	if err := validateAddresses(m.AddCreators...); err != nil {
		return err
	}
	return validateAddresses(m.RemoveCreators...)
}

func (m MsgUpdateAllowedCreditClassCreators) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.RootAddress)
	return []sdk.AccAddress{addr}
}

func (m MsgUpdateAllowedCreditClassCreators) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))

}

func (m MsgUpdateAllowedCreditClassCreators) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateAllowedCreditClassCreators) Type() string { return sdk.MsgTypeURL(&m) }
