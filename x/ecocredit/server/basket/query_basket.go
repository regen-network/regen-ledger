package basket

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func (k Keeper) Basket(ctx context.Context, request *baskettypes.QueryBasketRequest) (*baskettypes.QueryBasketResponse, error) {
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

	basketGogo := &baskettypes.Basket{}
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

	basketInfo := &baskettypes.BasketInfo{
		BasketDenom:       basket.BasketDenom,
		Name:              basket.Name,
		CreditTypeAbbrev:  basket.CreditTypeAbbrev,
		DisableAutoRetire: basket.DisableAutoRetire,
		Exponent:          basket.Exponent, //nolint:staticcheck
		Curator:           sdk.AccAddress(basket.Curator).String(),
	}

	if basket.DateCriteria != nil {
		criteria := &baskettypes.DateCriteria{}
		if err := ormutil.PulsarToGogoSlow(basket.DateCriteria, criteria); err != nil {
			return nil, err
		}

		basketInfo.DateCriteria = criteria
	}

	return &baskettypes.QueryBasketResponse{Basket: basketGogo, BasketInfo: basketInfo, Classes: classes}, nil
}
