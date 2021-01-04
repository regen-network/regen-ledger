package types

import (
	"context"
	"time"

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
