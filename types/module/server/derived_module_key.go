package server

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types/module"
	"google.golang.org/grpc"
)

type DerivedModuleKey struct {
	moduleName     string
	path           []byte
	invokerFactory InvokerFactory
}

var _ ModuleKey = DerivedModuleKey{}

func (d DerivedModuleKey) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, V ...grpc.CallOption) error {
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

func (d DerivedModuleKey) ModuleID() module.ModuleID {
	return module.ModuleID{
		ModuleName: d.moduleName,
		Path:       d.path,
	}
}

func (d DerivedModuleKey) Address() sdk.AccAddress {
	return d.ModuleID().Address()
}
