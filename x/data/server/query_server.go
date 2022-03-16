package server

import (
	"context"

	gogotypes "github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/ormstore"
	"github.com/regen-network/regen-ledger/x/data"
)

var _ data.QueryServer = serverImpl{}

// ByIRI queries data based on its ContentHash.
func (s serverImpl) ByIRI(ctx context.Context, request *data.QueryByIRIRequest) (*data.QueryByIRIResponse, error) {
	id, err := s.getID(ctx, request.Iri)
	if err != nil {
		return nil, err
	}

	store := sdk.UnwrapSDKContext(ctx).KVStore(s.storeKey)
	entry, err := s.getEntry(store, id)
	if err != nil {
		return nil, err
	}

	return &data.QueryByIRIResponse{
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
func (s serverImpl) BySigner(ctx context.Context, request *data.QueryBySignerRequest) (*data.QueryBySignerResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.Signer)
	if err != nil {
		return nil, err
	}

	store := types.UnwrapSDKContext(ctx).KVStore(s.storeKey)
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
func (s serverImpl) Signers(ctx context.Context, request *data.QuerySignersRequest) (*data.QuerySignersResponse, error) {
	id, err := s.getID(ctx, request.Iri)
	if err != nil {
		return nil, err
	}

	store := types.UnwrapSDKContext(ctx).KVStore(s.storeKey)
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

func (s serverImpl) Resolvers(ctx context.Context, request *data.QueryResolversRequest) (*data.QueryResolversResponse, error) {
	id, err := s.getID(ctx, request.Iri)
	if err != nil {
		return nil, err
	}

	if request.Pagination == nil {
		request.Pagination = &query.PageRequest{}
	}

	it, err := s.stateStore.DataResolverTable().
		List(ctx, api.DataResolverPrimaryKey{}.WithId(id),
			ormlist.Paginate(ormstore.GogoPageReqToPulsarPageReq(request.Pagination)))
	if err != nil {
		return nil, err
	}

	res := &data.QueryResolversResponse{}
	for it.Next() {
		item, err := it.Value()
		if err != nil {
			return nil, err
		}

		resolverInfo, err := s.stateStore.ResolverInfoTable().Get(ctx, item.ResolverId)
		if err != nil {
			return nil, err
		}

		res.ResolverUrls = append(res.ResolverUrls, resolverInfo.Url)
	}

	if it.PageResponse() != nil {
		res.Pagination = ormstore.PulsarPageResToGogoPageRes(it.PageResponse())
	}

	return res, nil
}

func (s serverImpl) ResolverInfo(ctx context.Context, request *data.QueryResolverInfoRequest) (*data.QueryResolverInfoResponse, error) {
	res, err := s.stateStore.ResolverInfoTable().GetByUrl(ctx, request.Url)
	if err != nil {
		return nil, err
	}

	acct := sdk.AccAddress(res.Manager)

	return &data.QueryResolverInfoResponse{
		Id:      res.Id,
		Manager: acct.String(),
	}, nil
}

func (s serverImpl) getID(ctx context.Context, iri string) ([]byte, error) {
	store := sdk.UnwrapSDKContext(ctx).KVStore(s.storeKey)
	id := s.iriIDTable.GetID(store, []byte(iri))
	if len(id) == 0 {
		return nil, status.Errorf(codes.NotFound, "can't find %s", iri)
	}

	return id, nil
}
