package server

import (
	"context"

	regenerrors "github.com/regen-network/regen-ledger/errors"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/data"
)

// AnchorByIRI queries a data anchor by the IRI of the data.
func (s serverImpl) AnchorByIRI(ctx context.Context, request *data.QueryAnchorByIRIRequest) (*data.QueryAnchorByIRIResponse, error) {
	if len(request.Iri) == 0 {
		return nil, regenerrors.ErrInvalidArgument.Wrap("IRI cannot be empty")
	}

	contentHash, err := data.ParseIRI(request.Iri)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrapf("failed to parse IRI: %s", err.Error())
	}

	dataID, err := s.stateStore.DataIDTable().GetByIri(ctx, request.Iri)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("data record with IRI: %s", request.Iri)
	}

	anchor, err := s.stateStore.DataAnchorTable().Get(ctx, dataID.Id)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrap(err.Error())
	}

	return &data.QueryAnchorByIRIResponse{
		Anchor: &data.AnchorInfo{
			Iri:         request.Iri,
			ContentHash: contentHash,
			Timestamp:   types.ProtobufToGogoTimestamp(anchor.Timestamp),
		},
	}, nil
}
