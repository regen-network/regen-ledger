package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgGovSendFromFeePool{}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgGovSendFromFeePool) ValidateBasic() error {
	if m.Coins == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("coins cannot be nil")
	}

	return m.Coins.Validate()
}

// Route implements the LegacyMsg interface.
func (m *MsgGovSendFromFeePool) Route() string { return sdk.MsgTypeURL(m) }

// Type implements the LegacyMsg interface.
func (m *MsgGovSendFromFeePool) Type() string { return sdk.MsgTypeURL(m) }
