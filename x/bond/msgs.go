package bond

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/regen-network/regen-ledger/types/math"
)

var (
	_, _ sdk.Msg            = &MsgIssueBond{}, &MsgSellBond{}
	_, _ legacytx.LegacyMsg = &MsgIssueBond{}, &MsgSellBond{}
)

// MaxMetadataLength defines the max length of the metadata bytes field
// for the issue-bond message.
const MaxMetadataLength = 256

// Route implements the LegacyMsg interface.
func (m MsgIssueBond) Route() string { return sdk.MsgTypeURL(&m) }

// ValidateBasic does a sanity check on the provided data.
func (m *MsgIssueBond) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Holder); err != nil {
		return sdkerrors.Wrap(err, "holder")
	}

	if len(m.Metadata) > MaxMetadataLength {
		return ErrMaxLimit.Wrap("Bond metadata")
	}

	if m.MaturityDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("must provide a maturity date for the bond to issue")
	}

	if m.IssuanceDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("must provide a issuance date for the bond to issue")
	}

	if m.MaturityDate.Before(*m.IssuanceDate) {
		return sdkerrors.ErrInvalidRequest.Wrapf(
			"the bond issuance date (%s) must be the same as or after the bond maturity date (%s)",
			m.MaturityDate.Format("2006-01-02"),
			m.IssuanceDate.Format("2006-01-02"))
	}

	if m.FaceValue != "" {
		if _, err := math.NewNonNegativeDecFromString(m.FaceValue); err != nil {
			return err
		}
	} else {
		return sdkerrors.ErrInvalidRequest.Wrap("FaceValue has to be positive decimal number")
	}

	if m.EmissionDenom == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("EmissionDenom has to not empty string")
	}

	if m.Name == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("Name has to not empty string")
	}

	if m.FaceCurrency == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("FaceCurrency has to not empty string")
	}

	if m.CouponRate == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("CouponRate has to not empty string")
	}

	if m.CouponFrequency == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("CouponFrequency has to not empty string")
	}

	if m.Project == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("Project has to not empty string")
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateClass.
func (m *MsgIssueBond) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Holder)
	return []sdk.AccAddress{addr}
}

// Type implements the LegacyMsg interface.
func (m MsgIssueBond) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgIssueBond) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// Route implements the LegacyMsg interface.
func (m MsgSellBond) Route() string { return sdk.MsgTypeURL(&m) }

// ValidateBasic does a sanity check on the provided data.
func (m *MsgSellBond) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Buyer); err != nil {
		return sdkerrors.Wrap(err, "Buyer address is not valid")
	}

	if m.Amount != "" {
		if _, err := math.NewNonNegativeDecFromString(m.Amount); err != nil {
			return err
		}
	} else {
		return sdkerrors.ErrInvalidRequest.Wrap("Amount has to be positive decimal number")
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateClass.
func (m *MsgSellBond) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Holder)
	return []sdk.AccAddress{addr}
}

// Type implements the LegacyMsg interface.
func (m MsgSellBond) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgSellBond) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
