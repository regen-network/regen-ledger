package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Send sends credits to a recipient.
// Send also retires credits if retirement jurisdiction is specified in the request.
func (k Keeper) Send(ctx context.Context, req *core.MsgSend) (*core.MsgSendResponse, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	sender, _ := sdk.AccAddressFromBech32(req.Sender)
	recipient, _ := sdk.AccAddressFromBech32(req.Recipient)
	credit := &core.MsgSendBulk_SendCredits{
		BatchDenom:     req.BatchDenom,
		TradableAmount: req.Amount,
	}
	err := k.sendEcocredits(ctx, credit, recipient, sender)
	if err != nil {
		return nil, err
	}
	if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventTransfer{
		Sender:         req.Sender,
		Recipient:      req.Recipient,
		BatchDenom:     credit.BatchDenom,
		TradableAmount: credit.TradableAmount,
	}); err != nil {
		return nil, err
	}

	sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/core/MsgSend credit iteration")

	return &core.MsgSendResponse{}, nil
}
