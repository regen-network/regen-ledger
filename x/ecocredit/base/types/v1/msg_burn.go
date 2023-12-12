package v1

import (
	"fmt"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgBurnRegen{}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgBurnRegen) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(err, "invalid from address")
	}
	amount, ok := math.NewIntFromString(m.Amount)
	if !ok {
		return fmt.Errorf("invalid amount: %s", m.Amount)
	}
	if !amount.IsPositive() {
		return fmt.Errorf("amount must be positive: %s", m.Amount)
	}
	if len(m.Reason) > MAX_REASON_LEN {
		return fmt.Errorf("reason must be less than 256 characters")
	}
	return nil
}

const MAX_REASON_LEN = 256

// GetSigners returns the expected signers for MsgBurnRegen.
func (m *MsgBurnRegen) GetSigners() []sdk.AccAddress {
	addr := sdk.MustAccAddressFromBech32(m.From)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgBurnRegen) Route() string { return sdk.MsgTypeURL(m) }

// Type implements the LegacyMsg interface.
func (m MsgBurnRegen) Type() string { return sdk.MsgTypeURL(m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgBurnRegen) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
