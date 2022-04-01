package server

import (
	"context"

	gogotypes "github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/data"
)

func (s serverImpl) getID(ctx context.Context, iri string) ([]byte, error) {
	store := sdk.UnwrapSDKContext(ctx).KVStore(s.storeKey)
	id := s.iriIDTable.GetID(store, []byte(iri))
	if len(id) == 0 {
		return nil, status.Errorf(codes.NotFound, "can't find %s", iri)
	}

	return id, nil
}

func (s serverImpl) getEntry(ctx context.Context, store sdk.KVStore, id []byte) (*data.ContentEntry, error) {

	dataAnchor, err := s.stateStore.DataAnchorTable().Get(ctx, id)
	if err != nil {
		return nil, err
	}

	timestamp := &gogotypes.Timestamp{
		Seconds: dataAnchor.Timestamp.Seconds,
		Nanos:   dataAnchor.Timestamp.Nanos,
	}

	iri := string(s.iriIDTable.GetValue(store, id))
	contentHash, err := data.ParseIRI(iri)
	if err != nil {
		return nil, err
	}

	entry := &data.ContentEntry{
		Timestamp: timestamp,
		Iri:       iri,
		Hash:      contentHash,
	}

	return entry, nil
}
