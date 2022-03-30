package marketplace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgUpdateSellOrders{}

// Route implements the LegacyMsg interface.
func (m MsgUpdateSellOrders) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgUpdateSellOrders) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgUpdateSellOrders) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgUpdateSellOrders) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	for _, update := range m.Updates {

		if _, err := math.NewPositiveDecFromString(update.NewQuantity); err != nil {
			return sdkerrors.Wrapf(err, "quantity must be positive decimal: %s", update.NewQuantity)
		}

		if update.NewAskPrice == nil {
			return sdkerrors.ErrInvalidRequest.Wrap("new ask price cannot be empty")
		}

		if err := update.NewAskPrice.Validate(); err != nil {
			return err
		}

		if !update.NewAskPrice.Amount.IsPositive() {
			return sdkerrors.ErrInvalidRequest.Wrap("ask price must be positive amount")
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdateSellOrders.
func (m *MsgUpdateSellOrders) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}
