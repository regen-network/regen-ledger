package server

import (
	"context"
	"fmt"
	"github.com/regen-network/regen-ledger/types/module"
	"google.golang.org/grpc"
)

type DerivedModuleKey struct {
	moduleName string
	path       []byte
	invoker    Invoker
}

var _ ModuleKey = DerivedModuleKey{}

func (d DerivedModuleKey) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, V ...grpc.CallOption) error {
	return d.invoker(CallInfo{
		Method: method,
		Caller: d.ModuleID(),
	})(ctx, args, reply)
}

func (d DerivedModuleKey) NewStream(context.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("unsupported")
}

func (d DerivedModuleKey) ModuleID() module.ModuleID {
	return module.ModuleID{
		ModuleName: d.moduleName,
		Path:       d.path,
	}
}

func (d DerivedModuleKey) Address() []byte {
	return d.ModuleID().Address()
}
