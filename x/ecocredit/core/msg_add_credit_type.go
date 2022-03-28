package core

import (
	"github.com/regen-network/regen-ledger/x/ecocredit"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

var _ legacytx.LegacyMsg = &MsgAddCreditType{}

func (m MsgAddCreditType) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.RootAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}
	if len(m.CreditTypes) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("credit types cannot be empty")
	}
	for _, ct := range m.CreditTypes {
		if err := ecocredit.ValidateCreditTypeAbbreviation(ct.Abbreviation); err != nil {
			return err
		}
		if len(ct.Name) == 0 {
			return sdkerrors.ErrInvalidRequest.Wrap("name is required")
		}
		if len(ct.Unit) == 0 {
			return sdkerrors.ErrInvalidRequest.Wrap("unit of measurement is required")
		}
		if ct.Precision == 0 {
			return sdkerrors.ErrInvalidRequest.Wrap("precision must be greater than 0")
		}
	}
	return nil
}

func (m MsgAddCreditType) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.RootAddress)
	return []sdk.AccAddress{addr}
}

func (m MsgAddCreditType) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgAddCreditType) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgAddCreditType) Type() string { return sdk.MsgTypeURL(&m) }
