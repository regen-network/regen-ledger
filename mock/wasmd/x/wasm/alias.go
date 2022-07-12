package wasm

import (
	"github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/CosmWasm/wasmd/x/wasm/types"
)

const (
	StoreKey = types.StoreKey
)

type (
	Config = types.WasmConfig
	Keeper = keeper.Keeper
	Option = keeper.Option
)
