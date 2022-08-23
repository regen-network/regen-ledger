package server

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/data"
)

// AnchorByIRI queries a data anchor by the IRI of the data.
func (s serverImpl) AnchorByIRI(ctx context.Context, request *data.QueryAnchorByIRIRequest) (*data.QueryAnchorByIRIResponse, error) {
	if len(request.Iri) == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("IRI cannot be empty")
	}

	contentHash, err := data.ParseIRI(request.Iri)
	if err != nil {
		return nil, err
	}

	dataID, err := s.stateStore.DataIDTable().GetByIri(ctx, request.Iri)
	if err != nil {
		return nil, sdkerrors.ErrNotFound.Wrapf("data record with IRI")
	}

	anchor, err := s.stateStore.DataAnchorTable().Get(ctx, dataID.Id)
	if err != nil {
		return nil, err
	}

	return &data.QueryAnchorByIRIResponse{
		Anchor: &data.AnchorInfo{
			Iri:         request.Iri,
			ContentHash: contentHash,
			Timestamp:   types.ProtobufToGogoTimestamp(anchor.Timestamp),
		},
	}, nil
}
