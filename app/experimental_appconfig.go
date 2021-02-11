// +build !stable

package app

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	moduletypes "github.com/regen-network/regen-ledger/types/module"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	data "github.com/regen-network/regen-ledger/x/data/module"
	ecocredit "github.com/regen-network/regen-ledger/x/ecocredit/module"
	group "github.com/regen-network/regen-ledger/x/group/module"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{
		wasm.AppModuleBasic{},
		data.Module{},
		ecocredit.Module{},
		group.Module{},
	}
}

func setCustomModules(app *RegenApp, interfaceRegistry types.InterfaceRegistry) {

	/* New Module Wiring START */
	newModuleManager := servermodule.NewManager(app.BaseApp, codec.NewProtoCodec(interfaceRegistry))

	err := newModuleManager.RegisterModules([]moduletypes.Module{
		ecocredit.Module{},
		data.Module{},
		group.Module{},
	})

	if err != nil {
		panic(err)
	}

	err = newModuleManager.CompleteInitialization()
	if err != nil {
		panic(err)
	}
	/* New Module Wiring END */
}
