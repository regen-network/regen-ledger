package v1

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/base"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/basket"
)

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgTake) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	if err := basket.ValidateBasketDenom(m.BasketDenom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("basket denom: %s", err)
	}

	if len(m.Amount) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("amount cannot be empty")
	}

	amount, ok := math.NewIntFromString(m.Amount)
	if !ok {
		return sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid integer", m.Amount)
	}

	if !amount.IsPositive() {
		return sdkerrors.ErrInvalidRequest.Wrapf("amount must be positive, got %s", m.Amount)
	}

	if m.RetireOnTake {
		// retirement_location is deprecated but still supported
		if len(m.RetirementLocation) == 0 && len(m.RetirementJurisdiction) == 0 {
			return sdkerrors.ErrInvalidRequest.Wrap("retirement jurisdiction cannot be empty if retire on take is true")
		}

		// retirement_location is deprecated but still supported
		if len(m.RetirementLocation) != 0 {
			if err := base.ValidateJurisdiction(m.RetirementLocation); err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("retirement location: %s", err)
			}
		}

		if len(m.RetirementJurisdiction) != 0 {
			if err := base.ValidateJurisdiction(m.RetirementJurisdiction); err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("retirement jurisdiction: %s", err)
			}
		}

		if len(m.RetirementReason) > base.MaxNoteLength {
			return ecocredit.ErrMaxLimit.Wrapf("retirement reason: max length %d", base.MaxNoteLength)
		}
	}

	return nil
}
