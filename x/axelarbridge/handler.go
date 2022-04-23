package axelarbridge

import (
	"context"
)

// Handler is a function that handlers an arbitray byte array coming from the
// origin chain, from a give senderAddr.
type Handler func(ctx context.Context, originChain, senderAddr string, payload []byte) error

// HandlerMap is a container for registered Handlers.
type HandlerMap = map[string]Handler
