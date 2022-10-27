package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/regen-network/regen-ledger/x/data"
)

// ConvertHashToIRI converts a ContentHash to an IRI.
func (s serverImpl) ConvertHashToIRI(_ context.Context, request *data.ConvertHashToIRIRequest) (*data.ConvertHashToIRIResponse, error) {
	if request.ContentHash == nil {
		return nil, status.Error(codes.InvalidArgument, "content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse IRI: %s", err.Error())
	}

	return &data.ConvertHashToIRIResponse{
		Iri: iri,
	}, nil
}
