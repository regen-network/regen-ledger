package esp

import (
"fmt"

sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgRegisterESPVersion:
			return handleMsgRegisterESP(ctx, keeper, msg)
		case MsgReportESPResult:
			return handleMsgReportESPResult(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized data Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgRegisterESP(ctx sdk.Context, keeper Keeper, msg MsgRegisterESPVersion) sdk.Result {
	code := keeper.RegisterESPVersion(ctx, msg.Curator, msg.Name, msg.Version, msg.Spec, msg.Signers)
	return sdk.Result{ Code:code }
}

func handleMsgReportESPResult(ctx sdk.Context, keeper Keeper, msg MsgReportESPResult) sdk.Result {
	code := keeper.ReportESPResult(ctx, msg.Curator, msg.Name, msg.Version, msg.Verifier, msg.Result, msg.Signers)
	return sdk.Result{ Code:code }
}

