package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgCreateProject{}

// Route implements the LegacyMsg interface.
func (m MsgCreateProject) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCreateProject) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCreateProject) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateProject) ValidateBasic() error {

	if _, err := sdk.AccAddressFromBech32(m.Issuer); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	if err := ValidateClassID(m.ClassId); err != nil {
		return err
	}

	if len(m.Metadata) > MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrap("create project metadata")
	}

	if err := ValidateJurisdiction(m.ProjectJurisdiction); err != nil {
		return err
	}

	if m.ProjectId != "" {
		if err := ValidateProjectID(m.ProjectId); err != nil {
			return err
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateProject.
func (m *MsgCreateProject) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Issuer)
	return []sdk.AccAddress{addr}
}
