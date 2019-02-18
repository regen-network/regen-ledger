package group

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "data" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreateGroup:
			return handleMsgCreateGroup(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized data Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgCreateGroup(ctx sdk.Context, keeper Keeper, msg MsgCreateGroup) sdk.Result {
	id := keeper.CreateGroup(ctx, msg.Data)
	return sdk.Result{
		Tags: sdk.NewTags("group.id", []byte(id.String())),
	}
}
