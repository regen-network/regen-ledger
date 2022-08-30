package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

var _ legacytx.LegacyMsg = &MsgTake{}

// Route implements LegacyMsg.
func (m MsgTake) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements LegacyMsg.
func (m MsgTake) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements LegacyMsg.
func (m MsgTake) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgTake) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf(err.Error())
	}

	if err := basket.ValidateBasketDenom(m.BasketDenom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("basket denom: %s", err)
	}

	if len(m.Amount) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("amount cannot be empty")
	}

	amount, ok := sdk.NewIntFromString(m.Amount)
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
			if err := core.ValidateJurisdiction(m.RetirementLocation); err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("retirement location: %s", err)
			}
		}

		if len(m.RetirementJurisdiction) != 0 {
			if err := core.ValidateJurisdiction(m.RetirementJurisdiction); err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("retirement jurisdiction: %s", err)
			}
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgTake.
func (m MsgTake) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}
