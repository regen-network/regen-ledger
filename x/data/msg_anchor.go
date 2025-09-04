package data

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ValidateBasic does a sanity check on the provided data.
func (m *MsgAnchor) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}

	if m.ContentHash == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
	}

	return m.ContentHash.Validate()
}
