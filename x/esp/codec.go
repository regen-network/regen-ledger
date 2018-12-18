package esp

import (
"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgRegisterESPVersion{}, "esp/MsgRegisterESPVersion", nil)
	cdc.RegisterConcrete(MsgReportESPResult{}, "esp/MsgReportESPResult", nil)
}
