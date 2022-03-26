package marketplace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgCancelSellOrder{}

// Route implements the LegacyMsg interface.
func (m MsgCancelSellOrder) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCancelSellOrder) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCancelSellOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCancelSellOrder) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Seller); err != nil {
		return sdkerrors.ErrInvalidAddress
	}
	if m.SellOrderId == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("0 is not a valid sell order id")
	}
	return nil
}

// GetSigners returns the expected signers for MsgCancelSellOrder.
func (m *MsgCancelSellOrder) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Seller)
	return []sdk.AccAddress{addr}
}
