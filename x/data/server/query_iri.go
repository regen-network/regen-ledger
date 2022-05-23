package server

import (
	"context"
	"encoding/base64"
	"strings"

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
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("failed to convert content hash to IRI: %s", err)
	}

	return &data.QueryIRIByHashResponse{
		Iri: iri,
	}, nil
}

// IRIByRawHash queries IRI based on ContentHash_Raw properties.
func (s serverImpl) IRIByRawHash(ctx context.Context, request *data.QueryIRIByRawHashRequest) (*data.QueryIRIByRawHashResponse, error) {
	if len(request.Hash) == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("hash cannot be empty")
	}

	if request.DigestAlgorithm == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("digest algorithm cannot be unspecified")
	}

	decodedHash, err := decodeBase64String(request.Hash)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("failed to decode base64 string %s: %s", request.Hash, err)
	}

	chr := data.ContentHash_Raw{
		Hash:            decodedHash,
		DigestAlgorithm: request.DigestAlgorithm,
		MediaType:       request.MediaType,
	}

	iri, err := chr.ToIRI()
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("failed to convert content hash to IRI: %s", err)
	}

	return &data.QueryIRIByRawHashResponse{
		Iri: iri,
	}, nil
}

// IRIByGraphHash queries IRI based on ContentHash_Graph properties.
func (s serverImpl) IRIByGraphHash(ctx context.Context, request *data.QueryIRIByGraphHashRequest) (*data.QueryIRIByGraphHashResponse, error) {
	if len(request.Hash) == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("hash cannot be empty")
	}

	if request.DigestAlgorithm == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("digest algorithm cannot be unspecified")
	}

	if request.CanonicalizationAlgorithm == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("canonicalization algorithm cannot be unspecified")
	}

	decodedHash, err := decodeBase64String(request.Hash)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("failed to decode base64 string %s: %s", request.Hash, err)
	}

	chg := data.ContentHash_Graph{
		Hash:                      decodedHash,
		DigestAlgorithm:           request.DigestAlgorithm,
		CanonicalizationAlgorithm: request.CanonicalizationAlgorithm,
		MerkleTree:                request.MerkleTree,
	}

	iri, err := chg.ToIRI()
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("failed to convert content hash to IRI: %s", err)
	}

	return &data.QueryIRIByGraphHashResponse{
		Iri: iri,
	}, nil
}

func decodeBase64String(str string) ([]byte, error) {
	// replace all instances of "%2b" with "+"
	str = strings.Replace(str, "%2b", "+", -1)
	// decode base64 string to base64 bytes
	return base64.StdEncoding.DecodeString(str)
}
