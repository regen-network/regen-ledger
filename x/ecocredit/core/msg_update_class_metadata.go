package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgUpdateClassMetadata{}

func (m MsgUpdateClassMetadata) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassMetadata) Type() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassMetadata) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

func (m *MsgUpdateClassMetadata) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	if err := ValidateClassId(m.ClassId); err != nil {
		return err
	}

	if len(m.NewMetadata) > MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrap("credit class metadata")
	}

	return nil
}

func (m *MsgUpdateClassMetadata) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}
