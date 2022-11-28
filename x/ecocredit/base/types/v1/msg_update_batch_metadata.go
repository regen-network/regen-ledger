package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
)

var _ legacytx.LegacyMsg = &MsgUpdateBatchMetadata{}

// Route implements the LegacyMsg interface.
func (m MsgUpdateBatchMetadata) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgUpdateBatchMetadata) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgUpdateBatchMetadata) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgUpdateBatchMetadata) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Issuer); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("issuer: %s", err)
	}

	if err := base.ValidateBatchDenom(m.BatchDenom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("batch denom: %s", err)
	}

	// we allow removing metadata for class and project but not for batch
	if m.NewMetadata == "" {
		return sdkerrors.ErrInvalidRequest.Wrapf("metadata: cannot be empty")
	}

	if len(m.NewMetadata) > base.MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrapf("metadata: max length %d", base.MaxMetadataLength)
	}

	return nil
}

// GetSigners returns the expected signers for the message.
func (m *MsgUpdateBatchMetadata) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Issuer)
	return []sdk.AccAddress{addr}
}
