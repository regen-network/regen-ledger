package v1

import (
	"fmt"

	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/v2/eth"
	"github.com/regen-network/regen-ledger/types/v2/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/base"
)

var _ sdk.Msg = &MsgBridge{}

// Route implements the LegacyMsg interface.
func (m MsgBridge) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgBridge) Type() string { return sdk.MsgTypeURL(&m) }

// ValidateBasic does a sanity check on the provided data.
func (m *MsgBridge) ValidateBasic() error {
	if m.Target == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("target cannot be empty")
	}

	if m.Recipient == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("recipient cannot be empty")
	}

	if !eth.IsValidAddress(m.Recipient) {
		return sdkerrors.ErrInvalidAddress.Wrap("recipient must be a valid ethereum address")
	}

	if len(m.Credits) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("credits cannot be empty")
	}

	for i, credit := range m.Credits {
		creditIndex := fmt.Sprintf("credits[%d]", i)

		if err := base.ValidateBatchDenom(credit.BatchDenom); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: batch denom: %s", creditIndex, err)
		}

		if credit.Amount == "" {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: amount cannot be empty", creditIndex)
		}

		if _, err := math.NewPositiveDecFromString(credit.Amount); err != nil {
			return errors.Wrapf(err, "%s: amount", creditIndex)
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgCancel.
func (m *MsgBridge) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}
