package core

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types/math"
)

// Validate checks if Credits is valid.
func (c *Credits) Validate() error {
	if err := ValidateBatchDenom(c.BatchDenom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	if c.Amount == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("amount cannot be empty")
	}

	if _, err := math.NewPositiveDecFromString(c.Amount); err != nil {
		return err
	}

	return nil
}
