package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

// ClassIssuers returns a list of addresses that are allowed to issue batches from the given class.
func (k Keeper) ClassIssuers(ctx context.Context, req *types.QueryClassIssuersRequest) (*types.QueryClassIssuersResponse, error) {
	classInfo, err := k.stateStore.ClassTable().GetById(ctx, req.ClassId)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("could not get class with id %s: %s", req.ClassId, err.Error())
	}

	pg := ormutil.PageReqToOrmPaginate(req.Pagination)
	it, err := k.stateStore.ClassIssuerTable().List(
		ctx, api.ClassIssuerClassKeyIssuerIndexKey{}.WithClassKey(classInfo.Key), pg)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	issuers := make([]string, 0)
	for it.Next() {
		val, err := it.Value()
		if err != nil {
			return nil, err
		}
		issuers = append(issuers, sdk.AccAddress(val.Issuer).String())
	}
	return &types.QueryClassIssuersResponse{
		Issuers:    issuers,
		Pagination: ormutil.PageResToCosmosTypes(it.PageResponse()),
	}, nil
}
