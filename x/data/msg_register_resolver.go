package data

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ValidateBasic does a sanity check on the provided data.
func (m *MsgRegisterResolver) ValidateBasic() error {
	if m.ResolverId == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("resolver id cannot be empty")
	}

	if len(m.ContentHashes) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("content hashes cannot be empty")
	}

	for _, ch := range m.ContentHashes {
		if err := ch.Validate(); err != nil {
			return err
		}
	}

	return nil
}
