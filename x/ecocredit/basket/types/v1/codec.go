package v1

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterTypes(registry codectypes.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreate{}, "regen.basket/MsgCreate", nil)
	cdc.RegisterConcrete(&MsgPut{}, "regen.basket/MsgPut", nil)
	cdc.RegisterConcrete(&MsgTake{}, "regen.basket/MsgTake", nil)
	cdc.RegisterConcrete(&MsgUpdateCurator{}, "regen.basket/MsgUpdateCurator", nil)
	cdc.RegisterConcrete(&MsgUpdateBasketFee{}, "regen.basket/MsgUpdateBasketFee", nil)
	cdc.RegisterConcrete(&MsgUpdateDateCriteria{}, "regen.basket/MsgUpdateDateCriteria", nil)
}
