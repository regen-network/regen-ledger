package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	"github.com/regen-network/regen-ledger/x/data/v3"
)

// ResolversByIRI queries resolvers with registered data by the IRI of the data.
func (s serverImpl) ResolversByIRI(ctx context.Context, request *data.QueryResolversByIRIRequest) (*data.QueryResolversByIRIResponse, error) {
	if len(request.Iri) == 0 {
		return nil, regenerrors.ErrInvalidArgument.Wrap("IRI cannot be empty")
	}

	// check for valid IRI
	_, err := data.ParseIRI(request.Iri)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	dataID, err := s.stateStore.DataIDTable().GetByIri(ctx, request.Iri)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("data record with IRI: %s", request.Iri)
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

	res := &data.QueryResolversByIRIResponse{}
	for it.Next() {
		item, err := it.Value()
		if err != nil {
			return nil, err
		}

		resolver, err := s.stateStore.ResolverTable().Get(ctx, item.ResolverId)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrapf("resolver with ID: %d", item.ResolverId)
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
