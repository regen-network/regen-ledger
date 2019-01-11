package geo

import (
	"fmt"
	"github.com/twpayne/go-geom/encoding/wkbcommon"
	"gitlab.com/regen-network/regen-ledger/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
	hash := keeper.StoreGeometry(ctx, msg.Data)
	tags := sdk.EmptyTags()
	tags.AppendTag("geo.id", []byte(utils.MustEncodeBech32("xrngeo", hash)))
	return sdk.Result{ Tags: tags }
}


