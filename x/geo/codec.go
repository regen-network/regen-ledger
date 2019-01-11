package geo

import "github.com/cosmos/cosmos-sdk/codec"

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgStoreGeometry{}, "geo/MsgStoreGeometry", nil)
	cdc.RegisterConcrete(Geometry{}, "geo/Geometry", nil)
}
