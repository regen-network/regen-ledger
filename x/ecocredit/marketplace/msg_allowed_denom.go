package marketplace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (m AllowedDenom) Validate() error {
	if err := sdk.ValidateDenom(m.BankDenom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("invalid denom: %s", err.Error())
	}
	if err := sdk.ValidateDenom(m.DisplayDenom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("invalid display_denom: %s", err.Error())
	}
	if m.Exponent > 18 {
		return sdkerrors.ErrInvalidRequest.Wrap("exponent cannot be more than 18")
	}
	return nil
}
