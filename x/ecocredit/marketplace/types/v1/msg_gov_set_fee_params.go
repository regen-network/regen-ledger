package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgGovSetFeeParams{}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgGovSetFeeParams) ValidateBasic() error {
	if m.Fees == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("fees cannot be nil")
	}

	err := m.Fees.Validate()
	if err != nil {
		return err
	}

	_, err = sdk.AccAddressFromBech32(m.Authority)
	return err
}

// Route implements the LegacyMsg interface.
func (m *MsgGovSetFeeParams) Route() string { return sdk.MsgTypeURL(m) }

// Type implements the LegacyMsg interface.
func (m *MsgGovSetFeeParams) Type() string { return sdk.MsgTypeURL(m) }
