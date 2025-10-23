package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddCreditType{}

// Route implements the LegacyMsg interface.
func (m MsgAddCreditType) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgAddCreditType) Type() string { return sdk.MsgTypeURL(&m) }

// ValidateBasic does a sanity check on the provided data.
func (m *MsgAddCreditType) ValidateBasic() error {
	if m.CreditType == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("credit type cannot be empty")
	}

	if err := m.CreditType.Validate(); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("credit type: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgAddCreditType.
func (m *MsgAddCreditType) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
