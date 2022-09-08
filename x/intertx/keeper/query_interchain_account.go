package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"

	types "github.com/regen-network/regen-ledger/x/intertx/types/v1"
)

// InterchainAccount implements the Query/InterchainAccount gRPC method
func (k Keeper) InterchainAccount(goCtx context.Context, req *types.QueryInterchainAccountRequest) (*types.QueryInterchainAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	portID, err := icatypes.NewControllerPortID(req.Owner)
	if err != nil {
		return nil, err
	}

	addr, found := k.icaControllerKeeper.GetInterchainAccountAddress(ctx, req.ConnectionId, portID)
	if !found {
		return nil, sdkerrors.ErrNotFound.Wrapf("no account found for portID %s", portID)
	}

	return &types.QueryInterchainAccountResponse{InterchainAccountAddress: addr}, nil
}
