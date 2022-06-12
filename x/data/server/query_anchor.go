package server

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/data"
)

// AnchorByIRI queries anchored data by IRI.
func (s serverImpl) AnchorByIRI(ctx context.Context, request *data.QueryAnchorByIRIRequest) (*data.QueryAnchorByIRIResponse, error) {
	ch, err := data.ParseIRI(request.Iri)
	if err != nil {
		return nil, err
	}

	anchor, err := s.getAnchorEntry(ctx, request.Iri, ch)
	if err != nil {
		return nil, err
	}

	return &data.QueryAnchorByIRIResponse{
		Anchor: anchor,
	}, nil
}

// AnchorByHash queries anchored data by ContentHash.
func (s serverImpl) AnchorByHash(ctx context.Context, request *data.QueryAnchorByHashRequest) (*data.QueryAnchorByHashResponse, error) {
	if request.ContentHash == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, err
	}

	anchor, err := s.getAnchorEntry(ctx, iri, request.ContentHash)
	if err != nil {
		return nil, err
	}

	return &data.QueryAnchorByHashResponse{
		Anchor: anchor,
	}, nil
}

func (s serverImpl) getAnchorEntry(ctx context.Context, iri string, ch *data.ContentHash) (*data.AnchorInfo, error) {
	dataId, err := s.stateStore.DataIDTable().GetByIri(ctx, iri)
	if err != nil {
		return nil, err
	}

	dataAnchor, err := s.stateStore.DataAnchorTable().Get(ctx, dataId.Id)
	if err != nil {
		return nil, err
	}

	return &data.AnchorInfo{
		Iri:         iri,
		ContentHash: ch,
		Timestamp:   types.ProtobufToGogoTimestamp(dataAnchor.Timestamp),
	}, nil
}
