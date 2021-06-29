package server

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"google.golang.org/grpc"

	"github.com/regen-network/regen-ledger/types"
)

// RootModuleKey is a master key for modules to derive module accounts. It can be used as
// a store key. It doesn't have address - only ModuleKey is addressable. It can't be used to
// instantiate gRPC clients - derive a module key to use it with clients.
type RootModuleKey interface {
	sdk.StoreKey

	Derive(key []byte) ModuleKey
}

// ModuleKey is an interface for module servers required by router. It's used to reference
// a moudle account, derive sub accounts and as a connection interface for gRPC clients.
type ModuleKey interface {
	types.InvokerConn

	ModuleAcc() types.ModuleAcc
	Address() sdk.AccAddress
	Derive(key []byte) ModuleKey
}

type moduleKey struct {
	moduleName string
	addr       []byte
	key        []byte
	i          InvokerFactory
}

// Invoker implements ModuleKey interface
func (d moduleKey) Invoker(methodName string) (types.Invoker, error) {
	return d.i(CallInfo{
		Method: methodName,
		Caller: d.ModuleAcc(),
	})
}

// Invoke implements ModuleKey interface
func (d moduleKey) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, _ ...grpc.CallOption) error {
	invoker, err := d.Invoker(method)
	if err != nil {
		return err
	}

	return invoker(ctx, args, reply)
}

// NewStream implements ModuleKey interface
func (d moduleKey) NewStream(context.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("unsupported")
}

// ModuleAcc implements ModuleKey interface
func (d moduleKey) ModuleAcc() types.ModuleAcc {
	return types.ModuleAcc{
		Module:  d.moduleName,
		Key:     d.key,
		Address: d.addr,
	}
}

// Address implements ModuleKey interface
func (d moduleKey) Address() sdk.AccAddress {
	return d.addr
}

// Derive implements ModuleKey interface
func (d moduleKey) Derive(key []byte) ModuleKey {
	return &moduleKey{d.moduleName, address.Derive(d.addr, key), key, d.i}
}

type rootModuleKey struct {
	moduleName string
	i          InvokerFactory
}

var _ RootModuleKey = rootModuleKey{}

func NewRootModuleKey(name string, i InvokerFactory) RootModuleKey {
	return &rootModuleKey{name, i}
}

// Name implements sdk.StoreKey interface
func (r rootModuleKey) Name() string {
	return r.moduleName
}

// String implements sdk.StoreKey interface
func (r rootModuleKey) String() string {
	return fmt.Sprintf("rootModuleKey{%p, %s}", &r, r.moduleName)
}

// Derive implements RootModuleKey interface
func (r rootModuleKey) Derive(key []byte) ModuleKey {
	return &moduleKey{r.moduleName, address.Module(r.moduleName, key), key, r.i}
}
