package server

import (
	"fmt"

	types2 "github.com/cosmos/cosmos-sdk/types"

	gogotypes "github.com/gogo/protobuf/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types"

	"github.com/regen-network/regen-ledger/x/data"
)

var _ data.MsgServer = serverImpl{}

func (s serverImpl) AnchorData(ctx types.Context, request *data.MsgAnchorDataRequest) (*data.MsgAnchorDataResponse, error) {
	alreadyAnchored, iri, _, timestamp, err := s.getIRIAndAnchor(ctx, request.Hash)
	if err != nil {
		return nil, err
	}

	if alreadyAnchored {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("IRI %s was already anchored", iri))
	}

	return &data.MsgAnchorDataResponse{Timestamp: timestamp}, nil
}

type ToIRI interface {
	ToIRI() (string, error)
}

func (s serverImpl) getIRIAndAnchor(ctx types.Context, toIRI ToIRI) (anchored bool, iri string, id []byte, timestamp *gogotypes.Timestamp, err error) {
	iri, err = toIRI.ToIRI()
	if err != nil {
		return false, iri, nil, nil, err
	}

	store := ctx.KVStore(s.storeKey)
	id = s.iriIdTable.GetOrCreateID(store, []byte(iri))
	anchored, timestamp, err = s.anchor(ctx, id, iri)
	if err != nil {
		return false, "", nil, nil, err
	}

	return anchored, iri, id, timestamp, err
}

func (s serverImpl) anchor(ctx types.Context, id []byte, iri string) (bool, *gogotypes.Timestamp, error) {
	timestamp, err := blockTimestamp(ctx)
	if err != nil {
		return false, nil, err
	}

	bz, err := timestamp.Marshal()
	if err != nil {
		return false, nil, err
	}

	store := ctx.KVStore(s.storeKey)
	key := AnchorTimestampKey(id)
	if store.Has(key) {
		return false, timestamp, nil
	}

	store.Set(key, bz)

	return true, timestamp, ctx.EventManager().EmitTypedEvent(&data.EventAnchorData{Iri: iri})
}

func blockTimestamp(ctx types.Context) (*gogotypes.Timestamp, error) {
	timestamp, err := gogotypes.TimestampProto(ctx.BlockTime())
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid block time")
	}

	return timestamp, err
}

var trueBz = []byte{1}

func (s serverImpl) SignData(ctx types.Context, request *data.MsgSignDataRequest) (*data.MsgSignDataResponse, error) {
	_, iri, id, timestamp, err := s.getIRIAndAnchor(ctx, request.Hash)
	if err != nil {
		return nil, err
	}

	store := ctx.KVStore(s.storeKey)
	timestampBz, err := timestamp.Marshal()
	if err != nil {
		return nil, err
	}

	for _, signer := range request.Signers {
		addr, err := types2.AccAddressFromBech32(signer)
		if err != nil {
			return nil, err
		}

		key := IDSignerTimestampKey(id, addr)
		if store.Has(key) {
			continue
		}

		store.Set(key, timestampBz)
		// set reverse lookup key
		store.Set(SignerIDKey(addr, id), trueBz)
	}

	err = ctx.EventManager().EmitTypedEvent(&data.EventSignData{
		Iri:     iri,
		Signers: request.Signers,
	})
	if err != nil {
		return nil, err
	}

	return &data.MsgSignDataResponse{}, nil
}

func (s serverImpl) StoreRawData(ctx types.Context, request *data.MsgStoreRawDataRequest) (*data.MsgStoreRawDataResponse, error) {
	// NOTE: hash verification already has happened in MsgStoreRawDataRequest.ValidateBasic()
	// TODO: add hash verification gas cost

	_, iri, id, _, err := s.getIRIAndAnchor(ctx, request.Hash)
	if err != nil {
		return nil, err
	}

	key := RawDataKey(id)
	store := ctx.KVStore(s.storeKey)
	if store.Has(key) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("IRI %s already has stored data", iri))
	}

	store.Set(key, request.Content)

	err = ctx.EventManager().EmitTypedEvent(&data.EventStoreRawData{Iri: iri})
	if err != nil {
		return nil, err
	}

	return &data.MsgStoreRawDataResponse{}, nil
}
