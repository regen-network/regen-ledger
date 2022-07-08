package core

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types/math"
)

// Validate checks if Credits is valid.
func (c *Credits) Validate() error {
	if c.BatchDenom == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("batch denom cannot be empty")
	}

	if err := ValidateBatchDenom(c.BatchDenom); err != nil {
		return err
	}

	if c.Amount == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("amount cannot be empty")
	}

	if _, err := math.NewPositiveDecFromString(c.Amount); err != nil {
		return err
	}

	return nil
}
