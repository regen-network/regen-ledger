package server

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/data"
)

// IRIByHash queries IRI based on ContentHash.
func (s serverImpl) IRIByHash(ctx context.Context, request *data.QueryIRIByHashRequest) (*data.QueryIRIByHashResponse, error) {
	if request.ContentHash == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, err
	}

	return &data.QueryIRIByHashResponse{
		Iri: iri,
	}, nil
}
