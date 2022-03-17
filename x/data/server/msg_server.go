package server

import (
	"bytes"
	"context"

	"github.com/gogo/protobuf/proto"
	gogotypes "github.com/gogo/protobuf/types"

	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	sdk "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/data"
)

var _ data.MsgServer = serverImpl{}

func (s serverImpl) AnchorData(goCtx context.Context, request *data.MsgAnchorData) (*data.MsgAnchorDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	iri, _, timestamp, err := s.anchorAndGetIRI(ctx, request.Hash)
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

func (s serverImpl) anchorAndGetIRI(ctx sdk.Context, toIRI ToIRI) (iri string, id []byte, timestamp *gogotypes.Timestamp, err error) {
	iri, err = toIRI.ToIRI()
	if err != nil {
		return "", nil, nil, err
	}

	store := ctx.KVStore(s.storeKey)
	id = s.iriIDTable.GetOrCreateID(store, []byte(iri))
	timestamp, err = s.anchorAndGetTimestamp(ctx, id, iri)
	return iri, id, timestamp, err
}

func (s serverImpl) anchorAndGetTimestamp(ctx sdk.Context, id []byte, iri string) (*gogotypes.Timestamp, error) {
	store := ctx.KVStore(s.storeKey)
	key := AnchorTimestampKey(id)
	bz := store.Get(key)
	if len(bz) != 0 {
		var timestamp gogotypes.Timestamp
		err := proto.Unmarshal(bz, &timestamp)
		if err != nil {
			return nil, err
		}

		return &timestamp, nil
	}

	timestamp, err := blockTimestamp(ctx)
	if err != nil {
		return nil, err
	}

	bz, err = proto.Marshal(timestamp)
	if err != nil {
		return nil, err
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
	iri, id, timestamp, err := s.anchorAndGetIRI(ctx, request.Hash)
	if err != nil {
		return nil, err
	}

	store := ctx.KVStore(s.storeKey)
	timestampBz, err := timestamp.Marshal()
	if err != nil {
		return nil, err
	}

	for _, signer := range request.Signers {
		addr, err := cosmossdk.AccAddressFromBech32(signer)
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

func (s serverImpl) DefineResolver(ctx context.Context, msg *data.MsgDefineResolver) (*data.MsgDefineResolverResponse, error) {
	manager, err := cosmossdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		return nil, err
	}

	id, err := s.stateStore.ResolverInfoTable().InsertReturningID(ctx, &api.ResolverInfo{
		Url:     msg.ResolverUrl,
		Manager: manager.Bytes(),
	})
	if err != nil {
		return nil, err
	}

	return &data.MsgDefineResolverResponse{ResolverId: id}, nil
}

func (s serverImpl) RegisterResolver(ctx context.Context, msg *data.MsgRegisterResolver) (*data.MsgRegisterResolverResponse, error) {
	resolverInfo, err := s.stateStore.ResolverInfoTable().Get(ctx, msg.ResolverId)
	if err != nil {
		return nil, err
	}

	manager, err := cosmossdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(resolverInfo.Manager, manager.Bytes()) {
		return nil, data.ErrUnauthorizedResolverManager
	}

	for _, datum := range msg.Data {
		_, id, _, err := s.anchorAndGetIRI(sdk.UnwrapSDKContext(ctx), datum)
		if err != nil {
			return nil, err
		}
		err = s.stateStore.DataResolverTable().Save(
			ctx,
			&api.DataResolver{
				ResolverId: msg.ResolverId,
				Id:         id,
			},
		)
		if err != nil {
			return nil, err
		}
	}
	return &data.MsgRegisterResolverResponse{}, nil
}
