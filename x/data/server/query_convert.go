package server

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/data"
)

// ConvertIRIToHash converts an IRI to a ContentHash.
func (s serverImpl) ConvertIRIToHash(ctx context.Context, request *data.ConvertIRIToHashRequest) (*data.ConvertIRIToHashResponse, error) {
	hash, err := data.ParseIRI(request.Iri)
	if err != nil {
		return nil, err
	}

	return &data.ConvertIRIToHashResponse{
		ContentHash: hash,
	}, nil
}

// ConvertHashToIRI converts a ContentHash to an IRI.
func (s serverImpl) ConvertHashToIRI(ctx context.Context, request *data.ConvertHashToIRIRequest) (*data.ConvertHashToIRIResponse, error) {
	if request.ContentHash == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("failed to convert content hash to IRI: %s", err)
	}

	return &data.ConvertHashToIRIResponse{
		Iri: iri,
	}, nil
}
