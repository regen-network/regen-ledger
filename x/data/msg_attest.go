package data

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ValidateBasic does a sanity check on the provided data.
func (m *MsgAttest) ValidateBasic() error {
	if len(m.ContentHashes) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("content hashes cannot be empty")
	}

	for _, hash := range m.ContentHashes {
		if hash == nil {
			return sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
		}
		err := hash.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}
