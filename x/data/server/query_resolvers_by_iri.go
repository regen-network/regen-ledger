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

// ResolversByIRI queries resolvers with registered data by the IRI of the data.
func (s serverImpl) ResolversByIRI(ctx context.Context, request *data.QueryResolversByIRIRequest) (*data.QueryResolversByIRIResponse, error) {
	if len(request.Iri) == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("IRI cannot be empty")
	}

	// check for valid IRI
	_, err := data.ParseIRI(request.Iri)
	if err != nil {
		return nil, err
	}

	dataId, err := s.stateStore.DataIDTable().GetByIri(ctx, request.Iri)
	if err != nil {
		return nil, sdkerrors.ErrNotFound.Wrap("data record with IRI")
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
	defer it.Close()

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
