package core

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types/eth"
)

// Validate performs basic validation of the BatchContract state type
func (b BatchContract) Validate() error {
	if b.BatchKey == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("batch key cannot be zero")
	}

	if b.ClassKey == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("class key cannot be zero")
	}

	if !eth.IsValidAddress(b.Contract) {
		return sdkerrors.ErrInvalidAddress.Wrap("contract must be a valid ethereum address")
	}

	return nil
}
