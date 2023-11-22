package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/basket"
)

var _ legacytx.LegacyMsg = &MsgUpdateDateCriteria{}

// Route implements LegacyMsg.
func (m MsgUpdateDateCriteria) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements LegacyMsg.
func (m MsgUpdateDateCriteria) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements LegacyMsg.
func (m MsgUpdateDateCriteria) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgUpdateDateCriteria) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid authority address: %s", err)
	}

	if err := basket.ValidateBasketDenom(m.Denom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("invalid basket denom: %s", err)
	}

	if err := m.NewDateCriteria.Validate(); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("invalid date criteria: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdateDateCriteria.
func (m MsgUpdateDateCriteria) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
