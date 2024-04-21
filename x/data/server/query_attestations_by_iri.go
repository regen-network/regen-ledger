package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	"github.com/regen-network/regen-ledger/types/v2"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	"github.com/regen-network/regen-ledger/x/data/v3"
)

// AttestationsByIRI queries data attestations by the IRI of the data.
func (s serverImpl) AttestationsByIRI(ctx context.Context, request *data.QueryAttestationsByIRIRequest) (*data.QueryAttestationsByIRIResponse, error) {
	if len(request.Iri) == 0 {
		return nil, regenerrors.ErrInvalidArgument.Wrap("IRI cannot be empty")
	}

	// check for valid IRI
	_, err := data.ParseIRI(request.Iri)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrapf("failed to parse IRI: %s", err.Error())
	}

	dataID, err := s.stateStore.DataIDTable().GetByIri(ctx, request.Iri)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("data record with IRI: %s", request.Iri)
	}

	it, err := s.stateStore.DataAttestorTable().List(
		ctx,
		api.DataAttestorIdAttestorIndexKey{}.WithId(dataID.Id),
		ormutil.PageReqToOrmPaginate(request.Pagination),
	)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	var attestations []*data.AttestationInfo
	for it.Next() {
		dataAttestor, err := it.Value()
		if err != nil {
			return nil, err
		}

		attestations = append(attestations, &data.AttestationInfo{
			Iri:       request.Iri,
			Attestor:  sdk.AccAddress(dataAttestor.Attestor).String(),
			Timestamp: types.ProtobufToGogoTimestamp(dataAttestor.Timestamp),
		})
	}

	return &data.QueryAttestationsByIRIResponse{
		Attestations: attestations,
		Pagination:   ormutil.PageResToCosmosTypes(it.PageResponse()),
	}, nil
}
