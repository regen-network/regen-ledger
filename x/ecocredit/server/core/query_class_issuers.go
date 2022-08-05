package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// ClassIssuers returns a list of addresses that are allowed to issue batches from the given class.
func (k Keeper) ClassIssuers(ctx context.Context, request *core.QueryClassIssuersRequest) (*core.QueryClassIssuersResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	classInfo, err := k.stateStore.ClassTable().GetById(ctx, request.ClassId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get class with id %s: %s", request.ClassId, err.Error())
	}

	it, err := k.stateStore.ClassIssuerTable().List(ctx, api.ClassIssuerClassKeyIssuerIndexKey{}.WithClassKey(classInfo.Key), ormlist.Paginate(pg))
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
	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &core.QueryClassIssuersResponse{
		Issuers:    issuers,
		Pagination: pr,
	}, nil
}
