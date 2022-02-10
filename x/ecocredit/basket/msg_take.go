package basket

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var (
	_ legacytx.LegacyMsg = &MsgTake{}
)

// Route Implements LegacyMsg.
func (m MsgTake) Route() string { return sdk.MsgTypeURL(&m) }

// Type Implements LegacyMsg.
func (m MsgTake) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes Implements LegacyMsg.
func (m MsgTake) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgTake) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf(err.Error())
	}
	if err := sdk.ValidateDenom(m.BasketDenom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid basket denom", m.BasketDenom)
	}
	if _, err := math.NewPositiveDecFromString(m.Amount); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}
	if m.RetireOnTake {
		if err := ecocredit.ValidateLocation(m.RetirementLocation); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
		}
	}
	return nil
}

// GetSigners returns the expected signers for MsgTake.
func (m MsgTake) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}
