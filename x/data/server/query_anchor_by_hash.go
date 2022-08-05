package server

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types"

	"github.com/regen-network/regen-ledger/x/data"
)

// AnchorByHash queries a data anchor by the ContentHash of the data.
func (s serverImpl) AnchorByHash(ctx context.Context, request *data.QueryAnchorByHashRequest) (*data.QueryAnchorByHashResponse, error) {
	if request.ContentHash == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, err
	}

	dataId, err := s.stateStore.DataIDTable().GetByIri(ctx, iri)
	if err != nil {
		return nil, sdkerrors.ErrNotFound.Wrap("data record with content hash")
	}

	anchor, err := s.stateStore.DataAnchorTable().Get(ctx, dataId.Id)
	if err != nil {
		return nil, err
	}

	return &data.QueryAnchorByHashResponse{
		Anchor: &data.AnchorInfo{
			Iri:         iri,
			ContentHash: request.ContentHash,
			Timestamp:   types.ProtobufToGogoTimestamp(anchor.Timestamp),
		},
	}, nil
}
