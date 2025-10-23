package v1

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/basket"
)

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgUpdateDateCriteria) ValidateBasic() error {
	if err := basket.ValidateBasketDenom(m.Denom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("invalid basket denom: %s", err)
	}

	if err := m.NewDateCriteria.Validate(); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("invalid date criteria: %s", err)
	}

	return nil
}
