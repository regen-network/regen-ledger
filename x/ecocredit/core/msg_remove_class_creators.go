package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgRemoveClassCreators{}

// Route implements the LegacyMsg interface.
func (m MsgRemoveClassCreators) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgRemoveClassCreators) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgRemoveClassCreators) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgRemoveClassCreators) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.Wrapf(err, "invalid authority address")
	}

	if len(m.Creators) == 0 {
		return sdkerrors.ErrInvalidType.Wrap("class creators cannot be empty")
	}

	for _, creator := range m.Creators {
		if _, err := sdk.AccAddressFromBech32(creator); err != nil {
			return err
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgRemoveClassCreators.
func (m *MsgRemoveClassCreators) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
