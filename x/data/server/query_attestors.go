package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/data"
)

// AttestorsByIRI queries attestor entries based on IRI.
func (s serverImpl) AttestorsByIRI(ctx context.Context, request *data.QueryAttestorsByIRIRequest) (*data.QueryAttestorsByIRIResponse, error) {
	dataId, err := s.stateStore.DataIDTable().GetByIri(ctx, request.Iri)
	if err != nil {
		return nil, err
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := s.stateStore.DataAttestorTable().List(
		ctx,
		api.DataAttestorIdAttestorIndexKey{}.WithId(dataId.Id),
		ormlist.Paginate(pg),
	)
	if err != nil {
		return nil, err
	}

	var attestors []*data.AttestorEntry
	for it.Next() {
		dataAttestor, err := it.Value()
		if err != nil {
			return nil, err
		}

		attestors = append(attestors, &data.AttestorEntry{
			Iri:       request.Iri,
			Attestor:  sdk.AccAddress(dataAttestor.Attestor).String(),
			Timestamp: types.ProtobufToGogoTimestamp(dataAttestor.Timestamp),
		})
	}

	pageRes, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &data.QueryAttestorsByIRIResponse{
		Attestors:  attestors,
		Pagination: pageRes,
	}, nil
}

// AttestorsByHash queries attestor entries based on ContentHash.
func (s serverImpl) AttestorsByHash(ctx context.Context, request *data.QueryAttestorsByHashRequest) (*data.QueryAttestorsByHashResponse, error) {
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

	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := s.stateStore.DataAttestorTable().List(
		ctx,
		api.DataAttestorIdAttestorIndexKey{}.WithId(dataId.Id),
		ormlist.Paginate(pg),
	)
	if err != nil {
		return nil, err
	}

	var attestors []*data.AttestorEntry
	for it.Next() {
		dataAttestor, err := it.Value()
		if err != nil {
			return nil, err
		}

		attestors = append(attestors, &data.AttestorEntry{
			Iri:       iri,
			Attestor:  sdk.AccAddress(dataAttestor.Attestor).String(),
			Timestamp: types.ProtobufToGogoTimestamp(dataAttestor.Timestamp),
		})
	}

	pageRes, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &data.QueryAttestorsByHashResponse{
		Attestors:  attestors,
		Pagination: pageRes,
	}, nil
}
