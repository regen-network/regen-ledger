package marketplace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

var _ legacytx.LegacyMsg = &MsgSell{}

// Route implements the LegacyMsg interface.
func (m MsgSell) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgSell) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgSell) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgSell) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	for _, order := range m.Orders {
		if err := core.ValidateBatchDenom(order.BatchDenom); err != nil {
			return err
		}

		if _, err := math.NewPositiveDecFromString(order.Quantity); err != nil {
			return sdkerrors.Wrapf(err, "quantity must be positive decimal: %s", order.Quantity)
		}

		if order.AskPrice == nil {
			return sdkerrors.ErrInvalidRequest.Wrap("ask price cannot be empty")
		}

		if err := order.AskPrice.Validate(); err != nil {
			return err
		}

		if !order.AskPrice.Amount.IsPositive() {
			return sdkerrors.ErrInvalidRequest.Wrap("ask price must be positive amount")
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgSell.
func (m *MsgSell) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}
