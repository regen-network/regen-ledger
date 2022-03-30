package core

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

var _, _, _, _ legacytx.LegacyMsg = &MsgAddCreditType{}, &MsgUpdateClassCreatorAllowlist{}, &MsgToggleAllowlist{}, &MsgUpdateClassFee{}

func (m MsgAddCreditType) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m MsgAddCreditType) GetSigners() []types.AccAddress {
	//TODO implement me
	panic("implement me")
}

func (m MsgAddCreditType) GetSignBytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (m MsgAddCreditType) Route() string {
	//TODO implement me
	panic("implement me")
}

func (m MsgAddCreditType) Type() string {
	//TODO implement me
	panic("implement me")
}

func (m MsgUpdateClassCreatorAllowlist) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m MsgUpdateClassCreatorAllowlist) GetSigners() []types.AccAddress {
	//TODO implement me
	panic("implement me")
}

func (m MsgUpdateClassCreatorAllowlist) GetSignBytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (m MsgUpdateClassCreatorAllowlist) Route() string {
	//TODO implement me
	panic("implement me")
}

func (m MsgUpdateClassCreatorAllowlist) Type() string {
	//TODO implement me
	panic("implement me")
}

func (m MsgToggleAllowlist) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m MsgToggleAllowlist) GetSigners() []types.AccAddress {
	//TODO implement me
	panic("implement me")
}

func (m MsgToggleAllowlist) GetSignBytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (m MsgToggleAllowlist) Route() string {
	//TODO implement me
	panic("implement me")
}

func (m MsgToggleAllowlist) Type() string {
	//TODO implement me
	panic("implement me")
}

func (m MsgUpdateClassFee) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m MsgUpdateClassFee) GetSigners() []types.AccAddress {
	//TODO implement me
	panic("implement me")
}

func (m MsgUpdateClassFee) GetSignBytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (m MsgUpdateClassFee) Route() string {
	//TODO implement me
	panic("implement me")
}

func (m MsgUpdateClassFee) Type() string {
	//TODO implement me
	panic("implement me")
}
