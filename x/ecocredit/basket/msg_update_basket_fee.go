package basket

import (
	"github.com/regen-network/regen-ledger/x/ecocredit"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

var _ legacytx.LegacyMsg = &MsgUpdateBasketFeeRequest{}

func (m *MsgUpdateBasketFeeRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.RootAddress); err != nil {
		return sdkerrors.Wrap(err, sdkerrors.ErrInvalidAddress.Error())
	}
	if len(m.AddFees) == 0 && len(m.RemoveDenoms) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("at least one of add_fees or remove_denoms must be specified")
	}
	for _, fee := range m.AddFees {
		if err := sdk.ValidateDenom(fee.Denom); err != nil {
			return err
		}
		if _, ok := sdk.NewIntFromString(fee.Amount); !ok {
			return sdkerrors.ErrInvalidRequest.Wrapf("could not convert %s to %T", fee.Amount, sdk.Int{})
		}
	}
	for _, denom := range m.RemoveDenoms {
		if err := sdk.ValidateDenom(denom); err != nil {
			return err
		}
	}
	return nil
}

func (m *MsgUpdateBasketFeeRequest) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.RootAddress)
	return []sdk.AccAddress{addr}
}

func (m *MsgUpdateBasketFeeRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgUpdateBasketFeeRequest) Route() string {
	return sdk.MsgTypeURL(m)
}

func (m *MsgUpdateBasketFeeRequest) Type() string {
	return sdk.MsgTypeURL(m)
}
