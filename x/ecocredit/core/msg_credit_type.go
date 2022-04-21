package core

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	maxCreditTypeNameLength = 75
)

func (m CreditType) Validate() error {
	if err := ValidateCreditTypeAbbreviation(m.Abbreviation); err != nil {
		return err
	}
	if len(m.Name) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("name cannot be empty")
	}
	if len(m.Name) > maxCreditTypeNameLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("credit type name cannot exceed %d characters", maxCreditTypeNameLength)
	}
	if len(m.Unit) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("unit cannot be empty")
	}
	if m.Precision != PRECISION {
		return sdkerrors.ErrInvalidRequest.Wrapf("credit type precision is currently locked to %d", PRECISION)
	}
	return nil
}
