package v1

import (
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddAllowedDenom{}

// Route implements the LegacyMsg interface.
func (m MsgAddAllowedDenom) Route() string { return sdk.MsgTypeURL(&m) }

// Route implements the LegacyMsg interface.
func (m MsgAddAllowedDenom) Type() string { return sdk.MsgTypeURL(&m) }

// ValidateBasic does a sanity check on the provided data.
func (m MsgAddAllowedDenom) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(err, "invalid authority address")
	}

	allowedDenom := AllowedDenom{
		BankDenom:    m.BankDenom,
		DisplayDenom: m.DisplayDenom,
		Exponent:     m.Exponent,
	}

	if err := allowedDenom.Validate(); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	return nil
}

// GetSigners returns the expected signers for MsgAddAllowedDenom.
func (m MsgAddAllowedDenom) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
