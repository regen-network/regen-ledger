package basket

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
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
	if err := sdk.ValidateDenom(m.BasketDenom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid basket denom", m.BasketDenom)
	}
	i, ok := sdk.NewIntFromString(m.Amount)
	if !ok {
		return sdkerrors.ErrInvalidRequest.Wrapf("%s can't be converted to an integer", m.Amount)
	}
	if !i.IsPositive() {
		return sdkerrors.ErrInvalidRequest.Wrapf("%s must be positive", m.Amount)
	}
	if m.RetireOnTake {
		if err := ecocredit.ValidateJurisdiction(m.RetirementJurisdiction); err != nil {
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
