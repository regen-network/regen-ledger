package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/data"
)

// ResolversByHash queries resolvers with registered data by the ContentHash of the data.
func (s serverImpl) ResolversByHash(ctx context.Context, request *data.QueryResolversByHashRequest) (*data.QueryResolversByHashResponse, error) {
	if request.ContentHash == nil {
		return nil, status.Error(codes.InvalidArgument, "content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, err
	}

	dataID, err := s.stateStore.DataIDTable().GetByIri(ctx, iri)
	if err != nil {
		return nil, status.Error(codes.NotFound, "data record with content hash")
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := s.stateStore.DataResolverTable().List(
		ctx,
		api.DataResolverPrimaryKey{}.WithId(dataID.Id),
		ormlist.Paginate(pg),
	)
	if err != nil {
		return nil, err
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
			return nil, status.Errorf(codes.NotFound, "failed to get resolver: %d", item.ResolverId)
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
