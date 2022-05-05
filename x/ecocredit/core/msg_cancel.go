package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgCancel{}

// Route implements the LegacyMsg interface.
func (m MsgCancel) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCancel) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCancel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCancel) ValidateBasic() error {

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
	return nil
}

// GetSigners returns the expected signers for MsgCancel.
func (m *MsgCancel) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Holder)
	return []sdk.AccAddress{addr}
}
