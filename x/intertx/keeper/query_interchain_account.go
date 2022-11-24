package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"

	regenerrors "github.com/regen-network/regen-ledger/types/errors"
	types "github.com/regen-network/regen-ledger/x/intertx/types/v1"
)

// InterchainAccount implements the Query/InterchainAccount gRPC method
func (k Keeper) InterchainAccount(goCtx context.Context, req *types.QueryInterchainAccountRequest) (*types.QueryInterchainAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	portID, err := icatypes.NewControllerPortID(req.Owner)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	addr, found := k.icaControllerKeeper.GetInterchainAccountAddress(ctx, req.ConnectionId, portID)
	if !found {
		return nil, regenerrors.ErrNotFound.Wrapf("no account found for portID %s", portID)
	}

	return &types.QueryInterchainAccountResponse{InterchainAccountAddress: addr}, nil
}
