package server

import (
	"context"

	types2 "github.com/cosmos/cosmos-sdk/types"

	gogotypes "github.com/gogo/protobuf/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/regen-network/regen-ledger/types"

	"github.com/regen-network/regen-ledger/x/data"
)

var _ data.MsgServer = serverImpl{}

func (s serverImpl) AnchorData(goCtx context.Context, request *data.MsgAnchorData) (*data.MsgAnchorDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	iri, _, timestamp, err := s.getIRIAndAnchor(ctx, request.Hash)
	if err != nil {
		return nil, err
	}

	return &data.MsgAnchorDataResponse{
		Timestamp: timestamp,
		Iri:       iri,
	}, nil
}

type ToIRI interface {
	ToIRI() (string, error)
}

func (s serverImpl) getIRIAndAnchor(ctx sdk.Context, toIRI ToIRI) (iri string, id []byte, timestamp *gogotypes.Timestamp, err error) {
	iri, err = toIRI.ToIRI()
	if err != nil {
		return iri, nil, nil, err
	}

	store := ctx.KVStore(s.storeKey)
	id = s.iriIDTable.GetOrCreateID(store, []byte(iri))
	timestamp, err = s.anchor(ctx, id, iri)
	return iri, id, timestamp, err
}

func (s serverImpl) anchor(ctx sdk.Context, id []byte, iri string) (*gogotypes.Timestamp, error) {
	timestamp, err := blockTimestamp(ctx)
	if err != nil {
		return nil, err
	}

	bz, err := timestamp.Marshal()
	if err != nil {
		return nil, err
	}

	store := ctx.KVStore(s.storeKey)
	key := AnchorTimestampKey(id)
	if store.Has(key) {
		return timestamp, nil
	}

	store.Set(key, bz)

	return timestamp, ctx.EventManager().EmitTypedEvent(&data.EventAnchorData{Iri: iri})
}

func blockTimestamp(ctx sdk.Context) (*gogotypes.Timestamp, error) {
	timestamp, err := gogotypes.TimestampProto(ctx.BlockTime())
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid block time")
	}

	return timestamp, err
}

var trueBz = []byte{1}

func (s serverImpl) SignData(goCtx context.Context, request *data.MsgSignData) (*data.MsgSignDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	iri, id, timestamp, err := s.getIRIAndAnchor(ctx, request.Hash)
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
