package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/v4/base"
)

var _ sdk.Msg = &MsgSealBatch{}

// Route implements the LegacyMsg interface.
func (m MsgSealBatch) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface. x
func (m MsgSealBatch) Type() string { return sdk.MsgTypeURL(&m) }

// ValidateBasic does a sanity check on the provided data.
func (m *MsgSealBatch) ValidateBasic() error {
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
