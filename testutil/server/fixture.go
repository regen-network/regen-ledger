package server

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	gogogrpc "github.com/gogo/protobuf/grpc"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	grpc "google.golang.org/grpc"
	"reflect"
	"testing"
)

type Fixture interface {
	Setup()
	Context() context.Context
	TxConn() grpc.ClientConnInterface
	QueryConn() grpc.ClientConnInterface
	Signers() []sdk.AccAddress
	Teardown()
}

type ConfiguratorFixture struct {
	queryRouter *testRouter
	msgRouter   *testRouter
	keys        []sdk.StoreKey
	t           *testing.T
	signers     []sdk.AccAddress
	ctx         context.Context
}

func (c ConfiguratorFixture) Context() context.Context {
	return c.ctx
}

func (c ConfiguratorFixture) TxConn() grpc.ClientConnInterface {
	return c.msgRouter
}

func (c ConfiguratorFixture) QueryConn() grpc.ClientConnInterface {
	return c.queryRouter
}

func (c ConfiguratorFixture) Signers() []sdk.AccAddress {
	return c.signers
}

func NewConfiguratorFixture(t *testing.T, storeKeys []sdk.StoreKey, signers []sdk.AccAddress) *ConfiguratorFixture {
	return &ConfiguratorFixture{
		queryRouter: newTestRouter(false),
		msgRouter:   newTestRouter(true),
		keys:        storeKeys,
		t:           t,
		signers:     signers,
	}
}

func (c *ConfiguratorFixture) Setup() {
	db := dbm.NewMemDB()

	ms := store.NewCommitMultiStore(db)
	for _, key := range c.keys {
		ms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	}
	err := ms.LoadLatestVersion()
	require.NoError(c.t, err)

	c.ctx = sdk.WrapSDKContext(sdk.NewContext(ms, tmproto.Header{}, false, log.NewNopLogger()))
}

func (c ConfiguratorFixture) Teardown() {}

func (c ConfiguratorFixture) MsgServer() gogogrpc.Server {
	return c.msgRouter
}

func (c ConfiguratorFixture) QueryServer() gogogrpc.Server {
	return c.queryRouter
}

var _ Fixture = &ConfiguratorFixture{}
var _ module.Configurator = ConfiguratorFixture{}

type testRouter struct {
	commitWrites bool
	handlers     map[string]func(ctx context.Context, args, reply interface{}) error
}

func newTestRouter(commitWrites bool) *testRouter {
	return &testRouter{
		commitWrites: commitWrites,
		handlers:     map[string]func(ctx context.Context, args interface{}, reply interface{}) error{},
	}
}

var _ gogogrpc.Server = testRouter{}
var _ grpc.ClientConnInterface = testRouter{}

func (t testRouter) RegisterService(sd *grpc.ServiceDesc, ss interface{}) {
	for _, method := range sd.Methods {
		fqName := fmt.Sprintf("/%s/%s", sd.ServiceName, method.MethodName)
		handler := method.Handler
		t.handlers[fqName] = func(ctx context.Context, args, reply interface{}) error {
			res, err := handler(ss, ctx, func(i interface{}) error { return nil },
				func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, unaryHandler grpc.UnaryHandler) (resp interface{}, err error) {
					return unaryHandler(ctx, args)
				})
			if err != nil {
				return err
			}
			reflect.ValueOf(reply).Elem().Set(reflect.ValueOf(res).Elem())
			return nil
		}
	}
}

func (t testRouter) Invoke(ctx context.Context, method string, args, reply interface{}, opts ...grpc.CallOption) error {
	handler := t.handlers[method]
	if handler == nil {
		return fmt.Errorf("can't find handler for method %s", method)
	}

	// cache wrap the multistore so that writes are batched
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	ms := sdkCtx.MultiStore()
	cacheMs := ms.CacheMultiStore()
	sdkCtx = sdkCtx.WithMultiStore(cacheMs)
	ctx = sdk.WrapSDKContext(sdkCtx)

	err := handler(ctx, args, reply)
	if err != nil {
		return err
	}

	// only commit writes if there are no errors and commitWrites is true
	if t.commitWrites {
		cacheMs.Write()
	}

	return nil
}

func (t testRouter) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("unsupported")
}
