/*Package server defines fixture interfaces and implementations for testing
server implementations with multiple backends.

Currently one backend - an in-memory store with no ABCI application is supported
in configuration.Fixture.

A multi-node in-process ABCI-based backend for full integration tests is planned
based on to the Cosmos SDK in-process integration test framework.
*/
package testutil

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc"
)

// FixtureFactory defines an interface for creating server test fixtures
type FixtureFactory interface {

	// Setup runs necessary fixture setup and returns a fresh Fixture environment.
	Setup() Fixture
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

	// Teardown performs any teardown actions for the fixture.
	Teardown()
}
