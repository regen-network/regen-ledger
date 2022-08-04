package core

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgCreateClass{}

// Route implements the LegacyMsg interface.
func (m MsgCreateClass) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCreateClass) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCreateClass) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateClass) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("admin: %s", err)
	}

	if len(m.Issuers) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("issuers cannot be empty")
	}

	duplicateMap := make(map[string]bool)
	for i, issuer := range m.Issuers {
		issuerIndex := fmt.Sprintf("issuers[%d]", i)

		if _, err := sdk.AccAddressFromBech32(issuer); err != nil {
			return sdkerrors.ErrInvalidAddress.Wrapf("%s: %s", issuerIndex, err)
		}

		if _, ok := duplicateMap[issuer]; ok {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: duplicate address %s", issuerIndex, issuer)
		}

		duplicateMap[issuer] = true
	}

	if len(m.Metadata) > MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrapf("metadata: max length %d", MaxMetadataLength)
	}

	if err := ValidateCreditTypeAbbreviation(m.CreditTypeAbbrev); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	if m.Fee != nil {
		if m.Fee.Denom == "" {
			return sdkerrors.ErrInvalidRequest.Wrap("fee denom cannot be empty")
		}

		if err := sdk.ValidateDenom(m.Fee.Denom); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s", err.Error())
		}

		if m.Fee.Amount.IsNil() {
			return sdkerrors.ErrInvalidRequest.Wrap("fee amount cannot be empty")
		}

		if !m.Fee.Amount.IsPositive() {
			return sdkerrors.ErrInsufficientFee.Wrap("fee amount must be positive")
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateClass.
func (m *MsgCreateClass) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}
