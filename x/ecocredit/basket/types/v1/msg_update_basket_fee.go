package v1

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgUpdateBasketFee) ValidateBasic() error {
	if m.Fee != nil {
		if err := m.Fee.Validate(); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s", err)
		}
	}

	return nil
}
