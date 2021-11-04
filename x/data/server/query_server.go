package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	gogotypes "github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/data"
)

var _ data.QueryServer = serverImpl{}

// ByHash queries data based on its ContentHash.
func (s serverImpl) ByHash(goCtx context.Context, request *data.QueryByHashRequest) (*data.QueryByHashResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)

	iri, err := request.Hash.ToIRI()
	if err != nil {
		return nil, err
	}

	store := ctx.KVStore(s.storeKey)
	id := s.iriIDTable.GetID(store, []byte(iri))
	if len(id) == 0 {
		return nil, status.Errorf(codes.NotFound, "can't find %s", iri)
	}

	entry, err := s.getEntry(store, id)
	if err != nil {
		return nil, err
	}

	return &data.QueryByHashResponse{
		Entry: entry,
	}, nil
}

func (s serverImpl) getEntry(store sdk.KVStore, id []byte) (*data.ContentEntry, error) {
	bz := store.Get(AnchorTimestampKey(id))
	if len(bz) == 0 {
		return nil, status.Error(codes.NotFound, "entry not found")
	}

	var timestamp gogotypes.Timestamp
	err := timestamp.Unmarshal(bz)
	if err != nil {
		return nil, err
	}

	var signerEntries []*data.SignerEntry
	idSignerIndexPrefix := IDSignerIndexPrefix(id)
	prefixStore := prefix.NewStore(store, idSignerIndexPrefix)
	iterator := prefixStore.Iterator(nil, nil)

	for iterator.Valid() {
		signer := sdk.AccAddress(iterator.Key())
		var timestamp gogotypes.Timestamp
		err = timestamp.Unmarshal(iterator.Value())
		if err != nil {
			return nil, err
		}

		signerEntries = append(signerEntries, &data.SignerEntry{
			Signer:    signer.String(),
			Timestamp: &timestamp,
		})

		iterator.Next()
	}

	iri := string(s.iriIDTable.GetValue(store, id))
	contentHash, err := data.ParseIRI(iri)
	if err != nil {
		return nil, err
	}

	entry := &data.ContentEntry{
		Timestamp: &timestamp,
		Iri:       iri,
		Hash:      contentHash,
	}

	return entry, nil
}

// BySigner queries data based on signers.
func (s serverImpl) BySigner(goCtx context.Context, request *data.QueryBySignerRequest) (*data.QueryBySignerResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)

	addr, err := sdk.AccAddressFromBech32(request.Signer)
	if err != nil {
		return nil, err
	}

	signerIDStore := prefix.NewStore(store, SignerIDIndexPrefix(addr))

	var entries []*data.ContentEntry
	pageRes, err := query.Paginate(signerIDStore, request.Pagination, func(key []byte, value []byte) error {
		entry, err := s.getEntry(store, key)
		if err != nil {
			return err
		}

		entries = append(entries, entry)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &data.QueryBySignerResponse{
		Entries:    entries,
		Pagination: pageRes,
	}, nil
}

// Signers queries the signers by IRI.
func (s serverImpl) Signers(goCtx context.Context, request *data.QuerySignersRequest) (*data.QuerySignersResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)

	id := s.iriIDTable.GetID(store, []byte(request.Iri))
	if len(id) == 0 {
		return nil,
			status.Errorf(codes.NotFound, "IRI %s not found", request.Iri)
	}

	signerIDStore := prefix.NewStore(store, IDSignerIndexPrefix(id))

	var signers []string
	pageRes, err := query.Paginate(signerIDStore, request.Pagination, func(key []byte, value []byte) error {
		signers = append(signers, sdk.AccAddress(key).String())
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &data.QuerySignersResponse{
		Signers:    signers,
		Pagination: pageRes,
	}, nil
}
