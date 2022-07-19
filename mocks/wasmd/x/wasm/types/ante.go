package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type contextKey int

const (
	contextKeyTXCount contextKey = iota
)

func WithTXCounter(ctx sdk.Context, counter uint32) sdk.Context {
	return ctx.WithValue(contextKeyTXCount, counter)
}
