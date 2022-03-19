package marketplace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgAllowAskDenom{}

// Route implements the LegacyMsg interface.
func (m MsgAllowAskDenom) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgAllowAskDenom) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgAllowAskDenom) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgAllowAskDenom) ValidateBasic() error {

	if _, err := sdk.AccAddressFromBech32(m.RootAddress); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	return sdk.ValidateDenom(m.Denom)
}

// GetSigners returns the expected signers for MsgAllowAskDenom.
func (m *MsgAllowAskDenom) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.RootAddress)
	return []sdk.AccAddress{addr}
}
