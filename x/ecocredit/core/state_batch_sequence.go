package core

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Validate performs basic validation of the BatchSequence state type
func (b BatchSequence) Validate() error {
	if b.ProjectKey == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("project key cannot be zero")
	}

	if b.NextSequence == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("next sequence cannot be zero")
	}

	return nil
}
