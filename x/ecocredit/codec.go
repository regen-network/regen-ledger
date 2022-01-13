package ecocredit

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers all the necessary ecocredit module concrete
// types with the provided codec reference.
// These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateClass{}, "regen-ledger/MsgCreateClass", nil)
	cdc.RegisterConcrete(&MsgCreateBatch{}, "regen-ledger/MsgCreateBatch", nil)
	cdc.RegisterConcrete(&MsgSend{}, "regen-ledger/MsgSend", nil)
	cdc.RegisterConcrete(&MsgRetire{}, "regen-ledger/MsgRetire", nil)
	cdc.RegisterConcrete(&MsgCancel{}, "regen-ledger/MsgCancel", nil)
	cdc.RegisterConcrete(&MsgCreateProject{}, "regen-ledger/MsgCreateProject", nil)
	cdc.RegisterConcrete(&MsgAddToBasket{}, "regen-ledger/MsgAddToBasket", nil)
	cdc.RegisterConcrete(&MsgPickFromBasket{}, "regen-ledger/MsgPickFromBasket", nil)
	cdc.RegisterConcrete(&MsgCreateBasket{}, "regen-ledger/MsgCreateBasket", nil)
	cdc.RegisterConcrete(&MsgTakeFromBasket{}, "regen-ledger/MsgTakeFromBasket", nil)
}

func RegisterTypes(registry codectypes.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
}
