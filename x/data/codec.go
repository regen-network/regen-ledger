package data

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterTypes(registry types.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &Msg_ServiceDesc)
}
