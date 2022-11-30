package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/v2/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
)

// Validate performs basic validation of the SellOrder state type
func (m *SellOrder) Validate() error {
	if m.Id == 0 {
		return ecocredit.ErrParseFailure.Wrapf("id cannot be zero")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Seller).String()); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("seller: %s", err)
	}

	if m.BatchKey == 0 {
		return ecocredit.ErrParseFailure.Wrapf("batch key cannot be zero")
	}

	if m.Quantity == "" {
		return ecocredit.ErrParseFailure.Wrapf("quantity cannot be empty")
	}

	if _, err := math.NewNonNegativeDecFromString(m.Quantity); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("quantity: %s", err)
	}

	if m.MarketId == 0 {
		return ecocredit.ErrParseFailure.Wrapf("market id cannot be zero")
	}

	if m.AskAmount == "" {
		return ecocredit.ErrParseFailure.Wrapf("ask amount cannot be empty")
	}

	if _, err := math.NewNonNegativeDecFromString(m.AskAmount); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("ask amount: %s", err)
	}

	return nil
}
