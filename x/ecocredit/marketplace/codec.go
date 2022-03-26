package marketplace

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterTypes(registry codectypes.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAllowAskDenom{}, "regen.marketplace/MsgAllowAskDenom", nil)
	cdc.RegisterConcrete(&MsgBuy{}, "regen.marketplace/MsgBuy", nil)
	cdc.RegisterConcrete(&MsgSell{}, "regen.marketplace/MsgSell", nil)
	cdc.RegisterConcrete(&MsgCancelSellOrder{}, "regen.marketplace/MsgCancelSellOrder", nil)
	cdc.RegisterConcrete(&MsgUpdateSellOrders{}, "regen.marketplace/MsgUpdateSellOrders", nil)
}
