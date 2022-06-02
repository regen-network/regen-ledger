package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/data"
)

// ResolversByIRI queries resolvers based on IRI.
func (s serverImpl) ResolversByIRI(ctx context.Context, request *data.QueryResolversByIRIRequest) (*data.QueryResolversByIRIResponse, error) {
	dataId, err := s.stateStore.DataIDTable().GetByIri(ctx, request.Iri)
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
		api.DataResolverPrimaryKey{}.WithId(dataId.Id),
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

		resolver, err := s.stateStore.ResolverTable().Get(ctx, item.ResolverId)
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

// ResolversByHash queries resolvers based on ContentHash.
func (s serverImpl) ResolversByHash(ctx context.Context, request *data.QueryResolversByHashRequest) (*data.QueryResolversByHashResponse, error) {
	if request.ContentHash == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, err
	}

	dataId, err := s.stateStore.DataIDTable().GetByIri(ctx, iri)
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
		api.DataResolverPrimaryKey{}.WithId(dataId.Id),
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

		resolver, err := s.stateStore.ResolverTable().Get(ctx, item.ResolverId)
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

// ResolversByUrl queries resolvers based on URL.
func (s serverImpl) ResolversByUrl(ctx context.Context, request *data.QueryResolversByUrlRequest) (*data.QueryResolversByUrlResponse, error) {
	if len(request.Url) == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("url cannot be empty")
	}

	if request.Pagination == nil {
		request.Pagination = &query.PageRequest{}
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

	res := &data.QueryResolversByUrlResponse{}
	for it.Next() {
		item, err := it.Value()
		if err != nil {
			return nil, err
		}

		manager := sdk.AccAddress(item.Manager).String()

		res.Resolvers = append(res.Resolvers, &data.ResolverInfo{
			Id:      item.Id,
			Url:     item.Url,
			Manager: manager,
		})
	}

	res.Pagination, err = ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return res, nil
}
