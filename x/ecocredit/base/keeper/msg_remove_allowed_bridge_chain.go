package keeper

import (
	"context"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func (k Keeper) RemoveAllowedBridgeChain(ctx context.Context, req *types.MsgRemoveAllowedBridgeChain) (*types.MsgRemoveAllowedBridgeChainResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	chainName := strings.ToLower(req.ChainName)

	err := k.stateStore.AllowedBridgeChainTable().Delete(ctx, &api.AllowedBridgeChain{ChainName: chainName})
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not delete chain name %s: %s", chainName, err.Error())
	}

	return &types.MsgRemoveAllowedBridgeChainResponse{}, nil
}
