package keeper

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MessageRouter interface {
	Handler(msg sdk.Msg) baseapp.MsgServiceHandler
}
