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
	err = baseApp.LoadLatestVersion()
	require.NoError(ff.t, err)

	return fixture{
		baseApp: baseApp,
		router:  mm.router,
		t:       ff.t,
		signers: ff.signers,
	}
}

type fixture struct {
	baseApp *baseapp.BaseApp
	router  *router
	t       *testing.T
	signers []sdk.AccAddress
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
