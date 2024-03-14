package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgGovSendFromFeePool{}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgGovSendFromFeePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		return err
	}

	_, err = sdk.AccAddressFromBech32(m.Recipient)
	if err != nil {
		return err
	}

	if m.Coins == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("coins cannot be nil")
	}

	return m.Coins.Validate()
}

// GetSigners implements the LegacyMsg interface.
func (m *MsgGovSendFromFeePool) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Authority)}
}

// Route implements the LegacyMsg interface.
func (m *MsgGovSendFromFeePool) Route() string { return sdk.MsgTypeURL(m) }

// Type implements the LegacyMsg interface.
func (m *MsgGovSendFromFeePool) Type() string { return sdk.MsgTypeURL(m) }

// GetSignBytes implements the LegacyMsg interface.
func (m *MsgGovSendFromFeePool) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}
