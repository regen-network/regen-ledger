package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/data"
)

// ByIRI queries data based on its ContentHash.
func (s serverImpl) ByIRI(ctx context.Context, request *data.QueryByIRIRequest) (*data.QueryByIRIResponse, error) {
	id, err := s.getID(ctx, request.Iri)
	if err != nil {
		return nil, err
	}

	store := sdk.UnwrapSDKContext(ctx).KVStore(s.storeKey)

	entry, err := s.getEntry(ctx, store, id)
	if err != nil {
		return nil, err
	}

	return &data.QueryByIRIResponse{
		Entry: entry,
	}, nil
}

// ByHash queries data based on ContentHash.
func (s serverImpl) ByHash(ctx context.Context, request *data.QueryByHashRequest) (*data.QueryByHashResponse, error) {
	if request.ContentHash == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, err
	}

	id, err := s.getID(ctx, iri)
	if err != nil {
		return nil, err
	}

	store := sdk.UnwrapSDKContext(ctx).KVStore(s.storeKey)

	entry, err := s.getEntry(ctx, store, id)
	if err != nil {
		return nil, err
	}

	return &data.QueryByHashResponse{
		Entry: entry,
	}, nil
}

// ByAttestor queries data based on attestor.
func (s serverImpl) ByAttestor(ctx context.Context, request *data.QueryByAttestorRequest) (*data.QueryByAttestorResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.Attestor)
	if err != nil {
		return nil, err
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := s.stateStore.DataAttestorTable().List(
		ctx,
		api.DataAttestorAttestorIndexKey{}.WithAttestor(addr),
		ormlist.Paginate(pg),
	)
	if err != nil {
		return nil, err
	}

	store := sdk.UnwrapSDKContext(ctx).KVStore(s.storeKey)

	var entries []*data.ContentEntry
	for it.Next() {
		dataAttestor, err := it.Value()
		if err != nil {
			return nil, err
		}

		entry, err := s.getEntry(ctx, store, dataAttestor.Id)
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	pageRes, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &data.QueryByAttestorResponse{
		Entries:    entries,
		Pagination: pageRes,
	}, nil
}
