package testdata

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
)

// RegisterLegacyAminoCodec registers all the necessary module concrete types and
// interfaces with the provided codec reference.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgPropose{}, "testdata/MsgPropose", nil)
	// oh man... amino
	cdc.RegisterInterface((*isMyAppMsg_Sum)(nil), nil)
	cdc.RegisterConcrete(&MsgAlwaysSucceed{}, "testdata/MsgAlwaysSucceed", nil)
	cdc.RegisterConcrete(&MsgAlwaysFail{}, "testdata/MsgAlwaysFail", nil)
	cdc.RegisterConcrete(&MyAppMsg_A{}, "testdata/MyAppMsg_A", nil)
	cdc.RegisterConcrete(&MyAppMsg_B{}, "testdata/MyAppMsg_B", nil)
}

var (
	amino = codec.NewLegacyAmino()

	// moduleCdc references the global x/transfer module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding as Amino
	// is still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/transfer and
	// defined at the application level.
	moduleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
