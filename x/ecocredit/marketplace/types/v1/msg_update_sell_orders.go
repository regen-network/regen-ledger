package v1

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/types/v2/math"
)

var _ legacytx.LegacyMsg = &MsgUpdateSellOrders{}

// Route implements the LegacyMsg interface.
func (m MsgUpdateSellOrders) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgUpdateSellOrders) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgUpdateSellOrders) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgUpdateSellOrders) ValidateBasic() error {
	if len(m.Seller) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("seller cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.Seller); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("seller is not a valid address: %s", err)
	}

	if len(m.Updates) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("updates cannot be empty")
	}

	for i, update := range m.Updates {
		updateIndex := fmt.Sprintf("updates[%d]", i)

		if update.SellOrderId == 0 {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: sell order id cannot be empty", updateIndex)
		}

		if len(update.NewQuantity) == 0 {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: new quantity cannot be empty", updateIndex)
		}

		if _, err := math.NewPositiveDecFromString(update.NewQuantity); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: new quantity must be a positive decimal", updateIndex)
		}

		// sdk.Coin.Validate panics if coin is nil
		if update.NewAskPrice == nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: new ask price cannot be empty", updateIndex)
		}

		// sdk.Coin.Validate provides inadequate error if coin denom is empty
		if len(update.NewAskPrice.Denom) == 0 {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: new ask price: denom cannot be empty", updateIndex)
		}

		if err := sdk.ValidateDenom(update.NewAskPrice.Denom); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: new ask price: %s", updateIndex, err)
		}

		// sdk.Coin.Validate panics if coin amount is nil
		if update.NewAskPrice.Amount.IsNil() {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: new ask price: amount cannot be empty", updateIndex)
		}

		// sdk.Coin.Validate provides inadequate error if coin amount is not a positive integer
		if !update.NewAskPrice.Amount.IsPositive() {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: new ask price: amount must be a positive integer", updateIndex)
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdateSellOrders.
func (m *MsgUpdateSellOrders) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Seller)
	return []sdk.AccAddress{addr}
}
