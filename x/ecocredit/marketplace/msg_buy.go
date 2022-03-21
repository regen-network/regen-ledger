package marketplace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

var _ legacytx.LegacyMsg = &MsgBuy{}

// Route implements the LegacyMsg interface.
func (m MsgBuy) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgBuy) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgBuy) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgBuy) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Buyer); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	for _, order := range m.Orders {

		if _, err := math.NewPositiveDecFromString(order.Quantity); err != nil {
			return sdkerrors.Wrapf(err, "quantity must be positive decimal: %s", order.Quantity)
		}

		if order.BidPrice == nil {
			return sdkerrors.ErrInvalidRequest.Wrap("bid price cannot be empty")
		}

		if err := order.BidPrice.Validate(); err != nil {
			return err
		}

		if !order.BidPrice.Amount.IsPositive() {
			return sdkerrors.ErrInvalidRequest.Wrap("bid price must be positive amount")
		}

		if !order.DisableAutoRetire {
			if err := core.ValidateLocation(order.RetirementLocation); err != nil {
				// ValidateLocation returns an sdkerrors.ErrInvalidRequest, so we can just wrap it here
				return sdkerrors.Wrap(err, "a valid retirement location is required when DisableAutoRetire is false")
			}
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgBuy.
func (m *MsgBuy) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Buyer)
	return []sdk.AccAddress{addr}
}
