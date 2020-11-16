package server

import (
	"bytes"
	"context"
	"fmt"

	gogotypes "github.com/gogo/protobuf/types"
	gocid "github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/data"
)

var _ data.MsgServer = serverImpl{}

func (s serverImpl) AnchorData(goCtx context.Context, request *data.MsgAnchorDataRequest) (*data.MsgAnchorDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cidBz := request.Cid
	key := AnchorKey(cidBz)
	store := ctx.KVStore(s.storeKey)
	if store.Has(key) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("CID f%x is already anchored", cidBz))
	}

	timestamp, err := blockTimestamp(ctx)
	if err != nil {
		return nil, err
	}

	err = s.anchorCid(ctx, timestamp, cidBz)
	if err != nil {
		return nil, err
	}

	return &data.MsgAnchorDataResponse{Timestamp: timestamp}, nil
}

func blockTimestamp(ctx sdk.Context) (*gogotypes.Timestamp, error) {
	timestamp, err := gogotypes.TimestampProto(ctx.BlockTime())
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid block time")
	}

	return timestamp, err
}

func (s serverImpl) anchorCidIfNeeded(ctx sdk.Context, timestamp *gogotypes.Timestamp, cid []byte) error {
	store := ctx.KVStore(s.storeKey)
	key := AnchorKey(cid)
	if store.Has(key) {
		return nil
	}

	return s.anchorCid(ctx, timestamp, cid)
}

func (s serverImpl) anchorCid(ctx sdk.Context, timestamp *gogotypes.Timestamp, cidBytes []byte) error {
	bz, err := timestamp.Marshal()
	if err != nil {
		return err
	}

	store := ctx.KVStore(s.storeKey)
	key := AnchorKey(cidBytes)
	store.Set(key, bz)

	return ctx.EventManager().EmitTypedEvent(&data.EventAnchorData{Cid: cidBytes})
}

var emptyBz = []byte{0}

func (s serverImpl) SignData(goCtx context.Context, request *data.MsgSignDataRequest) (*data.MsgSignDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cidBz := request.Cid

	timestamp, err := blockTimestamp(ctx)
	if err != nil {
		return nil, err
	}

	err = s.anchorCidIfNeeded(ctx, timestamp, cidBz)
	if err != nil {
		return nil, err
	}

	cidStr := CIDBase64String(cidBz)
	store := ctx.KVStore(s.storeKey)

	for _, signer := range request.Signers {
		key := CIDSignerKey(cidStr, signer)
		if store.Has(key) {
			continue
		}

		store.Set(key, emptyBz)
		// set reverse lookup key
		store.Set(SignerCIDKey(signer, cidStr), emptyBz)
	}

	err = ctx.EventManager().EmitTypedEvent(&data.EventSignData{
		Cid:     cidBz,
		Signers: request.Signers,
	})
	if err != nil {
		return nil, err
	}

	return &data.MsgSignDataResponse{}, nil
}

func (s serverImpl) StoreData(goCtx context.Context, request *data.MsgStoreDataRequest) (*data.MsgStoreDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cidBz := request.Cid

	timestamp, err := blockTimestamp(ctx)
	if err != nil {
		return nil, err
	}

	err = s.anchorCidIfNeeded(ctx, timestamp, cidBz)
	if err != nil {
		return nil, err
	}

	key := DataKey(cidBz)
	store := ctx.KVStore(s.storeKey)
	if store.Has(key) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("CID %s already has stored data", cidBz))
	}

	cid, err := gocid.Cast(cidBz)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("bad CID f%x", cidBz))
	}

	mh := cid.Hash()

	decodedMultihash, err := multihash.Decode(mh)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "can't retrieve multihash")
	}

	switch decodedMultihash.Name {
	case "sha2-256":
		// TODO: gas
	case "blake2b-256":
		// TODO: gas
	default:
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("unsupported hash function %s", decodedMultihash.Name))
	}

	reqMh, err := multihash.Sum(request.Content, decodedMultihash.Code, -1)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "unable to perform multihash")
	}

	if !bytes.Equal(mh, reqMh) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "multihash verification failed, please check that cid and content are correct")
	}

	store.Set(key, request.Content)

	err = ctx.EventManager().EmitTypedEvent(&data.EventStoreData{Cid: cidBz})
	if err != nil {
		return nil, err
	}

	return &data.MsgStoreDataResponse{}, nil
}
