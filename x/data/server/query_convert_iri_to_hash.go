package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/regen-network/regen-ledger/x/data"
)

// ConvertIRIToHash converts an IRI to a ContentHash.
func (s serverImpl) ConvertIRIToHash(_ context.Context, request *data.ConvertIRIToHashRequest) (*data.ConvertIRIToHashResponse, error) {
	if len(request.Iri) == 0 {
		return nil, status.Error(codes.InvalidArgument, "IRI cannot be empty")
	}

	hash, err := data.ParseIRI(request.Iri)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &data.ConvertIRIToHashResponse{
		ContentHash: hash,
	}, nil
}
