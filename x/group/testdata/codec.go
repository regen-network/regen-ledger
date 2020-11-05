package testdata

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// RegisterCodec registers all the necessary crisis module concrete types and
// interfaces with the provided codec reference.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgPropose{}, "testdata/MsgPropose", nil)
	// oh man... amino
	cdc.RegisterInterface((*isMyAppMsg_Sum)(nil), nil)
	cdc.RegisterConcrete(&MsgAlwaysSucceed{}, "testdata/MsgAlwaysSucceed", nil)
	cdc.RegisterConcrete(&MsgAlwaysFail{}, "testdata/MsgAlwaysFail", nil)
	cdc.RegisterConcrete(&MyAppMsg_A{}, "testdata/MyAppMsg_A", nil)
	cdc.RegisterConcrete(&MyAppMsg_B{}, "testdata/MyAppMsg_B", nil)
}

var (
	amino = codec.New()

	// moduleCdc references the global x/transfer module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding as Amino
	// is still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/transfer and
	// defined at the application level.
	moduleCdc = codec.NewHybridCodec(amino, cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	codec.RegisterCrypto(amino)
	amino.Seal()
}
