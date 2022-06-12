package server

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/data"
)

// ConvertHashToIRI converts a ContentHash to an IRI.
func (s serverImpl) ConvertHashToIRI(_ context.Context, request *data.ConvertHashToIRIRequest) (*data.ConvertHashToIRIResponse, error) {
	if request.ContentHash == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	return &data.ConvertHashToIRIResponse{
		Iri: iri,
	}, nil
}
