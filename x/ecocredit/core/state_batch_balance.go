package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types/math"
)

// Validate performs basic validation of the BatchBalance state type
func (b BatchBalance) Validate() error {
	if b.BatchKey == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("batch key cannot be zero")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(b.Address).String()); err != nil {
		return sdkerrors.Wrap(err, "address")
	}

	if _, err := math.NewNonNegativeDecFromString(b.TradableAmount); err != nil {
		return sdkerrors.Wrapf(err, "tradable amount")
	}

	if _, err := math.NewNonNegativeDecFromString(b.RetiredAmount); err != nil {
		return sdkerrors.Wrapf(err, "retired amount")
	}

	if _, err := math.NewNonNegativeDecFromString(b.EscrowedAmount); err != nil {
		return sdkerrors.Wrapf(err, "escrowed amount")
	}

	return nil
}
