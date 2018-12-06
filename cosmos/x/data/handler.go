package data

import (
"fmt"

sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "data" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgStoreData:
			return handleMsgStoreData(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized data Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgStoreData(ctx sdk.Context, keeper Keeper, msg MsgStoreData) sdk.Result {
	hash := keeper.StoreData(ctx, msg.Data)
	return sdk.Result{
		Data: hash,
	}
}
