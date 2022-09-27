package marketplace

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterTypes(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*govtypes.Content)(nil), &AllowDenomProposal{})
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgBuyDirect{}, "regen.marketplace/MsgBuyDirect", nil)
	cdc.RegisterConcrete(&MsgSell{}, "regen.marketplace/MsgSell", nil)
	cdc.RegisterConcrete(&MsgCancelSellOrder{}, "regen.marketplace/MsgCancelSellOrder", nil)
	cdc.RegisterConcrete(&MsgUpdateSellOrders{}, "regen.marketplace/MsgUpdateSellOrders", nil)
	cdc.RegisterConcrete(&AllowDenomProposal{}, "regen.marketplace/AllowDenomProposal", nil)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
}
