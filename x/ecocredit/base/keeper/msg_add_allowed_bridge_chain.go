package keeper

import (
	"context"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func (k Keeper) AddAllowedBridgeChain(ctx context.Context, req *types.MsgAddAllowedBridgeChain) (*types.MsgAddAllowedBridgeChainResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	chainName := strings.ToLower(req.ChainName)

	err := k.stateStore.AllowedBridgeChainTable().Insert(ctx, &api.AllowedBridgeChain{ChainName: chainName})
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not insert chain name %s: %s", req.ChainName, err.Error())
	}

	return &types.MsgAddAllowedBridgeChainResponse{}, nil
}
