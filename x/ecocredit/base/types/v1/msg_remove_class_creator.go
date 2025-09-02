package v1

import (
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgRemoveClassCreator{}

// Route implements the LegacyMsg interface.
func (m MsgRemoveClassCreator) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgRemoveClassCreator) Type() string { return sdk.MsgTypeURL(&m) }

// // GetSignBytes implements the LegacyMsg interface.

// ValidateBasic does a sanity check on the provided data.
func (m *MsgRemoveClassCreator) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(err, "invalid authority address")
	}

	_, err := sdk.AccAddressFromBech32(m.Creator)
	return err
}
