package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/data"
)

// ResolversByURL queries resolvers by URL.
func (s serverImpl) ResolversByURL(ctx context.Context, request *data.QueryResolversByURLRequest) (*data.QueryResolversByURLResponse, error) {
	if len(request.Url) == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("URL cannot be empty")
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := s.stateStore.ResolverTable().List(
		ctx,
		api.ResolverUrlIndexKey{}.WithUrl(request.Url),
		ormlist.Paginate(pg),
	)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	res := &data.QueryResolversByURLResponse{}
	for it.Next() {
		resolver, err := it.Value()
		if err != nil {
			return nil, err
		}

		manager := sdk.AccAddress(resolver.Manager).String()

		res.Resolvers = append(res.Resolvers, &data.ResolverInfo{
			Id:      resolver.Id,
			Url:     resolver.Url,
			Manager: manager,
		})
	}

	res.Pagination, err = ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return res, nil
}
