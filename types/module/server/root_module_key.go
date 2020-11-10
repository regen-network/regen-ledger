package server

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types/module"
	"google.golang.org/grpc"
)

type RootModuleKey struct {
	moduleName string
	invoker    Invoker
}

var _ ModuleKey = RootModuleKey{}
var _ sdk.StoreKey = RootModuleKey{}

func (r RootModuleKey) Name() string {
	return r.moduleName
}

func (r RootModuleKey) String() string {
	return fmt.Sprintf("RootModuleKey{%p, %s}", r, r.moduleName)
}

func (r RootModuleKey) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	return r.invoker(CallInfo{
		Method: method,
		Caller: r.ModuleID(),
	})(ctx, args, reply)
}

func (r RootModuleKey) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("unsupported")
}

func (r RootModuleKey) ModuleID() module.ModuleID {
	return module.ModuleID{ModuleName: r.moduleName}
}

func (r RootModuleKey) Address() []byte {
	return r.ModuleID().Address()
}

func (r RootModuleKey) Derive(path []byte) DerivedModuleKey {
	return DerivedModuleKey{
		moduleName: r.moduleName,
		path:       path,
		invoker:    r.invoker,
	}
}
