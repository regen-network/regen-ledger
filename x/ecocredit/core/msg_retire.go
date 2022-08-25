package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgRetire{}

// Route implements the LegacyMsg interface.
func (m MsgRetire) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgRetire) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgRetire) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgRetire) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("owner: %s", err)
	}

	if len(m.Credits) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("credits cannot be empty")
	}

	for i, credits := range m.Credits {
		if err := credits.Validate(); err != nil {
			return sdkerrors.Wrapf(err, "credits[%d]", i)
		}
	}

	if err := ValidateJurisdiction(m.Jurisdiction); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("jurisdiction: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgRetire.
func (m *MsgRetire) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}
