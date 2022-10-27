package server

import (
	"context"

	"github.com/regen-network/regen-ledger/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/regen-network/regen-ledger/x/data"
)

// AnchorByHash queries a data anchor by the ContentHash of the data.
func (s serverImpl) AnchorByHash(ctx context.Context, request *data.QueryAnchorByHashRequest) (*data.QueryAnchorByHashResponse, error) {
	if request.ContentHash == nil {
		return nil, status.Error(codes.InvalidArgument, "content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dataID, err := s.stateStore.DataIDTable().GetByIri(ctx, iri)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "data record with IRI: %s", iri)
	}

	anchor, err := s.stateStore.DataAnchorTable().Get(ctx, dataID.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &data.QueryAnchorByHashResponse{
		Anchor: &data.AnchorInfo{
			Iri:         iri,
			ContentHash: request.ContentHash,
			Timestamp:   types.ProtobufToGogoTimestamp(anchor.Timestamp),
		},
	}, nil
}
