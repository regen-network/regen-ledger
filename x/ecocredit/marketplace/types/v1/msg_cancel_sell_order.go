package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCancelSellOrder{}

// Route implements the LegacyMsg interface.
func (m MsgCancelSellOrder) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCancelSellOrder) Type() string { return sdk.MsgTypeURL(&m) }

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCancelSellOrder) ValidateBasic() error {
	if len(m.Seller) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrapf("seller cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.Seller); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("seller is not a valid address: %s", err)
	}

	if m.SellOrderId == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("sell order id cannot be empty")
	}

	return nil
}
