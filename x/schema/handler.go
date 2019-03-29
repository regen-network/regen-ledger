package schema

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "schema" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case PropertyDefinition:
			return handlePropertyDefinition(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized data Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handlePropertyDefinition(ctx sdk.Context, keeper Keeper, def PropertyDefinition) sdk.Result {
	id, uri, err := keeper.DefineProperty(ctx, def)
	if err != nil {
		return err.Result()
	}
	res := sdk.Result{}
	res.Tags = res.Tags.AppendTag("property.url", uri.String()).AppendTag("property.id", fmt.Sprintf("%d", id))
	return res
}
