package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"google.golang.org/grpc"

	"github.com/regen-network/regen-ledger/testutil/server"
	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/module"
)

type fixtureFactory struct {
	t       *testing.T
	modules []module.Module
	signers []sdk.AccAddress
}

var _ server.FixtureFactory = fixtureFactory{}

func NewFixtureFactory(t *testing.T, numSigners int, modules []module.Module) server.FixtureFactory {
	signers := makeTestAddresses(numSigners)
	return fixtureFactory{
		t:       t,
		modules: modules,
		signers: signers,
	}
}

func makeTestAddresses(count int) []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, count)
	for i := 0; i < count; i++ {
		_, _, addrs[i] = testdata.KeyTestPubAddr()
	}
	return addrs
}

func (ff fixtureFactory) Setup() server.Fixture {
	registry := types.NewInterfaceRegistry()
	baseApp := baseapp.NewBaseApp("test", log.NewNopLogger(), dbm.NewMemDB(), nil)
	baseApp.MsgServiceRouter().SetInterfaceRegistry(registry)
	baseApp.GRPCQueryRouter().SetInterfaceRegistry(registry)

	cdc := codec.NewProtoCodec(registry)

	mm := NewManager(baseApp, cdc)
	err := mm.RegisterModules(ff.modules)
	require.NoError(ff.t, err)
	err = mm.CompleteInitialization()
	require.NoError(ff.t, err)

	// Setting up bank keeper to use with group module tests
	// TODO: remove once #225 addressed
	banktypes.RegisterInterfaces(registry)
	authtypes.RegisterInterfaces(registry)

	paramsKey := sdk.NewKVStoreKey(paramstypes.StoreKey)
	authKey := sdk.NewKVStoreKey(authtypes.StoreKey)
	bankKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	tkey := sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	amino := codec.NewLegacyAmino()

	authSubspace := paramstypes.NewSubspace(mm.cdc, amino, paramsKey, tkey, authtypes.ModuleName)
	bankSubspace := paramstypes.NewSubspace(mm.cdc, amino, paramsKey, tkey, banktypes.ModuleName)

	accountKeeper := authkeeper.NewAccountKeeper(
		mm.cdc, authKey, authSubspace, authtypes.ProtoBaseAccount, map[string][]string{},
	)
	bankKeeper := bankkeeper.NewBaseKeeper(
		mm.cdc, bankKey, accountKeeper, bankSubspace, map[string]bool{},
	)

	baseApp.Router().AddRoute(sdk.NewRoute(banktypes.ModuleName, bank.NewHandler(bankKeeper)))
	baseApp.MountStore(tkey, sdk.StoreTypeTransient)
	baseApp.MountStore(paramsKey, sdk.StoreTypeIAVL)
	baseApp.MountStore(authKey, sdk.StoreTypeIAVL)
	baseApp.MountStore(bankKey, sdk.StoreTypeIAVL)

	err = baseApp.LoadLatestVersion()
	require.NoError(ff.t, err)

	return fixture{
		baseApp:    baseApp,
		router:     mm.router,
		t:          ff.t,
		signers:    ff.signers,
		bankKeeper: bankKeeper,
	}
}

type fixture struct {
	baseApp    *baseapp.BaseApp
	router     *router
	t          *testing.T
	signers    []sdk.AccAddress
	bankKeeper bankkeeper.BaseKeeper
}

func (f fixture) Context() context.Context {
	return regentypes.Context{Context: f.baseApp.NewUncachedContext(false, tmproto.Header{})}
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

func (f fixture) BankKeeper() bankkeeper.BaseKeeper {
	return f.bankKeeper
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
