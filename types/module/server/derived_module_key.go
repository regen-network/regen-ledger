package server

import (
	"context"
	"fmt"

	"github.com/regen-network/regen-ledger/types"

	"google.golang.org/grpc"
)

type DerivedModuleKey struct {
	moduleName     string
	path           []byte
	invokerFactory InvokerFactory
}

var _ ModuleKey = DerivedModuleKey{}

func (d DerivedModuleKey) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, _ ...grpc.CallOption) error {
	invoker, err := d.invokerFactory(CallInfo{
		Method: method,
		Caller: d.ModuleID(),
	})

	if err != nil {
		return err
	}

	return invoker(ctx, args, reply)
}

func (d DerivedModuleKey) NewStream(context.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("unsupported")
}

func (d DerivedModuleKey) ModuleID() types.ModuleID {
	return types.ModuleID{
		ModuleName: d.moduleName,
		Path:       d.path,
	}
}

func (d DerivedModuleKey) Address() []byte {
	return d.ModuleID().Address()
}
