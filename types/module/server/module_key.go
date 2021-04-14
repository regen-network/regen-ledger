package server

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/address"
)

type ModuleKey interface {
	types.InvokerConn

	ModuleID() types.ModuleID
	Address() sdk.AccAddress
	Derive(key []byte) ModuleKey
}

type RootModuleKey interface {
	ModuleKey
	sdk.StoreKey
}

type moduleKey struct {
	moduleName string
	addr       []byte
	i          InvokerFactory
}

func NewDerivedModuleKey(modName string, parentAddr, derivationKey []byte, i InvokerFactory) ModuleKey {
	return moduleKey{modName, address.Derive(parentAddr, derivationKey), i}
}

func (d moduleKey) Invoker(methodName string) (types.Invoker, error) {
	return d.i(CallInfo{
		Method: methodName,
		Caller: d.ModuleID(),
	})
}

func (d moduleKey) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, _ ...grpc.CallOption) error {
	invoker, err := d.Invoker(method)
	if err != nil {
		return err
	}

	return invoker(ctx, args, reply)
}

func (d moduleKey) NewStream(context.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("unsupported")
}

func (d moduleKey) ModuleID() types.ModuleID {
	return types.ModuleID{
		Name:    d.moduleName,
		Address: d.addr,
	}
}

func (d moduleKey) Address() sdk.AccAddress {
	return d.addr
}

func (d moduleKey) Derive(key []byte) ModuleKey {
	return NewDerivedModuleKey(d.moduleName, d.addr, key, d.i)
}

type rootModuleKey struct {
	moduleKey
}

var _ RootModuleKey = rootModuleKey{}

func NewRootModuleKey(name string, i InvokerFactory) RootModuleKey {
	return &rootModuleKey{moduleKey{name, address.Module(name), i}}
}

func (r rootModuleKey) Name() string {
	return r.moduleName
}

func (r rootModuleKey) String() string {
	return fmt.Sprintf("rootModuleKey{%p, %s}", &r, r.moduleName)
}
