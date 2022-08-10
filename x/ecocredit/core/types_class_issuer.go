package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Validate performs a basic validation of credit class issuers
func (c ClassIssuer) Validate() error {
	if c.ClassKey == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("class key cannot be zero")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(c.Issuer).String()); err != nil {
		return sdkerrors.Wrap(err, "issuer")
	}

	return nil
}
