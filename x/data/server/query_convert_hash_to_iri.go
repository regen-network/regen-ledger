package server

import (
	"context"

	errors "github.com/regen-network/regen-ledger/errors"
	"github.com/regen-network/regen-ledger/x/data"
)

// ConvertHashToIRI converts a ContentHash to an IRI.
func (s serverImpl) ConvertHashToIRI(_ context.Context, request *data.ConvertHashToIRIRequest) (*data.ConvertHashToIRIResponse, error) {
	if request.ContentHash == nil {
		return nil, errors.ErrInvalidArgument.Wrap("content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, errors.ErrInvalidArgument.Wrapf("failed to parse IRI: %s", err.Error())
	}

	return &data.ConvertHashToIRIResponse{
		Iri: iri,
	}, nil
}
