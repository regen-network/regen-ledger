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

// AttestationsByAttestor queries anchored data based on attestor.
func (s serverImpl) AttestationsByAttestor(ctx context.Context, request *data.QueryAttestationsByAttestorRequest) (*data.QueryAttestationsByAttestorResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.Attestor)
	if err != nil {
		return nil, err
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := s.stateStore.DataAttestorTable().List(
		ctx,
		api.DataAttestorAttestorIndexKey{}.WithAttestor(addr),
		ormlist.Paginate(pg),
	)
	if err != nil {
		return nil, err
	}

	var attestation []*data.AttestationInfo
	for it.Next() {
		dataAttestor, err := it.Value()
		if err != nil {
			return nil, err
		}

		dataId, err := s.stateStore.DataIDTable().Get(ctx, dataAttestor.Id)
		if err != nil {
			return nil, err
		}

		dataAnchor, err := s.stateStore.DataAnchorTable().Get(ctx, dataAttestor.Id)
		if err != nil {
			return nil, err
		}

		attestation = append(attestation, &data.AttestationInfo{
			Iri:       dataId.Iri,
			Attestor:  request.Attestor,
			Timestamp: types.ProtobufToGogoTimestamp(dataAnchor.Timestamp),
		})
	}

	pageRes, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &data.QueryAttestationsByAttestorResponse{
		Attestations: attestation,
		Pagination:   pageRes,
	}, nil
}

// AttestationsByIRI queries attestations based on IRI.
func (s serverImpl) AttestationsByIRI(ctx context.Context, request *data.QueryAttestationsByIRIRequest) (*data.QueryAttestationsByIRIResponse, error) {
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

// AttestationsByHash queries attestations based on ContentHash.
func (s serverImpl) AttestationsByHash(ctx context.Context, request *data.QueryAttestationsByHashRequest) (*data.QueryAttestationsByHashResponse, error) {
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

	var attestations []*data.AttestationInfo
	for it.Next() {
		dataAttestor, err := it.Value()
		if err != nil {
			return nil, err
		}

		attestations = append(attestations, &data.AttestationInfo{
			Iri:       iri,
			Attestor:  sdk.AccAddress(dataAttestor.Attestor).String(),
			Timestamp: types.ProtobufToGogoTimestamp(dataAttestor.Timestamp),
		})
	}

	pageRes, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &data.QueryAttestationsByHashResponse{
		Attestations: attestations,
		Pagination:   pageRes,
	}, nil
}
