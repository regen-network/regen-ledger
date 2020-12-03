package app

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/server"
)

// WasmWrapper allows us to use namespacing in the config file
// This is only used for parsing in the app, x/wasm expects WasmConfig
type WasmWrapper struct {
	Wasm wasm.Config `mapstructure:"wasm"`
}

func getWasmConfig() wasm.Config {
	wasmWrap := WasmWrapper{Wasm: wasm.DefaultWasmConfig()}
	ctx := server.NewDefaultContext()
	err := ctx.Viper.Unmarshal(&wasmWrap)
	if err != nil {
		panic("error while reading wasm config: " + err.Error())
	}

	return wasmWrap.Wasm
}
