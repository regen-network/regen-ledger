package server

import (
	"context"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/regen-network/regen-ledger/x/data"
)

func (s serverImpl) getEntry(ctx context.Context, id []byte) (*data.ContentEntry, error) {
	dataId, err := s.stateStore.DataIDTable().Get(ctx, id)
	if err != nil {
		return nil, err
	}

	contentHash, err := data.ParseIRI(dataId.Iri)
	if err != nil {
		return nil, err
	}

	dataAnchor, err := s.stateStore.DataAnchorTable().Get(ctx, id)
	if err != nil {
		return nil, err
	}

	timestamp := &gogotypes.Timestamp{
		Seconds: dataAnchor.Timestamp.Seconds,
		Nanos:   dataAnchor.Timestamp.Nanos,
	}

	entry := &data.ContentEntry{
		Hash:      contentHash,
		Iri:       dataId.Iri,
		Timestamp: timestamp,
	}

	return entry, nil
}
