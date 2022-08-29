package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	types "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
)

func (k Keeper) BasketBalance(ctx context.Context, request *types.QueryBasketBalanceRequest) (*types.QueryBasketBalanceResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	basket, err := k.stateStore.BasketTable().GetByBasketDenom(ctx, request.BasketDenom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return nil, sdkerrors.Wrapf(err, "basket %s not found", request.BasketDenom)
		}
		return nil, sdkerrors.Wrapf(err, "failed to get basket %s", request.BasketDenom)
	}

	found, err := k.coreStore.BatchTable().HasByDenom(ctx, request.BatchDenom)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to get credit batch %s", request.BatchDenom)
	}
	if !found {
		return nil, ormerrors.NotFound.Wrapf("credit batch %s not found", request.BatchDenom)
	}

	balance, err := k.stateStore.BasketBalanceTable().Get(ctx, basket.Id, request.BatchDenom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return &types.QueryBasketBalanceResponse{Balance: "0"}, nil
		}
		return nil, sdkerrors.Wrapf(err, "failed to get basket balance for %s", request.BasketDenom)
	}

	return &types.QueryBasketBalanceResponse{Balance: balance.Balance}, nil
}
