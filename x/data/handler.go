package data

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "data" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgStoreGraph:
			return handleMsgStoreData(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized data Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgStoreData(ctx sdk.Context, keeper Keeper, msg MsgStoreGraph) sdk.Result {
	hash, err := keeper.StoreGraph(ctx, msg.Hash, msg.Data)
	if err != nil {
		return err.Result()
	}
	tags := sdk.EmptyTags()
	tags.AppendTag("data.hash", hex.EncodeToString(hash))
	return sdk.Result{
		Data: hash,
		Tags: tags,
	}
}
