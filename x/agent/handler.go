package agent

import (
"fmt"

sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "data" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreateAgent:
			return handleMsgCreateAgent(ctx, keeper, msg)
		case MsgUpdateAgent:
			return handleMsgUpdateAgent(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized data Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}


func handleMsgCreateAgent(ctx sdk.Context, keeper Keeper, msg MsgCreateAgent) sdk.Result {
	id := keeper.CreateAgent(ctx, msg.Data)
	return sdk.Result{
		Tags: sdk.NewTags("agent.id", []byte(MustEncodeBech32AgentID(id))),
	}
}

func handleMsgUpdateAgent(ctx sdk.Context, keeper Keeper, msg MsgUpdateAgent) sdk.Result {
	keeper.UpdateAgentInfo(ctx, msg.Id, msg.Signers, msg.Data)
	return sdk.Result{
		Tags: sdk.NewTags("agent.id", []byte(MustEncodeBech32AgentID(msg.Id))),
	}
}
