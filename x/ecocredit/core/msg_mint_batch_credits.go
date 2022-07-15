package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgMintBatchCredits{}

// Route implements the LegacyMsg interface.
func (m MsgMintBatchCredits) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgMintBatchCredits) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgMintBatchCredits) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgMintBatchCredits) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Issuer); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("issuer: %s", err)
	}

	if m.BatchDenom == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("batch denom cannot be empty")
	}

	if err := ValidateBatchDenom(m.BatchDenom); err != nil {
		return err
	}

	if len(m.Issuance) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("issuance cannot be empty")
	}

	for i, issuance := range m.Issuance {
		if err := issuance.Validate(); err != nil {
			return sdkerrors.Wrapf(err, "issuance[%d]", i)
		}
	}

	if m.OriginTx == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("origin tx cannot be empty")
	}

	return m.OriginTx.Validate()
}

// GetSigners returns the expected signers for MsgMintBatchCredits.
func (m *MsgMintBatchCredits) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Issuer)
	return []sdk.AccAddress{addr}
}
