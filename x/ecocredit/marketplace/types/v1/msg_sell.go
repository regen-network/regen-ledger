package v1

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
)

var _ legacytx.LegacyMsg = &MsgSell{}

// Route implements the LegacyMsg interface.
func (m MsgSell) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgSell) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgSell) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgSell) ValidateBasic() error {
	if len(m.Seller) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("seller cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.Seller); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("seller is not a valid address: %s", err)
	}

	if len(m.Orders) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("orders cannot be empty")
	}

	for i, order := range m.Orders {
		// orderIndex is used for more granular error messages when
		// an individual order in a list of orders fails to process
		orderIndex := fmt.Sprintf("orders[%d]", i)

		if err := base.ValidateBatchDenom(order.BatchDenom); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: batch denom: %s", orderIndex, err)
		}

		if len(order.Quantity) == 0 {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: quantity cannot be empty", orderIndex)
		}

		if _, err := math.NewPositiveDecFromString(order.Quantity); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: quantity must be a positive decimal", orderIndex)
		}

		// sdk.Coin.Validate panics if coin is nil
		if order.AskPrice == nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: ask price: cannot be empty", orderIndex)
		}

		// sdk.Coin.Validate provides inadequate error if coin denom is empty
		if len(order.AskPrice.Denom) == 0 {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: ask price: denom cannot be empty", orderIndex)
		}

		if err := sdk.ValidateDenom(order.AskPrice.Denom); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: ask price: %s", orderIndex, err)
		}

		// sdk.Coin.Validate panics if coin amount is nil
		if order.AskPrice.Amount.IsNil() {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: ask price: amount cannot be empty", orderIndex)
		}

		// sdk.Coin.Validate provides inadequate error if coin amount is not a positive integer
		if !order.AskPrice.Amount.IsPositive() {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: ask price: amount must be a positive integer", orderIndex)
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgSell.
func (m *MsgSell) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Seller)
	return []sdk.AccAddress{addr}
}
