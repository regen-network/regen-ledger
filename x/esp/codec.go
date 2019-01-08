package esp

import (
"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(ActionRegisterESPVersion{}, "esp/ActionRegisterESPVersion", nil)
	cdc.RegisterConcrete(ActionReportESPResult{}, "esp/ActionReportESPResult", nil)
	cdc.RegisterConcrete(ESPVersionSpec{}, "esp/ESPVersionSpec", nil)
	cdc.RegisterConcrete(ESPResult{}, "esp/ESPResult", nil)
}
