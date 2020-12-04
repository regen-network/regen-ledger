package server

import (
	"fmt"
	"reflect"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
	gogogrpc "github.com/gogo/protobuf/grpc"

	"github.com/regen-network/regen-ledger/types/module"
)

type Manager struct {
	sdkmodule.Manager
	baseApp          *baseapp.BaseApp
	cdc              *codec.ProtoCodec
	keys             map[string]ModuleKey
	router           *router
	requiredServices map[reflect.Type]bool
}

func NewManager(baseApp *baseapp.BaseApp, cdc *codec.ProtoCodec) *Manager {
	return &Manager{
		baseApp: baseApp,
		cdc:     cdc,
		keys:    map[string]ModuleKey{},
		router: &router{
			handlers:         map[string]handler{},
			providedServices: map[reflect.Type]bool{},
			antiReentryMap:   map[string]bool{},
		},
	}
}

func (mm *Manager) RegisterModules(modules []module.Module) error {
	for _, mod := range modules {
		// check if we actually have a server module, otherwise skip
		serverMod, ok := mod.(Module)
		if !ok {
			continue
		}

		serverMod.RegisterInterfaces(mm.cdc.InterfaceRegistry())
	}

	for _, mod := range modules {
		// check if we actually have a server module, otherwise skip
		serverMod, ok := mod.(Module)
		if !ok {
			continue
		}

		name := serverMod.Name()

		invokerFactory := mm.router.invokerFactory(name)

		key := &rootModuleKey{
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
			moduleName:   name,
		}

		queryRegistrar := registrar{
			router:       mm.router,
			baseServer:   mm.baseApp.GRPCQueryRouter(),
			commitWrites: true,
			moduleName:   name,
		}

		cfg := &configurator{
			msgServer:        msgRegistrar,
			queryServer:      queryRegistrar,
			key:              key,
			cdc:              mm.cdc,
			requiredServices: map[reflect.Type]bool{},
		}

		serverMod.RegisterServices(cfg)

		for typ := range cfg.requiredServices {
			mm.requiredServices[typ] = true
		}
	}

	return nil
}

type AuthorizationMiddleware func(ctx sdk.Context, methodName string, req sdk.MsgRequest, signer sdk.AccAddress) bool

func (mm *Manager) SetAuthorizationMiddleware(authzFunc AuthorizationMiddleware) {
	mm.router.authzMiddleware = authzFunc
}

func (mm *Manager) CompleteInitialization() error {
	for typ := range mm.requiredServices {
		if _, found := mm.router.providedServices[typ]; !found {
			return fmt.Errorf("initialization error, service %s was required, but not provided", typ)
		}

	}

	return nil
}

type configurator struct {
	msgServer        gogogrpc.Server
	queryServer      gogogrpc.Server
	key              *rootModuleKey
	cdc              codec.BinaryMarshaler
	requiredServices map[reflect.Type]bool
}

var _ Configurator = &configurator{}

func (c *configurator) MsgServer() gogogrpc.Server {
	return c.msgServer
}

func (c *configurator) QueryServer() gogogrpc.Server {
	return c.queryServer
}

func (c *configurator) ModuleKey() RootModuleKey {
	return c.key
}

func (c *configurator) BinaryMarshaler() codec.BinaryMarshaler {
	return c.cdc
}

func (c *configurator) RequireServer(serverInterface interface{}) {
	c.requiredServices[reflect.TypeOf(serverInterface)] = true
}
