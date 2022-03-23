package core

import (
	"github.com/regen-network/regen-ledger/x/ecocredit"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

var _ legacytx.LegacyMsg = &MsgUpdateCreditClassFeeRequest{}

func (m *MsgUpdateCreditClassFeeRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.RootAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}
	if len(m.AddFees) == 0 && len(m.RemoveDenoms) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("at least one of add_fees or remove_denoms must be specified")
	}
	for _, fee := range m.AddFees {
		if err := sdk.ValidateDenom(fee.Denom); err != nil {
			return err
		}
		if fee.IsZero() {
			return sdkerrors.ErrInsufficientFee.Wrap("fee must be non-zero")
		}
	}
	for _, denom := range m.RemoveDenoms {
		if err := sdk.ValidateDenom(denom); err != nil {
			return err
		}
	}
	return nil
}

func (m *MsgUpdateCreditClassFeeRequest) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.RootAddress)
	return []sdk.AccAddress{addr}
}

func (m *MsgUpdateCreditClassFeeRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgUpdateCreditClassFeeRequest) Route() string {
	return sdk.MsgTypeURL(m)
}

func (m *MsgUpdateCreditClassFeeRequest) Type() string {
	return sdk.MsgTypeURL(m)
}
