package core

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/math"
)

// Validate checks if a BatchIssuance is valid.
func (i *BatchIssuance) Validate() error {
	if _, err := sdk.AccAddressFromBech32(i.Recipient); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("recipient: %s", err)
	}

	if i.TradableAmount == "" && i.RetiredAmount == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("tradable amount or retired amount required")
	}

	if i.TradableAmount != "" {
		if _, err := math.NewNonNegativeDecFromString(i.TradableAmount); err != nil {
			return fmt.Errorf("tradable amount: %s", err)
		}
	}

	if i.RetiredAmount != "" {
		retiredAmount, err := math.NewNonNegativeDecFromString(i.RetiredAmount)
		if err != nil {
			return fmt.Errorf("retired amount: %s", err)
		}

		if !retiredAmount.IsZero() {
			if i.RetirementJurisdiction == "" {
				return sdkerrors.ErrInvalidRequest.Wrap("retirement jurisdiction cannot be empty")
			}

			if err = ValidateJurisdiction(i.RetirementJurisdiction); err != nil {
				return fmt.Errorf("retirement jurisdiction: %s", err)
			}
		}
	}

	return nil
}
