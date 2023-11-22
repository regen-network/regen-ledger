package fixture

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"
	"google.golang.org/grpc"

	sdkmodules "github.com/cosmos/cosmos-sdk/types/module"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ Factory = &factoryImpl{}

type factoryImpl struct {
	t       gocuke.TestingT
	modules []sdkmodules.AppModule
	signers []sdk.AccAddress
	cdc     *codec.ProtoCodec
	baseApp *baseapp.BaseApp
}

func NewFixtureFactory(t gocuke.TestingT, numSigners int) Factory {
	signers := makeTestAddresses(numSigners)
	return &factoryImpl{
		t:       t,
		signers: signers,
		// cdc and baseApp are initialized here just for compatibility with legacy modules which don't use ADR 033
		// TODO: remove once all code using this uses ADR 033 module wiring
		cdc:     codec.NewProtoCodec(types.NewInterfaceRegistry()),
		baseApp: baseapp.NewBaseApp("test", log.NewNopLogger(), dbm.NewMemDB(), nil),
	}
}

func (ff *factoryImpl) SetModules(modules []sdkmodules.AppModule) {
	ff.modules = modules
	// we append the mock module below in order to bypass the check for validator updates.
	// since we are testing with a fixture with no validators, we must inject a mock module and
	// force it to inject a validator update.
	ff.modules = append(ff.modules, MockModule{})
}

// Codec is exposed just for compatibility of these test suites with legacy modules and can be removed when everything
// has been migrated to ADR 033
func (ff *factoryImpl) Codec() *codec.ProtoCodec {
	return ff.cdc
}

// BaseApp is exposed just for compatibility of these test suites with legacy modules and can be removed when everything
// has been migrated to ADR 033
func (ff *factoryImpl) BaseApp() *baseapp.BaseApp {
	return ff.baseApp
}

func makeTestAddresses(count int) []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, count)
	for i := 0; i < count; i++ {
		// generate from secret so that keys are deterministic
		key := secp256k1.GenPrivKeyFromSecret([]byte{byte(i)})
		addrs[i] = sdk.AccAddress(key.PubKey().Address())
	}
	return addrs
}

func (ff factoryImpl) Setup() Fixture {
	cdc := ff.cdc
	registry := cdc.InterfaceRegistry()
	baseApp := ff.baseApp
	baseApp.MsgServiceRouter().SetInterfaceRegistry(registry)
	baseApp.GRPCQueryRouter().SetInterfaceRegistry(registry)
	mm := sdkmodules.NewManager(ff.modules...)
	cfg := sdkmodules.NewConfigurator(cdc, baseApp.MsgServiceRouter(), baseApp.GRPCQueryRouter())
	for _, module := range mm.Modules {
		module.RegisterInterfaces(ff.cdc.InterfaceRegistry())
		module.RegisterServices(cfg)
	}

	err := baseApp.LoadLatestVersion()
	require.NoError(ff.t, err)

	return fixture{
		baseApp: baseApp,
		mm:      mm,
		cdc:     cdc,
		router: &router{
			cdc:                cdc.GRPCCodec(),
			msgServiceRouter:   baseApp.MsgServiceRouter(),
			queryServiceRouter: baseApp.GRPCQueryRouter(),
		},
		t:       ff.t,
		signers: ff.signers,
	}
}

type fixture struct {
	baseApp *baseapp.BaseApp
	mm      *sdkmodules.Manager
	router  *router
	cdc     *codec.ProtoCodec
	t       gocuke.TestingT
	signers []sdk.AccAddress
}

func (f fixture) Context() context.Context {
	return f.baseApp.NewUncachedContext(false, tmproto.Header{})
}

func (f fixture) TxConn() grpc.ClientConnInterface {
	return testKey{invokerFactory: f.router.testTxFactory(f.signers)}
}

func (f fixture) QueryConn() grpc.ClientConnInterface {
	return testKey{invokerFactory: f.router.testQueryFactory()}
}

func (f fixture) Signers() []sdk.AccAddress {
	return f.signers
}

func (f fixture) InitGenesis(ctx sdk.Context, genesisData map[string]json.RawMessage) (abci.ResponseInitChain, error) {
	// we inject the mock module genesis in order to bypass the check for validator updates.
	// since the testing fixture doesn't require validators/validator updates, the check fails otherwise.
	genesisData[MockModule{}.Name()] = []byte(`{}`)
	return f.mm.InitGenesis(ctx, f.cdc, genesisData), nil
}

func (f fixture) ExportGenesis(ctx sdk.Context) (map[string]json.RawMessage, error) {
	return f.mm.ExportGenesis(ctx, f.cdc), nil
}

func (f fixture) Codec() *codec.ProtoCodec {
	return f.cdc
}

func (f fixture) Teardown() {}

type testKey struct {
	invokerFactory InvokerFactory
}

var _ grpc.ClientConnInterface = testKey{}

func (t testKey) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, _ ...grpc.CallOption) error {
	invoker, err := t.invokerFactory(CallInfo{Method: method})
	if err != nil {
		return err
	}

	return invoker(ctx, args, reply)
}

func (t testKey) NewStream(context.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("unsupported")
}
