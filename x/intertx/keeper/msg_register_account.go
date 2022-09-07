package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/regen-network/regen-ledger/x/intertx/types/v1"
)

// RegisterAccount implements the Msg/RegisterAccount interface
func (k Keeper) RegisterAccount(goCtx context.Context, msg *types.MsgRegisterAccount) (*types.MsgRegisterAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.icaControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, msg.Owner, msg.Version); err != nil {
		return nil, err
	}

	return &types.MsgRegisterAccountResponse{}, nil
}
