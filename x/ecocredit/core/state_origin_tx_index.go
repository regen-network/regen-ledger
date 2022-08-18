package core

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Validate performs basic validation of the OriginTxIndex state type
func (b OriginTxIndex) Validate() error {
	if b.ClassKey == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("class key cannot be zero")
	}

	if b.Id == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("origin_tx.id cannot be empty")
	}

	if !reOriginTxId.MatchString(b.Id) {
		return sdkerrors.ErrInvalidRequest.Wrap("origin_tx.id must be at most 128 characters long, valid characters: alpha-numberic, space, '-' or '_'")
	}

	if b.Source == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("origin_tx.source cannot be empty")
	}

	if !reOriginTxSource.MatchString(b.Source) {
		return sdkerrors.ErrInvalidRequest.Wrap("origin_tx.source must be at most 32 characters long, valid characters: alpha-numberic, space, '-' or '_'")
	}

	return nil
}
