package server

import (
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

func (s serverImpl) ByHash(ctx types.Context, request *data.QueryByHashRequest) (*data.QueryByHashResponse, error) {
	iri, err := request.Hash.ToIRI()
	if err != nil {
		return nil, err
	}

	store := ctx.KVStore(s.storeKey)
	id := s.iriIDTable.GetOrCreateID(store, []byte(iri))

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
		return nil, status.Error(codes.NotFound, "not found")
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
	contentHash, _, err := data.ParseIRI(sdk.GetConfig().GetBech32AccountAddrPrefix(), iri)
	if err != nil {
		return nil, err
	}

	entry := &data.ContentEntry{
		Timestamp: &timestamp,
		Signers:   signerEntries,
		Iri:       iri,
		Hash:      contentHash,
	}

	rawData := store.Get(RawDataKey(id))
	if rawData != nil {
		entry.Content = &data.Content{Sum: &data.Content_RawData{
			RawData: rawData,
		}}
	}

	return entry, nil
}

func (s serverImpl) BySigner(ctx types.Context, request *data.QueryBySignerRequest) (*data.QueryBySignerResponse, error) {
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
