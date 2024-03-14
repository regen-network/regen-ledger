package v1

import (
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgGovSetFeeParams{}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgGovSetFeeParams) ValidateBasic() error {
	if m.Fees == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("fees cannot be nil")
	}

	err := m.Fees.Validate()
	if err != nil {
		return err
	}

	_, err = types.AccAddressFromBech32(m.Authority)
	return err
}

// GetSigners implements the LegacyMsg interface.
func (m *MsgGovSetFeeParams) GetSigners() []types.AccAddress {
	return []types.AccAddress{types.MustAccAddressFromBech32(m.Authority)}
}

// Route implements the LegacyMsg interface.
func (m *MsgGovSetFeeParams) Route() string { return types.MsgTypeURL(m) }

// Type implements the LegacyMsg interface.
func (m *MsgGovSetFeeParams) Type() string { return types.MsgTypeURL(m) }

// GetSignBytes implements the LegacyMsg interface.
func (m *MsgGovSetFeeParams) GetSignBytes() []byte {
	return types.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}
