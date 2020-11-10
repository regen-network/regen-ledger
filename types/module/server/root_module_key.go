package server

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types/module"
	"google.golang.org/grpc"
)

type RootModuleKey struct {
	moduleName     string
	invokerFactory InvokerFactory
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
	invoker, err := r.invokerFactory(CallInfo{
		Method: method,
		Caller: r.ModuleID(),
	})
	if err != nil {
		return err
	}

	return invoker(ctx, args, reply)
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
		moduleName:     r.moduleName,
		path:           path,
		invokerFactory: r.invokerFactory,
	}
}
