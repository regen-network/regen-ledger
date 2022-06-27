package basket

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
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

	if len(m.BasketDenom) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("basket denom cannot be empty")
	}

	if err := ValidateBasketDenom(m.BasketDenom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	if len(m.Credits) > 0 {
		for _, credit := range m.Credits {
			if len(credit.BatchDenom) == 0 {
				return sdkerrors.ErrInvalidRequest.Wrap("credit batch denom cannot be empty")
			}

			if err := core.ValidateBatchDenom(credit.BatchDenom); err != nil {
				return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
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

// GetSigners returns the expected signers for MsgCreate.
func (m MsgPut) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}
