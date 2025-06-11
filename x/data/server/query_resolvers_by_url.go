package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	"github.com/regen-network/regen-ledger/x/data/v4"
)

// ResolversByURL queries resolvers by URL.
func (s serverImpl) ResolversByURL(ctx context.Context, request *data.QueryResolversByURLRequest) (*data.QueryResolversByURLResponse, error) {
	if len(request.Url) == 0 {
		return nil, regenerrors.ErrInvalidArgument.Wrap("URL cannot be empty")
	}

	it, err := s.stateStore.ResolverTable().List(
		ctx,
		api.ResolverUrlIndexKey{}.WithUrl(request.Url),
		ormutil.PageReqToOrmPaginate(request.Pagination),
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
	res.Pagination = ormutil.PageResToCosmosTypes(it.PageResponse())

	return res, nil
}
