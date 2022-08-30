package marketplace

import (
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgRemoveAllowedDenom{}

// Route implements the LegacyMsg interface.
func (m MsgRemoveAllowedDenom) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgRemoveAllowedDenom) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgRemoveAllowedDenom) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m MsgRemoveAllowedDenom) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(err, "invalid authority address")
	}

	if m.Denom == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("denom: %s", err.Error())
	}

	return nil
}

// GetSigners returns the expected signers for MsgCancelSellOrder.
func (m MsgRemoveAllowedDenom) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
