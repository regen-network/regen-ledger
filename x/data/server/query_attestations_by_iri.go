package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/types"
	regenerrors "github.com/regen-network/regen-ledger/types/errors"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/data"
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

	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	it, err := s.stateStore.DataAttestorTable().List(
		ctx,
		api.DataAttestorIdAttestorIndexKey{}.WithId(dataID.Id),
		ormlist.Paginate(pg),
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

	pageRes, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}

	return &data.QueryAttestationsByIRIResponse{
		Attestations: attestations,
		Pagination:   pageRes,
	}, nil
}
