package keeper

import (
	"context"

	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/regen-network/regen-ledger/x/intertx/types/v1"
)

// RegisterAccount implements the Msg/RegisterAccount interface
func (k Keeper) RegisterAccount(ctx context.Context, msg *types.MsgRegisterAccount) (*types.MsgRegisterAccountResponse, error) {
	if err := k.icaControllerKeeper.RegisterInterchainAccount(sdk.UnwrapSDKContext(ctx), msg.ConnectionId, msg.Owner, msg.Version); err != nil {
		wrappedErr := errors.Wrap(err, "error in ICAControllerKeeper.RegisterInterchainAccount")
		return nil, wrappedErr
	}
	return &types.MsgRegisterAccountResponse{}, nil
}
