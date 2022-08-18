package core

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types/math"
)

// Validate performs basic validation of the BatchSupply state type
func (b BatchSupply) Validate() error {
	if b.BatchKey == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("batch key cannot be zero")
	}

	if _, err := math.NewDecFromString(b.TradableAmount); err != nil {
		return sdkerrors.Wrapf(err, "tradable amount")
	}

	if _, err := math.NewDecFromString(b.RetiredAmount); err != nil {
		return sdkerrors.Wrapf(err, "retired amount")
	}

	if _, err := math.NewDecFromString(b.CancelledAmount); err != nil {
		return sdkerrors.Wrapf(err, "cancelled amount")
	}

	return nil
}
