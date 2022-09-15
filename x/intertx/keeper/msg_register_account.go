package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/regen-network/regen-ledger/x/intertx/types/v1"
)

// RegisterAccount implements the Msg/RegisterAccount interface
func (k Keeper) RegisterAccount(ctx context.Context, msg *types.MsgRegisterAccount) (*types.MsgRegisterAccountResponse, error) {
	if err := k.icaControllerKeeper.RegisterInterchainAccount(sdk.UnwrapSDKContext(ctx), msg.ConnectionId, msg.Owner, msg.Version); err != nil {
		realErr := fmt.Errorf("ICAControllerKeeper.RegisterInterchainAccount: %w", err)
		return nil, realErr
	}
	return &types.MsgRegisterAccountResponse{}, nil
}
