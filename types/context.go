package types

import (
	"context"
	"time"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Context struct {
	sdk.Context
}

var _ context.Context = Context{}

func (c Context) Deadline() (deadline time.Time, ok bool) {
	return c.Context.Context().Deadline()
}

func (c Context) Done() <-chan struct{} {
	return c.Context.Context().Done()
}

func (c Context) Err() error {
	return c.Context.Context().Err()
}

func UnwrapSDKContext(ctx context.Context) Context {
	if sdkCtx, ok := ctx.(Context); ok {
		return sdkCtx
	}
	return Context{sdk.UnwrapSDKContext(ctx)}
}

type ContextBase interface {
	context.Context

	BlockHeader() tmproto.Header
	TxBytes() []byte
	TxHash() []byte
	EventManager() sdk.EventManager
}

type SyncContext interface {
	ContextBase

	// GetRef returns a reference to a KV-pair that can be read and written in
	// a callback registered with Exec.
	GetRef(key []byte) KVRef

	// GetRef returns an iterator that can be read and written in
	// a callback registered with Exec.
	GetIteratorRef(key []byte) IteratorRef

	// Exec queues a function to be executed later when state access is safely isolated
	// between transactions running in parallel. Repeated calls to Exec will queue functions
	// to be run sequentially in the order than Exec was called during the prepare phase.
	Exec(func(ExecContext) error)
}

type KVRef interface {
	Value(ExecContext) []byte
	SetValue(ExecContext, []byte)
}

type IteratorRef interface {
	Next(ExecContext) bool
	Key() []byte
	Value() []byte
	SetValue([]byte)
}

type ExecContext interface {
	ContextBase

	QueueSideEffect(func(SyncContext) error)
}

type BeginEndBlockerContext interface {
	ContextBase

	QueueWorkItem(func(SyncContext) error)
}
