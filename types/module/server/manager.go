package server

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogogrpc "github.com/gogo/protobuf/grpc"
)

type Manager struct {
	baseApp baseapp.BaseApp
	cdc     *codec.ProtoCodec
	keys    map[string]ModuleKey
	router  *router
}

func NewManager(baseApp baseapp.BaseApp, cdc *codec.ProtoCodec) *Manager {
	return &Manager{
		baseApp: baseApp,
		cdc:     cdc,
		keys:    map[string]ModuleKey{},
		router:  &router{handlers: map[string]handler{}},
	}
}

func (mm *Manager) RegisterModules(modules map[string]Module) error {
	for _, mod := range modules {
		mod.RegisterTypes(mm.cdc.InterfaceRegistry())
	}

	for name, mod := range modules {
		invokerFactory := mm.router.invokerFactory(name)

		key := RootModuleKey{
			moduleName:     name,
			invokerFactory: invokerFactory,
		}

		if _, found := mm.keys[name]; found {
			return fmt.Errorf("module named %s defined twice", name)
		}

		mm.keys[name] = key
		mm.baseApp.MountStore(key, sdk.StoreTypeIAVL)

		msgRegistrar := registrar{
			router:       mm.router,
			baseServer:   mm.baseApp.MsgServiceRouter(),
			commitWrites: true,
		}

		queryRegistrar := registrar{
			router:       mm.router,
			baseServer:   mm.baseApp.GRPCQueryRouter(),
			commitWrites: true,
		}

		cfg := configurator{
			msgServer:   msgRegistrar,
			queryServer: queryRegistrar,
			key:         key,
			cdc:         mm.cdc,
		}

		mod.RegisterServices(cfg)
	}

	return nil
}

type configurator struct {
	msgServer   gogogrpc.Server
	queryServer gogogrpc.Server
	key         ModuleKey
	cdc         codec.BinaryMarshaler
}

var _ Configurator = configurator{}

func (c configurator) MsgServer() gogogrpc.Server {
	return c.msgServer
}

func (c configurator) QueryServer() gogogrpc.Server {
	return c.queryServer
}

func (c configurator) ModuleKey() ModuleKey {
	return c.key
}

func (c configurator) BinaryMarshaler() codec.BinaryMarshaler {
	return c.cdc
}
