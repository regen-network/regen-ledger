package basket

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgPut{}

// Route implements LegacyMsg.
func (m MsgPut) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements LegacyMsg.
func (m MsgPut) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements LegacyMsg.
func (m MsgPut) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgPut) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf(err.Error())
	}
	if err := sdk.ValidateDenom(m.BasketDenom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid basket denom", m.BasketDenom)
	}
	if len(m.Credits) > 0 {
		for _, credit := range m.Credits {
			if err := ecocredit.ValidateDenom(credit.BatchDenom); err != nil {
				return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
			}
			if _, err := math.NewPositiveDecFromString(credit.Amount); err != nil {
				return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
			}
		}
	} else {
		return sdkerrors.ErrInvalidRequest.Wrap("no credits were specified to put into the basket")
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreate.
func (m MsgPut) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}
