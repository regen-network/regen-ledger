package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	baseapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/errors"
	"github.com/regen-network/regen-ledger/types/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// AllowedClassCreators queries list of allowed class creators.
func (k Keeper) AllowedClassCreators(ctx context.Context, req *types.QueryAllowedClassCreatorsRequest) (*types.QueryAllowedClassCreatorsResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	itr, err := k.stateStore.AllowedClassCreatorTable().List(ctx, baseapi.AllowedClassCreatorAddressIndexKey{}, ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer itr.Close()

	classCreators := make([]string, 0, 8) // pre-allocate some cap space
	for itr.Next() {
		val, err := itr.Value()
		if err != nil {
			return nil, err
		}

		classCreators = append(classCreators, sdk.AccAddress(val.Address).String())
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(itr.PageResponse())
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}

	return &types.QueryAllowedClassCreatorsResponse{
		ClassCreators: classCreators,
		Pagination:    pr,
	}, nil
}
