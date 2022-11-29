package v1

import (
	"cosmossdk.io/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
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
			return errors.Wrap(err, "tradable amount")
		}
	}

	if i.RetiredAmount != "" {
		retiredAmount, err := math.NewNonNegativeDecFromString(i.RetiredAmount)
		if err != nil {
			return errors.Wrap(err, "retired amount")
		}

		if !retiredAmount.IsZero() {
			if err = base.ValidateJurisdiction(i.RetirementJurisdiction); err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("retirement jurisdiction: %s", err)
			}

			if len(i.RetirementReason) > base.MaxNoteLength {
				return ecocredit.ErrMaxLimit.Wrapf("retirement reason: max length %d", base.MaxNoteLength)
			}
		}
	}

	return nil
}
