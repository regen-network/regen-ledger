package basket

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var (
	_ legacytx.LegacyMsg = &MsgCreate{}
)

const nameMaxLen = 32
const displayNameMinLen = 3
const displayNameMaxLen = 32
const exponentMax = 32
const creditTNameMaxLen = 32

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgCreate) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Curator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap("malformed curator address " + err.Error())
	}
	// TODO: add proper validation once we will have proper requirements.
	// https://github.com/regen-network/regen-ledger/issues/732
	if m.Name == "" || len(m.Name) > nameMaxLen {
		return sdkerrors.ErrInvalidRequest.Wrap("name must not be empty and must not be longer than 32 characters long")
	}
	if len(m.DisplayName) < displayNameMinLen || len(m.DisplayName) > displayNameMaxLen {
		return sdkerrors.ErrInvalidRequest.Wrapf("display_name must be between %d and %d characters long", displayNameMinLen, displayNameMaxLen)
	}
	if m.Exponent > exponentMax {
		return sdkerrors.ErrInvalidRequest.Wrapf("exponent must not be bigger than %d", exponentMax)
	}
	if m.CreditTypeName == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("credit_type_name must be defined")
	}
	if len(m.CreditTypeName) > creditTNameMaxLen {
		return sdkerrors.ErrInvalidRequest.Wrapf("credit_type_name must not be longer than %d", creditTNameMaxLen)
	}

	if len(m.AllowedClasses) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("allowed_classes is required")
	}
	for i := range m.AllowedClasses {
		if m.AllowedClasses[i] == "" {
			return sdkerrors.ErrInvalidRequest.Wrapf("allowed_classes[%d] must be defined", i)
		}
	}
	if err := m.Fee.Validate(); err != nil {
		return err
	}

	return nil
}

// Validate additional validation with access to the state data.
func (m MsgCreate) Validate(minFee sdk.Coin) error {
	// TODO: update function to do stateful validation
	return nil
}

// GetSigners returns the expected signers for MsgCreate.
func (m MsgCreate) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Curator)
	return []sdk.AccAddress{addr}
}

// GetSignBytes Implements LegacyMsg.
func (m MsgCreate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// Route Implements LegacyMsg.
func (m MsgCreate) Route() string { return sdk.MsgTypeURL(&m) }

// Type Implements LegacyMsg.
func (m MsgCreate) Type() string { return sdk.MsgTypeURL(&m) }
