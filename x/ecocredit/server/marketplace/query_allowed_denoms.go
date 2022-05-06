package marketplace

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"

	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func (k Keeper) AllowedDenoms(ctx context.Context, req *marketplace.QueryAllowedDenomsRequest) (*marketplace.QueryAllowedDenomsResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.AllowedDenomTable().List(ctx, &marketplacev1.AllowedDenomPrimaryKey{}, ormlist.Paginate(pg))
	defer it.Close()

	allowedDenoms := make([]*marketplace.AllowedDenom, 0)
	for it.Next() {
		allowedDenom, err := it.Value()
		if err != nil {
			return nil, err
		}
		var ad marketplace.AllowedDenom
		if err = ormutil.PulsarToGogoSlow(allowedDenom, &ad); err != nil {
			return nil, err
		}
		allowedDenoms = append(allowedDenoms, &ad)
	}
	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &marketplace.QueryAllowedDenomsResponse{AllowedDenoms: allowedDenoms, Pagination: pr}, nil
}
