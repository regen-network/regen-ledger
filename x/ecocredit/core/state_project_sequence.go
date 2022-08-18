package core

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Validate performs basic validation of the ProjectSequence state type
func (b ProjectSequence) Validate() error {
	if b.ClassKey == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("class key cannot be zero")
	}

	if b.NextSequence == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("next sequence cannot be zero")
	}

	return nil
}
