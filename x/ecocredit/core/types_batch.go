package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Validate performs a basic validation of credit batch
func (b Batch) Validate() error {
	if err := ValidateBatchDenom(b.Denom); err != nil {
		return err
	}

	if b.ProjectKey == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("project key cannot be zero")
	}

	if b.StartDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("must provide a start date for the credit batch")
	}
	if b.EndDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("must provide an end date for the credit batch")
	}
	if b.EndDate.Compare(*b.StartDate) != 1 {
		return sdkerrors.ErrInvalidRequest.Wrapf("the batch end date (%s) must be the same as or after the batch start date (%s)", b.EndDate.String(), b.StartDate.String())
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(b.Issuer).String()); err != nil {
		return sdkerrors.Wrap(err, "issuer")
	}

	return nil
}
