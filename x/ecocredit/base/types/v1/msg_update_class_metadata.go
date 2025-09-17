package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/base"
)

var _ sdk.Msg = &MsgUpdateClassMetadata{}

func (m MsgUpdateClassMetadata) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassMetadata) Type() string { return sdk.MsgTypeURL(&m) }

func (m *MsgUpdateClassMetadata) ValidateBasic() error {
	if err := base.ValidateClassID(m.ClassId); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("class id: %s", err)
	}

	if len(m.NewMetadata) > base.MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrapf("metadata: max length %d", base.MaxMetadataLength)
	}

	return nil
}

func (m *MsgUpdateClassMetadata) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}
