package v1alpha1

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterTypes(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateClass{},
		&MsgCreateBatch{},
		&MsgSend{},
		&MsgRetire{},
		&MsgCancel{},
		&MsgUpdateClassAdmin{},
		&MsgUpdateClassIssuers{},
		&MsgUpdateClassMetadata{},
	)
}
