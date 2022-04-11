package core

import (
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgUpdateProjectMetadata{}

func (m MsgUpdateProjectMetadata) ValidateBasic() error {
	if _, err := types.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}
	if len(m.NewMetadata) > MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrapf("create project metadata: max length is %d", MaxMetadataLength)
	}
	if err := ValidateProjectID(m.ProjectId); err != nil {
		return err
	}
	return nil
}

func (m MsgUpdateProjectMetadata) GetSigners() []types.AccAddress {
	addr, _ := types.AccAddressFromBech32(m.Admin)
	return []types.AccAddress{addr}
}

func (m MsgUpdateProjectMetadata) GetSignBytes() []byte {
	return types.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgUpdateProjectMetadata) Route() string { return types.MsgTypeURL(&m) }

func (m MsgUpdateProjectMetadata) Type() string { return types.MsgTypeURL(&m) }
