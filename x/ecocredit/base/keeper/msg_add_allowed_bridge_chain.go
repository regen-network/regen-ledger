package keeper

import (
	"context"
	"strings"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

func (k Keeper) AddAllowedBridgeChain(ctx context.Context, req *types.MsgAddAllowedBridgeChain) (*types.MsgAddAllowedBridgeChainResponse, error) {
	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	authorityBz, err := k.ac.StringToBytes(req.Authority)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid authority address")
	}

	authority := sdk.AccAddress(authorityBz)
	if !authority.Equals(k.authority) {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	chainName := strings.ToLower(req.ChainName)

	err = k.stateStore.AllowedBridgeChainTable().Insert(ctx, &api.AllowedBridgeChain{ChainName: chainName})
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not insert chain name %s: %s", req.ChainName, err.Error())
	}

	return &types.MsgAddAllowedBridgeChainResponse{}, nil
}
