package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddAllowedBridgeChain{}

// Route implements the LegacyMsg interface.
func (m MsgAddAllowedBridgeChain) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgAddAllowedBridgeChain) Type() string { return sdk.MsgTypeURL(&m) }

// ValidateBasic does a sanity check on the provided data.
func (m *MsgAddAllowedBridgeChain) ValidateBasic() error {
	if m.ChainName == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("chain_name cannot be empty")
	}
	return nil
}

// GetSigners returns the expected signers for MsgAddAllowedBridgeChain.
func (m *MsgAddAllowedBridgeChain) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
