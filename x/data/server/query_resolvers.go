package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/data"
)

// ResolversByIRI queries resolvers based on IRI.
func (s serverImpl) ResolversByIRI(ctx context.Context, request *data.QueryResolversByIRIRequest) (*data.QueryResolversByIRIResponse, error) {
	id, err := s.getID(ctx, request.Iri)
	if err != nil {
		return nil, err
	}

	if request.Pagination == nil {
		request.Pagination = &query.PageRequest{}
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := s.stateStore.DataResolverTable().List(
		ctx,
		api.DataResolverPrimaryKey{}.WithId(id),
		ormlist.Paginate(pg),
	)
	if err != nil {
		return nil, err
	}

	res := &data.QueryResolversByIRIResponse{}
	for it.Next() {
		item, err := it.Value()
		if err != nil {
			return nil, err
		}

		resolverInfo, err := s.stateStore.ResolverInfoTable().Get(ctx, item.ResolverId)
		if err != nil {
			return nil, err
		}

		res.ResolverUrls = append(res.ResolverUrls, resolverInfo.Url)
	}

	res.Pagination, err = ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ResolversByHash queries resolvers based on ContentHash.
func (s serverImpl) ResolversByHash(ctx context.Context, request *data.QueryResolversByHashRequest) (*data.QueryResolversByHashResponse, error) {
	if request.ContentHash == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, err
	}

	id, err := s.getID(ctx, iri)
	if err != nil {
		return nil, err
	}

	if request.Pagination == nil {
		request.Pagination = &query.PageRequest{}
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := s.stateStore.DataResolverTable().List(
		ctx,
		api.DataResolverPrimaryKey{}.WithId(id),
		ormlist.Paginate(pg),
	)
	if err != nil {
		return nil, err
	}

	res := &data.QueryResolversByHashResponse{}
	for it.Next() {
		item, err := it.Value()
		if err != nil {
			return nil, err
		}

		resolverInfo, err := s.stateStore.ResolverInfoTable().Get(ctx, item.ResolverId)
		if err != nil {
			return nil, err
		}

		res.ResolverUrls = append(res.ResolverUrls, resolverInfo.Url)
	}

	res.Pagination, err = ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return res, nil
}
