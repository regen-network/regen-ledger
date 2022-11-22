package v1

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	"github.com/regen-network/regen-ledger/x/ecocredit"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
)

var _ legacytx.LegacyMsg = &MsgBuyDirect{}

func (m MsgBuyDirect) ValidateBasic() error {
	if len(m.Buyer) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("buyer cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.Buyer); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("buyer is not a valid address: %s", err)
	}

	if len(m.Orders) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("orders cannot be empty")
	}

	for i, order := range m.Orders {
		// orderIndex is used for more granular error messages when
		// an individual order in a list of orders fails to process
		orderIndex := fmt.Sprintf("orders[%d]", i)

		if order.SellOrderId == 0 {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: sell order id cannot be empty", orderIndex)
		}

		if len(order.Quantity) == 0 {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: quantity cannot be empty", orderIndex)
		}

		if _, err := math.NewPositiveDecFromString(order.Quantity); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: quantity must be a positive decimal", orderIndex)
		}

		// sdk.Coin.Validate panics if coin is nil
		if order.BidPrice == nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: bid price cannot be empty", orderIndex)
		}

		// sdk.Coin.Validate provides inadequate error if coin denom is empty
		if len(order.BidPrice.Denom) == 0 {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: bid price: denom cannot be empty", orderIndex)
		}

		if err := sdk.ValidateDenom(order.BidPrice.Denom); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: bid price: %s", orderIndex, err)
		}

		// sdk.Coin.Validate panics if coin amount is nil
		if order.BidPrice.Amount.IsNil() {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: bid price: amount cannot be empty", orderIndex)
		}

		// sdk.Coin.Validate provides inadequate error if coin amount is not a positive integer
		if !order.BidPrice.Amount.IsPositive() {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: bid price: amount must be a positive integer", orderIndex)
		}

		if !order.DisableAutoRetire {
			if err := base.ValidateJurisdiction(order.RetirementJurisdiction); err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("%s: retirement jurisdiction: %s", orderIndex, err)
			}

			if len(order.RetirementReason) > base.MaxNoteLength {
				return ecocredit.ErrMaxLimit.Wrapf("%s: retirement reason: max length %d", orderIndex, base.MaxNoteLength)
			}
		}
	}

	return nil
}

func (m MsgBuyDirect) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Buyer)
	return []sdk.AccAddress{addr}
}

func (m MsgBuyDirect) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgBuyDirect) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgBuyDirect) Type() string { return sdk.MsgTypeURL(&m) }
