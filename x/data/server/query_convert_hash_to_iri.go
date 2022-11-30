package server

import (
	"context"

	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/x/data/v2"
)

// ConvertHashToIRI converts a ContentHash to an IRI.
func (s serverImpl) ConvertHashToIRI(_ context.Context, request *data.ConvertHashToIRIRequest) (*data.ConvertHashToIRIResponse, error) {
	if request.ContentHash == nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap("content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrapf("failed to parse IRI: %s", err.Error())
	}

	return &data.ConvertHashToIRIResponse{
		Iri: iri,
	}, nil
}
