package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

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

	if len(m.Metadata) > MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrap("credit class metadata")
	}

	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	if len(m.Issuers) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("issuers cannot be empty")
	}

	if len(m.CreditTypeAbbrev) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("must specify a credit type abbreviation")
	}

	duplicateMap := make(map[string]bool)
	for _, issuer := range m.Issuers {

		if _, err := sdk.AccAddressFromBech32(issuer); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
		}

		if _, ok := duplicateMap[issuer]; ok {
			return sdkerrors.ErrInvalidRequest.Wrapf("duplicate issuer %s", issuer)
		}
		duplicateMap[issuer] = true
	}

	if m.Fee != nil {
		if m.Fee.Amount.IsZero() {
			return sdkerrors.ErrInsufficientFee.Wrap("must specify nonzero fee")
		}
		if err := sdk.ValidateDenom(m.Fee.Denom); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s", err.Error())
		}
	} else {
		return sdkerrors.ErrInvalidRequest.Wrap("fee must be specified")
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateClass.
func (m *MsgCreateClass) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}
