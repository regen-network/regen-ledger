package data

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterTypes(registry types.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers all the necessary data module concrete
// types with the provided codec reference.
// These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAnchor{}, "regen-ledger/MsgAnchor", nil)
	cdc.RegisterConcrete(&MsgAttest{}, "regen-ledger/MsgAttest", nil)
	cdc.RegisterConcrete(&MsgDefineResolver{}, "regen-ledger/MsgDefineResolver", nil)
	cdc.RegisterConcrete(&MsgRegisterResolver{}, "regen-ledger/MsgRegisterResolver", nil)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
}
