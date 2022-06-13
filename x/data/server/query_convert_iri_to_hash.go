package server

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/data"
)

// ConvertIRIToHash converts an IRI to a ContentHash.
func (s serverImpl) ConvertIRIToHash(_ context.Context, request *data.ConvertIRIToHashRequest) (*data.ConvertIRIToHashResponse, error) {
	if len(request.Iri) == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("IRI cannot be empty")
	}

	hash, err := data.ParseIRI(request.Iri)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	return &data.ConvertIRIToHashResponse{
		ContentHash: hash,
	}, nil
}
