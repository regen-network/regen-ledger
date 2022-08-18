package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs basic validation of the BatchBalance state type
func (b BatchBalance) Validate() error {
	if b.BatchKey == 0 {
		return ecocredit.ErrParseFailure.Wrapf("batch key cannot be zero")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(b.Address).String()); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("address: %s", err)
	}

	if _, err := math.NewNonNegativeDecFromString(b.TradableAmount); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("tradable amount: %s", err)
	}

	if _, err := math.NewNonNegativeDecFromString(b.RetiredAmount); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("retired amount: %s", err)
	}

	if _, err := math.NewNonNegativeDecFromString(b.EscrowedAmount); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("escrowed amount: %s", err)
	}

	return nil
}
