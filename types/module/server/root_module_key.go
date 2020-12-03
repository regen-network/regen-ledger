package server

import (
	"context"
	"fmt"

	"github.com/regen-network/regen-ledger/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc"
)

type RootModuleKey interface {
	ModuleKey
	sdk.StoreKey
}

type rootModuleKey struct {
	moduleName     string
	invokerFactory InvokerFactory
}

var _ RootModuleKey = &rootModuleKey{}

func (r *rootModuleKey) Name() string {
	return r.moduleName
}

func (r *rootModuleKey) String() string {
	return fmt.Sprintf("rootModuleKey{%p, %s}", r, r.moduleName)
}

func (r *rootModuleKey) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	invoker, err := r.invokerFactory(CallInfo{
		Method: method,
		Caller: r.ModuleID(),
	})
	if err != nil {
		return err
	}

	return invoker(ctx, args, reply)
}

func (r *rootModuleKey) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("unsupported")
}

func (r *rootModuleKey) ModuleID() types.ModuleID {
	return types.ModuleID{ModuleName: r.moduleName}
}

func (r *rootModuleKey) Address() []byte {
	return r.ModuleID().Address()
}

func (r *rootModuleKey) Derive(path []byte) DerivedModuleKey {
	return DerivedModuleKey{
		moduleName:     r.moduleName,
		path:           path,
		invokerFactory: r.invokerFactory,
	}
}
