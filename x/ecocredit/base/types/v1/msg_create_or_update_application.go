package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
)

var _ legacytx.LegacyMsg = &MsgCreateOrUpdateApplication{}

// Route implements the LegacyMsg interface.
func (m *MsgCreateOrUpdateApplication) Route() string { return sdk.MsgTypeURL(m) }

// Type implements the LegacyMsg interface.
func (m *MsgCreateOrUpdateApplication) Type() string { return sdk.MsgTypeURL(m) }

// GetSignBytes implements the LegacyMsg interface.
func (m *MsgCreateOrUpdateApplication) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateOrUpdateApplication) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.ProjectAdmin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("project admin: %s", err)
	}

	if err := base.ValidateProjectID(m.ProjectId); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("project id: %s", err)
	}

	if err := base.ValidateClassID(m.ClassId); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("class id: %s", err)
	}

	if len(m.Metadata) > base.MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrapf("metadata: max length %d", base.MaxMetadataLength)
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateOrUpdateApplication.
func (m *MsgCreateOrUpdateApplication) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.ProjectAdmin)}
}
