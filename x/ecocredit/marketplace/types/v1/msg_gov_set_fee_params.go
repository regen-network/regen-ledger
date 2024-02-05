package v1

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgGovSetFeeParams{}

func (m *MsgGovSetFeeParams) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m *MsgGovSetFeeParams) GetSigners() []types.AccAddress {
	//TODO implement me
	panic("implement me")
}

func (m *MsgGovSetFeeParams) GetSignBytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (m *MsgGovSetFeeParams) Route() string {
	//TODO implement me
	panic("implement me")
}

func (m *MsgGovSetFeeParams) Type() string {
	//TODO implement me
	panic("implement me")
}
