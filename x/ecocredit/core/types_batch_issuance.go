package core

import (
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
			return sdkerrors.Wrap(err, "tradable amount")
		}
	}

	if i.RetiredAmount != "" {
		retiredAmount, err := math.NewNonNegativeDecFromString(i.RetiredAmount)
		if err != nil {
			return sdkerrors.Wrap(err, "retired amount")
		}

		if !retiredAmount.IsZero() {
			if err = ValidateJurisdiction(i.RetirementJurisdiction); err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("retirement jurisdiction: %s", err.Error())
			}
		}
	}

	return nil
}
