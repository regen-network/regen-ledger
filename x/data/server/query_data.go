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

// AnchorByIRI queries data based on its ContentHash.
func (s serverImpl) AnchorByIRI(ctx context.Context, request *data.QueryAnchorByIRIRequest) (*data.QueryAnchorByIRIResponse, error) {
	contentHash, err := data.ParseIRI(request.Iri)
	if err != nil {
		return nil, err
	}

	entry, err := s.getEntry(ctx, contentHash, request.Iri)
	if err != nil {
		return nil, err
	}

	return &data.QueryAnchorByIRIResponse{
		Entry: entry,
	}, nil
}

// AnchorByHash queries data based on ContentHash.
func (s serverImpl) AnchorByHash(ctx context.Context, request *data.QueryAnchorByHashRequest) (*data.QueryAnchorByHashResponse, error) {
	if request.ContentHash == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, err
	}

	entry, err := s.getEntry(ctx, request.ContentHash, iri)
	if err != nil {
		return nil, err
	}

	return &data.QueryAnchorByHashResponse{
		Entry: entry,
	}, nil
}

// AnchorsByAttestor queries data based on attestor.
func (s serverImpl) AnchorsByAttestor(ctx context.Context, request *data.QueryAnchorsByAttestorRequest) (*data.QueryAnchorsByAttestorResponse, error) {
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

	var entries []*data.ContentEntry
	for it.Next() {
		dataAttestor, err := it.Value()
		if err != nil {
			return nil, err
		}

		dataId, err := s.stateStore.DataIDTable().Get(ctx, dataAttestor.Id)
		if err != nil {
			return nil, err
		}

		contentHash, err := data.ParseIRI(dataId.Iri)
		if err != nil {
			return nil, err
		}

		dataAnchor, err := s.stateStore.DataAnchorTable().Get(ctx, dataAttestor.Id)
		if err != nil {
			return nil, err
		}

		entries = append(entries, &data.ContentEntry{
			ContentHash: contentHash,
			Iri:         dataId.Iri,
			Timestamp:   types.ProtobufToGogoTimestamp(dataAnchor.Timestamp),
		})
	}

	pageRes, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &data.QueryAnchorsByAttestorResponse{
		Entries:    entries,
		Pagination: pageRes,
	}, nil
}

func (s serverImpl) getEntry(ctx context.Context, ch *data.ContentHash, iri string) (*data.ContentEntry, error) {
	dataId, err := s.stateStore.DataIDTable().GetByIri(ctx, iri)
	if err != nil {
		return nil, err
	}

	dataAnchor, err := s.stateStore.DataAnchorTable().Get(ctx, dataId.Id)
	if err != nil {
		return nil, err
	}

	return &data.ContentEntry{
		ContentHash: ch,
		Iri:         iri,
		Timestamp:   types.ProtobufToGogoTimestamp(dataAnchor.Timestamp),
	}, nil
}
