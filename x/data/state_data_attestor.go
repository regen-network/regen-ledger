package data

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate performs basic validation of the DataAttestor state type
func (m *DataAttestor) Validate() error {
	if len(m.Id) == 0 {
		return ErrParseFailure.Wrap("id cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Attestor).String()); err != nil {
		return ErrParseFailure.Wrapf("attestor: %s", err)
	}

	if m.Timestamp == nil {
		return ErrParseFailure.Wrapf("timestamp cannot be empty")
	}

	return nil
}
