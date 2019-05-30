package geo

import "github.com/cosmos/cosmos-sdk/codec"

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgStoreGeometry{}, "geo/MsgStoreGeometry", nil)
	cdc.RegisterConcrete(Geometry{}, "geo/Geometry", nil)
}

// generic sealed codec to be used throughout module
var moduleCdc *codec.Codec

func init() {
	cdc := codec.New()
	RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	moduleCdc = cdc.Seal()
}
