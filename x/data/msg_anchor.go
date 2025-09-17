package data

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ValidateBasic does a sanity check on the provided data.
func (m *MsgAnchor) ValidateBasic() error {
	if m.ContentHash == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
	}

	return m.ContentHash.Validate()
}
