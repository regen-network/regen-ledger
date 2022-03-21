package core

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterTypes(registry codectypes.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateClass{}, "regen.core/MsgCreateClass", nil)
	cdc.RegisterConcrete(&MsgCancel{}, "regen.core/MsgCancel", nil)
	cdc.RegisterConcrete(&MsgCreateBatch{}, "regen.core/MsgCreateBatch", nil)
	cdc.RegisterConcrete(&MsgRetire{}, "regen.core/MsgRetire", nil)
	cdc.RegisterConcrete(&MsgSend{}, "regen.core/MsgSend", nil)
	cdc.RegisterConcrete(&MsgUpdateClassAdmin{}, "regen.core/MsgUpdateClassAdmin", nil)
	cdc.RegisterConcrete(&MsgUpdateClassMetadata{}, "regen.core/MsgUpdateClassMetadata", nil)
}
