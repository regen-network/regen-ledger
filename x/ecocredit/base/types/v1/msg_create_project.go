package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
)

const MaxReferenceIDLength = 32

var _ legacytx.LegacyMsg = &MsgCreateProject{}

// Route implements the LegacyMsg interface.
func (m MsgCreateProject) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCreateProject) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCreateProject) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateProject) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("admin: %s", err)
	}

	if err := base.ValidateClassID(m.ClassId); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("class id: %s", err)
	}

	if len(m.Metadata) > base.MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrapf("metadata: max length %d", base.MaxMetadataLength)
	}

	if err := base.ValidateJurisdiction(m.Jurisdiction); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("jurisdiction: %s", err)
	}

	if m.ReferenceId != "" && len(m.ReferenceId) > MaxReferenceIDLength {
		return ecocredit.ErrMaxLimit.Wrapf("reference id: max length %d", MaxReferenceIDLength)
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateProject.
func (m *MsgCreateProject) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}
