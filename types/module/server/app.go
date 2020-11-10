package server

import "github.com/cosmos/cosmos-sdk/baseapp"

type ModuleApp struct {
	baseapp.BaseApp
	modules map[string]Module
}

func NewModuleApp(modules map[string]Module) *ModuleApp {
	return &ModuleApp{modules: modules}
}
