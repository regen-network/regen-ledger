package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgRemoveClassCreator{}

// Route implements the LegacyMsg interface.
func (m MsgRemoveClassCreator) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgRemoveClassCreator) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgRemoveClassCreator) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgRemoveClassCreator) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.Wrapf(err, "invalid authority address")
	}

	if len(m.Creator) == 0 {
		return sdkerrors.ErrInvalidType.Wrap("class creators cannot be empty")
	}

	for _, creator := range m.Creator {
		if _, err := sdk.AccAddressFromBech32(creator); err != nil {
			return err
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgRemoveClassCreator.
func (m *MsgRemoveClassCreator) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
