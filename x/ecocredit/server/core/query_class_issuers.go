package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
)

// ClassIssuers returns a list of addresses that are allowed to issue batches from the given class.
func (k Keeper) ClassIssuers(ctx context.Context, request *v1beta1.QueryClassIssuersRequest) (*v1beta1.QueryClassIssuersResponse, error) {
	pg, err := GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	classInfo, err := k.stateStore.ClassInfoStore().GetByName(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.ClassIssuerStore().List(ctx, ecocreditv1beta1.ClassIssuerClassIdIssuerIndexKey{}.WithClassId(classInfo.Id), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	issuers := make([]string, 0)
	for it.Next() {
		val, err := it.Value()
		if err != nil {
			return nil, err
		}
		issuers = append(issuers, sdk.AccAddress(val.Issuer).String())
	}
	pr, err := PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &v1beta1.QueryClassIssuersResponse{
		Issuers:    issuers,
		Pagination: pr,
	}, nil
}
