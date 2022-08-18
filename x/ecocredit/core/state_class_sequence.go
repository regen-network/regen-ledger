package core

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Validate performs basic validation of the ClassSequence state type
func (b ClassSequence) Validate() error {
	if err := ValidateCreditTypeAbbreviation(b.CreditTypeAbbrev); err != nil {
		return err
	}

	if b.NextSequence == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("next sequence cannot be zero")
	}

	return nil
}
