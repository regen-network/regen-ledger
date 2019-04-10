package claim

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "claim" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSignClaim:
			return handleMsgSignClaim(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized data Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgSignClaim(ctx sdk.Context, keeper Keeper, msg MsgSignClaim) sdk.Result {
	err := keeper.SignClaim(ctx, msg.Content, msg.Evidence, msg.Signers)
	if err != nil {
		return err.Result()
	}
	res := sdk.Result{}
	res.Tags = res.Tags.AppendTag("claim.id", msg.Content.String())
	return res
}
