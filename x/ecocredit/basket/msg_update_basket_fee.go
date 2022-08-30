package basket

import (
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgUpdateBasketFees{}

// Route implements LegacyMsg.
func (m MsgUpdateBasketFees) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements LegacyMsg.
func (m MsgUpdateBasketFees) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements LegacyMsg.
func (m MsgUpdateBasketFees) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgUpdateBasketFees) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(err, "invalid authority address")
	}

	return m.BasketFees.Validate()
}

// GetSigners returns the expected signers for MsgUpdateBasketFee.
func (m MsgUpdateBasketFees) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
