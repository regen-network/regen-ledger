package v1

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/v2/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/base"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/basket"
)

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgPut) ValidateBasic() error {
	if err := basket.ValidateBasketDenom(m.BasketDenom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("basket denom: %s", err)
	}

	if len(m.Credits) > 0 {
		for i, credit := range m.Credits {
			creditIndex := fmt.Sprintf("credits[%d]", i)

			if err := base.ValidateBatchDenom(credit.BatchDenom); err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("%s: batch denom: %s", creditIndex, err)
			}

			if len(credit.Amount) == 0 {
				return sdkerrors.ErrInvalidRequest.Wrap("credit amount cannot be empty")
			}

			if _, err := math.NewPositiveDecFromString(credit.Amount); err != nil {
				return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
			}
		}
	} else {
		return sdkerrors.ErrInvalidRequest.Wrap("credits cannot be empty")
	}

	return nil
}
