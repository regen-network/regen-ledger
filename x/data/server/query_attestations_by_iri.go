package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/data"
)

// AttestationsByIRI queries data attestations by the IRI of the data.
func (s serverImpl) AttestationsByIRI(ctx context.Context, request *data.QueryAttestationsByIRIRequest) (*data.QueryAttestationsByIRIResponse, error) {
	if len(request.Iri) == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("IRI cannot be empty")
	}

	// check for valid IRI
	_, err := data.ParseIRI(request.Iri)
	if err != nil {
		return nil, err
	}

	dataID, err := s.stateStore.DataIDTable().GetByIri(ctx, request.Iri)
	if err != nil {
		return nil, sdkerrors.ErrNotFound.Wrap("data record with IRI")
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &data.QueryAttestationsByIRIResponse{
		Attestations: attestations,
		Pagination:   pageRes,
	}, nil
}
