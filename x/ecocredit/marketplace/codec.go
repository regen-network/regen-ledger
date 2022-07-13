package marketplace

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func RegisterTypes(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*govv1beta1.Content)(nil), &AllowDenomProposal{})
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgBuyDirect{}, "regen.marketplace/MsgBuyDirect", nil)
	cdc.RegisterConcrete(&MsgSell{}, "regen.marketplace/MsgSell", nil)
	cdc.RegisterConcrete(&MsgCancelSellOrder{}, "regen.marketplace/MsgCancelSellOrder", nil)
	cdc.RegisterConcrete(&MsgUpdateSellOrders{}, "regen.marketplace/MsgUpdateSellOrders", nil)
}

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	RegisterLegacyAminoCodec(amino)
}
