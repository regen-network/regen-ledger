package core

import (
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgUpdateProjectMetadata{}

func (m MsgUpdateProjectMetadata) Route() string { return types.MsgTypeURL(&m) }

func (m MsgUpdateProjectMetadata) Type() string { return types.MsgTypeURL(&m) }

func (m MsgUpdateProjectMetadata) GetSignBytes() []byte {
	return types.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgUpdateProjectMetadata) ValidateBasic() error {
	if _, err := types.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("admin: %s", err)
	}

	if err := ValidateProjectId(m.ProjectId); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	if len(m.NewMetadata) > MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrapf("metadata: max length is %d", MaxMetadataLength)
	}

	return nil
}

func (m MsgUpdateProjectMetadata) GetSigners() []types.AccAddress {
	addr, _ := types.AccAddressFromBech32(m.Admin)
	return []types.AccAddress{addr}
}
