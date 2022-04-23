package axelarbridge

import (
	"context"

	bridgev1 "github.com/regen-network/regen-ledger/api/axelar/bridge/v1"
)

// Handler is a function that handlers an arbitray byte array coming from the
// origin chain, from a give senderAddr.
type Handler func(ctx context.Context, event bridgev1.Event) error

// HandlerMap is a container for registered Handlers.
type HandlerMap = map[string]Handler
