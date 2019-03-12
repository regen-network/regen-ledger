package geo

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/twpayne/go-geom/encoding/wkbcommon"
)

// NewHandler returns a handler for "data" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	// TODO set wkbcommon.MaxGeometryElements to a reasonable
	// value to prevent memory attacks
	// See https://github.com/twpayne/go-geom#protection-against-malicious-or-malformed-inputs
	wkbcommon.MaxGeometryElements = [4]int{16384, 16384, 16384, 16384}
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgStoreGeometry:
			return handleMsgStoreGeometry(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized data Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgStoreGeometry(ctx sdk.Context, keeper Keeper, msg MsgStoreGeometry) sdk.Result {
	addr, err := keeper.StoreGeometry(ctx, msg.Data)
	if err != nil {
		return err.Result()
	}
	tags := sdk.EmptyTags()
	tags = tags.AppendTag("geo.id", GeoURL(addr))
	return sdk.Result{Tags: tags}
}
