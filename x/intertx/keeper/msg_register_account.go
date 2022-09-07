package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/regen-network/regen-ledger/x/intertx/types/v1"
)

// RegisterAccount implements the Msg/RegisterAccount interface
func (k Keeper) RegisterAccount(ctx context.Context, msg *types.MsgRegisterAccount) (*types.MsgRegisterAccountResponse, error) {
	if err := k.icaControllerKeeper.RegisterInterchainAccount(sdk.UnwrapSDKContext(ctx), msg.ConnectionId, msg.Owner, msg.Version); err != nil {
		return nil, err
	}
	return &types.MsgRegisterAccountResponse{}, nil
}
