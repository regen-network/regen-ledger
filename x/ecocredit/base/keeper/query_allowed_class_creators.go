package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	baseapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

// AllowedClassCreators queries list of allowed class creators.
func (k Keeper) AllowedClassCreators(ctx context.Context, req *types.QueryAllowedClassCreatorsRequest) (*types.QueryAllowedClassCreatorsResponse, error) {
	if req == nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap("empty request")
	}
	pg := ormutil.PageReqToOrmPaginate(req.Pagination)
	it, err := k.stateStore.AllowedClassCreatorTable().List(ctx, baseapi.AllowedClassCreatorAddressIndexKey{}, pg)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	classCreators := make([]string, 0, 8) // pre-allocate some cap space
	for it.Next() {
		val, err := it.Value()
		if err != nil {
			return nil, err
		}
		classCreators = append(classCreators, sdk.AccAddress(val.Address).String())
	}

	return &types.QueryAllowedClassCreatorsResponse{
		ClassCreators: classCreators,
		Pagination:    ormutil.PageResToCosmosTypes(it.PageResponse()),
	}, nil
}
