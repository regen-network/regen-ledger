package wasm

import (
	"github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/CosmWasm/wasmd/x/wasm/types"
)

const (
	ModuleName = types.ModuleName
	StoreKey   = types.StoreKey
)

var (
	NewKeeper = keeper.NewKeeper
)

type (
	Config = types.WasmConfig
	Keeper = keeper.Keeper
	Option = keeper.Option
)
