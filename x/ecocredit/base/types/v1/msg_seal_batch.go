package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
)

var _ legacytx.LegacyMsg = &MsgSealBatch{}

// Route implements the LegacyMsg interface.
func (m MsgSealBatch) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface. x
func (m MsgSealBatch) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgSealBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgSealBatch) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Issuer); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("issuer: %s", err)
	}

	if err := base.ValidateBatchDenom(m.BatchDenom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("batch denom: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgSealBatch.
func (m *MsgSealBatch) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Issuer)
	return []sdk.AccAddress{addr}
}
