package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/data"
)

// AnchorByIRI queries a data anchor by the IRI of the data.
func (s serverImpl) AnchorByIRI(ctx context.Context, request *data.QueryAnchorByIRIRequest) (*data.QueryAnchorByIRIResponse, error) {
	if len(request.Iri) == 0 {
		return nil, status.Error(codes.InvalidArgument, "IRI cannot be empty")
	}

	contentHash, err := data.ParseIRI(request.Iri)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse IRI: %s", err.Error())
	}

	dataID, err := s.stateStore.DataIDTable().GetByIri(ctx, request.Iri)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "data record with IRI: %s", request.Iri)
	}

	anchor, err := s.stateStore.DataAnchorTable().Get(ctx, dataID.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &data.QueryAnchorByIRIResponse{
		Anchor: &data.AnchorInfo{
			Iri:         request.Iri,
			ContentHash: contentHash,
			Timestamp:   types.ProtobufToGogoTimestamp(anchor.Timestamp),
		},
	}, nil
}
