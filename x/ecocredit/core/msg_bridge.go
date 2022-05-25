package core

import (
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// TODO: remove after we open governance changes for bridge target
// https://github.com/regen-network/regen-ledger/issues/1119

const (
	BRIDGE_TARGET string = "polygon"
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

	if _, err := sdk.AccAddressFromBech32(m.Holder); err != nil {
		return sdkerrors.Wrap(err, "holder")
	}

	if len(m.Credits) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("credits should not be empty")
	}

	for _, credit := range m.Credits {
		if err := ValidateBatchDenom(credit.BatchDenom); err != nil {
			return err
		}

		if _, err := math.NewPositiveDecFromString(credit.Amount); err != nil {
			return err
		}
	}

	if m.Target != BRIDGE_TARGET {
		return sdkerrors.ErrInvalidRequest.Wrapf("expected %s got %s", BRIDGE_TARGET, m.Target)
	}

	if !isValidEthereumAddress(m.Recipient) {
		return sdkerrors.ErrInvalidAddress.Wrapf("%s is not a valid ethereum address", m.Recipient)
	}

	if !isValidEthereumAddress(m.Contract) {
		return sdkerrors.ErrInvalidAddress.Wrapf("%s is not a valid ethereum address", m.Contract)
	}

	return nil
}

// GetSigners returns the expected signers for MsgCancel.
func (m *MsgBridge) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Holder)
	return []sdk.AccAddress{addr}
}

func isValidEthereumAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}
