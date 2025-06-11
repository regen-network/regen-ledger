package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	"github.com/regen-network/regen-ledger/x/data/v4"
)

// ResolversByHash queries resolvers with registered data by the ContentHash of the data.
func (s serverImpl) ResolversByHash(ctx context.Context, request *data.QueryResolversByHashRequest) (*data.QueryResolversByHashResponse, error) {
	if request.ContentHash == nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap("content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	dataID, err := s.stateStore.DataIDTable().GetByIri(ctx, iri)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrap("data record with content hash")
	}

	it, err := s.stateStore.DataResolverTable().List(
		ctx,
		api.DataResolverPrimaryKey{}.WithId(dataID.Id),
		ormutil.PageReqToOrmPaginate(request.Pagination),
	)
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}
	defer it.Close()

	res := &data.QueryResolversByHashResponse{}
	for it.Next() {
		item, err := it.Value()
		if err != nil {
			return nil, err
		}

		resolver, err := s.stateStore.ResolverTable().Get(ctx, item.ResolverId)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("failed to get resolver by ID: %d", item.ResolverId)
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
