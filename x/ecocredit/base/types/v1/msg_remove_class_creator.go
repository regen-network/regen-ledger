package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgRemoveClassCreator{}

// Route implements the LegacyMsg interface.
func (m MsgRemoveClassCreator) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgRemoveClassCreator) Type() string { return sdk.MsgTypeURL(&m) }

// // GetSignBytes implements the LegacyMsg interface.
