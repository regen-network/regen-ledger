package v1

import (
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateProjectFee{}

// Route implements the LegacyMsg interface.
func (m MsgUpdateProjectFee) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgUpdateProjectFee) Type() string { return sdk.MsgTypeURL(&m) }

// ValidateBasic does a sanity check on the provided data.
func (m *MsgUpdateProjectFee) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(err, "invalid authority address")
	}

	if m.Fee != nil {
		if err := m.Fee.Validate(); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s", err)
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdateProjectFee.
func (m *MsgUpdateProjectFee) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
