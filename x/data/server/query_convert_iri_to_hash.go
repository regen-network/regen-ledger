package server

import (
	"context"

	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/x/data/v2"
)

// ConvertIRIToHash converts an IRI to a ContentHash.
func (s serverImpl) ConvertIRIToHash(_ context.Context, request *data.ConvertIRIToHashRequest) (*data.ConvertIRIToHashResponse, error) {
	if len(request.Iri) == 0 {
		return nil, regenerrors.ErrInvalidArgument.Wrap("IRI cannot be empty")
	}

	hash, err := data.ParseIRI(request.Iri)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	return &data.ConvertIRIToHashResponse{
		ContentHash: hash,
	}, nil
}
