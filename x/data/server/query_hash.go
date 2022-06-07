package server

import (
	"context"

	"github.com/regen-network/regen-ledger/x/data"
)

// ConvertIRIToHash queries ContentHash based on IRI.
func (s serverImpl) ConvertIRIToHash(ctx context.Context, request *data.ConvertIRIToHashRequest) (*data.ConvertIRIToHashResponse, error) {
	hash, err := data.ParseIRI(request.Iri)
	if err != nil {
		return nil, err
	}

	return &data.ConvertIRIToHashResponse{
		ContentHash: hash,
	}, nil
}
