package server

import (
	"context"

	"github.com/regen-network/regen-ledger/x/data"
)

// HashByIRI queries ContentHash based on IRI.
func (s serverImpl) HashByIRI(ctx context.Context, request *data.QueryHashByIRIRequest) (*data.QueryHashByIRIResponse, error) {
	hash, err := data.ParseIRI(request.Iri)
	if err != nil {
		return nil, err
	}

	return &data.QueryHashByIRIResponse{
		ContentHash: hash,
	}, nil
}
