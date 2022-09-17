package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit/basket"

	"github.com/regen-network/regen-ledger/x/ecocredit/base"
)

const descrMaxLen = 256

var _ legacytx.LegacyMsg = &MsgCreate{}

// Route implements LegacyMsg.
func (m MsgCreate) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements LegacyMsg.
func (m MsgCreate) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements LegacyMsg.
func (m MsgCreate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgCreate) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Curator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap("malformed curator address: " + err.Error())
	}

	if err := basket.ValidateBasketName(m.Name); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("name: %s", err)
	}

	if len(m.Description) > descrMaxLen {
		return sdkerrors.ErrInvalidRequest.Wrapf("description length cannot be greater than %d characters", descrMaxLen)
	}

	if err := base.ValidateCreditTypeAbbreviation(m.CreditTypeAbbrev); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("credit type abbrev: %s", err)
	}

	if len(m.AllowedClasses) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("allowed classes cannot be empty")
	}

	for i := range m.AllowedClasses {
		if err := base.ValidateClassID(m.AllowedClasses[i]); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("allowed classes [%d]: %s", i, err)
		}
	}

	if err := m.DateCriteria.Validate(); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("invalid date criteria: %s", err)
	}

	// In the next version of the basket package, this field will be updated to
	// a single Coin rather than a list of Coins. In the meantime, the message
	// will fail basic validation if more than one Coin is provided.
	if len(m.Fee) > 1 {
		return sdkerrors.ErrInvalidRequest.Wrap("more than one fee is not allowed")
	}

	return m.Fee.Validate()
}

// GetSigners returns the expected signers for MsgCreate.
func (m MsgCreate) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Curator)
	return []sdk.AccAddress{addr}
}
