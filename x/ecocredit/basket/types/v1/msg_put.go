package v1

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit/basket"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
)

var _ legacytx.LegacyMsg = &MsgPut{}

// Route implements LegacyMsg.
func (m MsgPut) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements LegacyMsg.
func (m MsgPut) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements LegacyMsg.
func (m MsgPut) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgPut) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf(err.Error())
	}

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

// GetSigners returns the expected signers for MsgCreate.
func (m MsgPut) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}
