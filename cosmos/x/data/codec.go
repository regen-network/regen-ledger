package data

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	//cdc.RegisterConcrete(MsgRegisterSchema{}, "data/RegisterSchema", nil)
	cdc.RegisterConcrete(MsgStoreData{}, "data/MsgStoreData", nil)
	cdc.RegisterConcrete(MsgStoreDataPointer{}, "data/MsgStoreDataPointer", nil)
}
