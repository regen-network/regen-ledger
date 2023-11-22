/*
Package fixture defines fixture interfaces and implementations for testing
server implementations with multiple backends.

Currently one backend - an in-memory store with no ABCI application is supported
in configuration.Fixture.

A multi-node in-process ABCI-based backend for full integration tests is planned
based on to the Cosmos SDK in-process integration test framework.
*/
package fixture

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmodules "github.com/cosmos/cosmos-sdk/types/module"
)

// Factory defines an interface for creating server test fixtures
type Factory interface {

	// Setup runs necessary fixture setup and returns a fresh Fixture environment.
	Setup() Fixture

	// BaseApp returns the baseApp used in the test fixture..
	BaseApp() *baseapp.BaseApp

	// Codec returns the codec.
	Codec() *codec.ProtoCodec

	// SetModules sets the modules to be used in the test fixture.
	SetModules(modules []sdkmodules.AppModule)
}

// Fixture defines an interface for interacting with app services in tests
// independent of the backend.
type Fixture interface {

	// Context is the context.Context to be used with gRPC generated client code.
	Context() context.Context

	// TxConn is the grpc.ClientConnInterface to be used when constructing Msg service clients.
	TxConn() grpc.ClientConnInterface

	// QueryConn is the grpc.ClientConnInterface to be used when constructing Query service clients.
	QueryConn() grpc.ClientConnInterface

	// Signers are a list of addresses which can be used to sign transactions. They may either be
	// random or correspond to nodes in a test network which have keyrings.
	Signers() []sdk.AccAddress

	// InitGenesis initializes genesis for all modules with provided genesisData.
	InitGenesis(ctx sdk.Context, genesisData map[string]json.RawMessage) (abci.ResponseInitChain, error)

	// ExportGenesis returns raw encoded JSON genesis state for all modules.
	ExportGenesis(ctx sdk.Context) (map[string]json.RawMessage, error)

	// Codec is the app ProtoCodec.
	Codec() *codec.ProtoCodec

	// Teardown performs any teardown actions for the fixture.
	Teardown()
}
