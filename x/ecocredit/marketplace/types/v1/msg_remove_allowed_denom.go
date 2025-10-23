package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRemoveAllowedDenom{}

// Route implements the LegacyMsg interface.
func (m MsgRemoveAllowedDenom) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgRemoveAllowedDenom) Type() string { return sdk.MsgTypeURL(&m) }

// ValidateBasic does a sanity check on the provided data.
func (m MsgRemoveAllowedDenom) ValidateBasic() error {
	if m.Denom == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("denom: %s", err.Error())
	}

	return nil
}
