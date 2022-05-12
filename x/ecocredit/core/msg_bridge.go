package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgBridge{}

// Route implements the LegacyMsg interface.
func (m MsgBridge) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgBridge) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgBridge) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgBridge) ValidateBasic() error {

	if err := m.MsgCancel.ValidateBasic(); err != nil {
		return err
	}

	if len(m.BridgeContract) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("bridge contract should not be empty")
	}

	if len(m.BridgeRecipient) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("bridge recipient address should not be empty")
	}

	if len(m.BridgeTarget) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("bridge target should not be empty")
	}

	return nil
}

// GetSigners returns the expected signers for MsgCancel.
func (m *MsgBridge) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.MsgCancel.Holder)
	return []sdk.AccAddress{addr}
}
