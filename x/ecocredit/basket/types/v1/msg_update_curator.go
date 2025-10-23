package v1

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/v4/basket"
)

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgUpdateCurator) ValidateBasic() error {
	if m.NewCurator == m.Curator {
		return sdkerrors.ErrInvalidAddress.Wrap("curator and new curator cannot be the same")
	}

	if err := basket.ValidateBasketDenom(m.Denom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("basket denom: %s", err)
	}

	return nil
}
