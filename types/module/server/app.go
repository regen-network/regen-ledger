package server

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

type ModuleApp struct {
	*baseapp.BaseApp
	modules map[string]Module
}

func NewModuleApp(name string, logger log.Logger, db dbm.DB, txDecoder sdk.TxDecoder, options ...func(app *baseapp.BaseApp)) *ModuleApp {
	return &ModuleApp{
		BaseApp: baseapp.NewBaseApp(name, logger, db, txDecoder, options...),
		modules: map[string]Module{},
	}
}
