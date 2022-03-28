package basket

import (
	"github.com/regen-network/regen-ledger/x/ecocredit"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

var _ legacytx.LegacyMsg = &MsgUpdateBasketFee{}

func (m MsgUpdateBasketFee) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.RootAddress); err != nil {
		return sdkerrors.Wrap(err, sdkerrors.ErrInvalidAddress.Error())
	}
	if len(m.AddFees) == 0 && len(m.RemoveFees) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("at least one of add_fees or remove_denoms must be specified")
	}
	for _, fee := range m.AddFees {
		if err := fee.Validate(); err != nil {
			return err
		}
		if fee.Amount.IsZero() {
			return sdkerrors.ErrInvalidRequest.Wrap("fee must be greater than zero")
		}
	}
	for _, denom := range m.RemoveFees {
		if err := sdk.ValidateDenom(denom); err != nil {
			return err
		}
	}
	return nil
}

func (m MsgUpdateBasketFee) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.RootAddress)
	return []sdk.AccAddress{addr}
}

func (m MsgUpdateBasketFee) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgUpdateBasketFee) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateBasketFee) Type() string { return sdk.MsgTypeURL(&m) }
