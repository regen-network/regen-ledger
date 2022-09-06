package v1

import (
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgRemoveAllowedBridgeChain{}

// Route implements the LegacyMsg interface.
func (m MsgRemoveAllowedBridgeChain) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgRemoveAllowedBridgeChain) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgRemoveAllowedBridgeChain) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgRemoveAllowedBridgeChain) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(err, "invalid authority address")
	}
	if m.ChainName == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("chain_name cannot be empty")
	}
	return nil
}

// GetSigners returns the expected signers for MsgRemoveAllowedBridgeChain.
func (m *MsgRemoveAllowedBridgeChain) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
