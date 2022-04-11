package core

import (
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgNewCreditType{}

func (m MsgNewCreditType) ValidateBasic() error {
	if _, err := types.AccAddressFromBech32(m.RootAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}
	return m.CreditType.Validate()
}

func (m MsgNewCreditType) GetSigners() []types.AccAddress {
	addr, _ := types.AccAddressFromBech32(m.RootAddress)
	return []types.AccAddress{addr}
}

func (m MsgNewCreditType) GetSignBytes() []byte {
	return types.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgNewCreditType) Route() string { return types.MsgTypeURL(&m) }

func (m MsgNewCreditType) Type() string { return types.MsgTypeURL(&m) }

func (m CreditType) Validate() error {
	if err := ValidateCreditTypeAbbreviation(m.Abbreviation); err != nil {
		return err
	}
	if len(m.Name) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("name cannot be empty")
	}
	if len(m.Unit) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("unit cannot be empty")
	}
	if m.Precision != PRECISION {
		return sdkerrors.ErrInvalidRequest.Wrapf("credit type precision is currently locked to %d", PRECISION)
	}
	return nil
}
