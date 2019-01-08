package consortium

import (
"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(ActionScheduleUpgrade{}, "consortium/ActionScheduleUpgrade", nil)
	cdc.RegisterConcrete(ActionChangeValidatorSet{}, "consortium/ActionChangeValidatorSet", nil)
}
