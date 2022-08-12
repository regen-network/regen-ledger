package marketplace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/x/ecocredit"

	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgAddAllowedDenom{}

// ValidateBasic does a sanity check on the provided data.
func (m MsgAddAllowedDenom) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.Wrapf(err, "invalid authority address")
	}

	msg := AllowedDenom{
		BankDenom:    m.BankDenom,
		DisplayDenom: m.DisplayDenom,
		Exponent:     m.Exponent,
	}

	return msg.Validate()
}

// GetSigners returns the expected signers for MsgAddAllowedDenom.
func (m MsgAddAllowedDenom) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

// GetSignBytes implements the LegacyMsg interface.
func (m MsgAddAllowedDenom) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// Route implements the LegacyMsg interface.
func (m MsgAddAllowedDenom) Route() string { return sdk.MsgTypeURL(&m) }

// Route implements the LegacyMsg interface.
func (m MsgAddAllowedDenom) Type() string { return sdk.MsgTypeURL(&m) }
