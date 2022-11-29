package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
)

func (k Keeper) Basket(ctx context.Context, request *types.QueryBasketRequest) (*types.QueryBasketResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	basket, err := k.stateStore.BasketTable().GetByBasketDenom(ctx, request.BasketDenom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return nil, errors.Wrapf(err, "basket %s not found", request.BasketDenom)
		}
		return nil, errors.Wrapf(err, "failed to get basket %s", request.BasketDenom)
	}

	basketGogo := &types.Basket{}
	err = ormutil.PulsarToGogoSlow(basket, basketGogo)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.BasketClassTable().List(ctx, api.BasketClassPrimaryKey{}.WithBasketId(basket.Id))
	if err != nil {
		return nil, err
	}

	var classes []string
	for it.Next() {
		class, err := it.Value()
		if err != nil {
			return nil, err
		}

		classes = append(classes, class.ClassId)
	}

	it.Close()

	basketInfo := &types.BasketInfo{
		BasketDenom:       basket.BasketDenom,
		Name:              basket.Name,
		CreditTypeAbbrev:  basket.CreditTypeAbbrev,
		DisableAutoRetire: basket.DisableAutoRetire,
		Exponent:          basket.Exponent, //nolint:staticcheck
		Curator:           sdk.AccAddress(basket.Curator).String(),
	}

	if basket.DateCriteria != nil {
		criteria := &types.DateCriteria{}
		if err := ormutil.PulsarToGogoSlow(basket.DateCriteria, criteria); err != nil {
			return nil, err
		}

		basketInfo.DateCriteria = criteria
	}

	return &types.QueryBasketResponse{Basket: basketGogo, BasketInfo: basketInfo, Classes: classes}, nil
}
