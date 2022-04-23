package ecocredit

import (
	context "context"

	"github.com/cosmos/cosmos-sdk/baseapp"

	"github.com/regen-network/regen-ledger/x/axelarbridge"
)

func NewRegenAxelarHandler(router *baseapp.MsgServiceRouter) map[string]axelarbridge.Handler {
	return map[string]axelarbridge.Handler{
		"regen_toucan_bridge": func(ctx context.Context, originChain, senderAddr string, payload []byte) error {
			// TODO:
			// - decode payload (TODO: define payload format)
			// - validation of payload
			// - create a new `regen.ecocredit.v1.Msg.CreateBatch` message
			// - run it via the msg service router
			return nil
		},
	}
}
