package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
)

var _ legacytx.LegacyMsg = &MsgCreateUnregisteredProject{}

// Route implements the LegacyMsg interface.
func (m *MsgCreateUnregisteredProject) Route() string { return sdk.MsgTypeURL(m) }

// Type implements the LegacyMsg interface.
func (m *MsgCreateUnregisteredProject) Type() string { return sdk.MsgTypeURL(m) }

// GetSignBytes implements the LegacyMsg interface.
func (m *MsgCreateUnregisteredProject) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateUnregisteredProject) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("admin: %s", err)
	}

	if err := base.ValidateJurisdiction(m.Jurisdiction); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("jurisdiction: %s", err)
	}

	if len(m.Metadata) > base.MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrapf("metadata: max length %d", base.MaxMetadataLength)
	}

	if m.ReferenceId != "" && len(m.ReferenceId) > MaxReferenceIDLength {
		return ecocredit.ErrMaxLimit.Wrapf("reference id: max length %d", MaxReferenceIDLength)
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateUnregisteredProject.
func (m *MsgCreateUnregisteredProject) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Admin)}
}
