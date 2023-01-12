package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3/basket"
)

var _ legacytx.LegacyMsg = &MsgUpdateCurator{}

// Route implements LegacyMsg.
func (m MsgUpdateCurator) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements LegacyMsg.
func (m MsgUpdateCurator) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements LegacyMsg.
func (m MsgUpdateCurator) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgUpdateCurator) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Curator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("curator: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.NewCurator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("new curator: %s", err)
	}

	if m.NewCurator == m.Curator {
		return sdkerrors.ErrInvalidAddress.Wrap("curator and new curator cannot be the same")
	}

	if err := basket.ValidateBasketDenom(m.Denom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("basket denom: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdateCurator.
func (m MsgUpdateCurator) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Curator)
	return []sdk.AccAddress{addr}
}
