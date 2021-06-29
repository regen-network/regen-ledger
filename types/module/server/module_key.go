package server

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"google.golang.org/grpc"

	"github.com/regen-network/regen-ledger/types"
)

// ModuleKey is an interface for module servers required by router.
type ModuleKey interface {
	types.InvokerConn

	ModuleID() types.ModuleID
	Address() sdk.AccAddress
	Derive(key []byte) ModuleKey
}

// RootModuleKey is a master key for modules to derive other keys. It doesn't have address -
// only ModuleKey is addressable.
type RootModuleKey interface {
	types.InvokerConn
	ModuleID() types.ModuleID
	Derive(key []byte) ModuleKey

	sdk.StoreKey
}

type moduleKey struct {
	moduleName string
	addr       []byte
	key        []byte
	i          InvokerFactory
}

// NewDerivedModuleKey creates a ModuleKey with a derived module address based on parent
// module address and derivation key.
func NewDerivedModuleKey(modName string, parentAddr, derivationKey []byte, i InvokerFactory) ModuleKey {
	return moduleKey{modName, address.Derive(parentAddr, derivationKey), derivationKey, i}
}

// Invoker implements ModuleKey interface
func (d moduleKey) Invoker(methodName string) (types.Invoker, error) {
	return d.i(CallInfo{
		Method: methodName,
		Caller: d.ModuleID(),
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

// ModuleID implements ModuleKey interface
func (d moduleKey) ModuleID() types.ModuleID {
	return types.ModuleID{
		Name: d.moduleName,
		Key:  d.key,
	}
}

// Address implements ModuleKey interface
func (d moduleKey) Address() sdk.AccAddress {
	return d.addr
}

// Derive implements ModuleKey interface
func (d moduleKey) Derive(key []byte) ModuleKey {
	return NewDerivedModuleKey(d.moduleName, d.addr, key, d.i)
}

type rootModuleKey struct {
	moduleKey
}

var _ RootModuleKey = rootModuleKey{}

func NewRootModuleKey(name string, i InvokerFactory) RootModuleKey {
	// return &rootModuleKey{moduleKey{name, address.Module(name), i}}  // TODO
	return &rootModuleKey{moduleKey{name, []byte(name), i}}
}

// Name implements sdk.StoreKey interface
func (r rootModuleKey) Name() string {
	return r.moduleName
}

// String implements sdk.StoreKey interface
func (r rootModuleKey) String() string {
	return fmt.Sprintf("rootModuleKey{%p, %s}", &r, r.moduleName)
}
