package ecocredit

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
