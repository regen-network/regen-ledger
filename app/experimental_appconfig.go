// +build experimental

package app

import (
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
		data.Module{},
		ecocredit.Module{},
		group.Module{},
	}
}

func setCustomModules(app *RegenApp, interfaceRegistry types.InterfaceRegistry) {

	/* New Module Wiring START */
	newModuleManager := servermodule.NewManager(app.BaseApp, codec.NewProtoCodec(interfaceRegistry))

	// BEGIN HACK: this is a total, ugly hack until x/auth supports ADR 033 or we have a suitable alternative
	groupModule := group.Module{AccountKeeper: app.AccountKeeper}
	// use a separate newModules from the global NewModules here because we need to pass state into the group module
	newModules := []moduletypes.Module{
		ecocredit.Module{},
		data.Module{},
		groupModule,
	}
	err := newModuleManager.RegisterModules(newModules)
	if err != nil {
		panic(err)
	}
	// END HACK

	err = newModuleManager.CompleteInitialization()
	if err != nil {
		panic(err)
	}
	/* New Module Wiring END */
}
