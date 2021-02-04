package server

import (
	"fmt"

	"github.com/regen-network/regen-ledger/types"

	"github.com/regen-network/regen-ledger/x/data"
)

var _ data.MsgServer = serverImpl{}

func (s serverImpl) AnchorData(ctx types.Context, request *data.MsgAnchorDataRequest) (*data.MsgAnchorDataResponse, error) {
	return nil, fmt.Errorf("not implemented")
	//cidBz := request.Cid
	//key := AnchorKey(cidBz)
	//store := ctx.KVStore(s.storeKey)
	//if store.Has(key) {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("CID f%x is already anchored", cidBz))
	//}
	//
	//timestamp, err := blockTimestamp(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = s.anchorCid(ctx, timestamp, cidBz)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &data.MsgAnchorDataResponse{Timestamp: timestamp}, nil
}

//func blockTimestamp(ctx types.Context) (*gogotypes.Timestamp, error) {
//	timestamp, err := gogotypes.TimestampProto(ctx.BlockTime())
//	if err != nil {
//		return nil, sdkerrors.Wrap(err, "invalid block time")
//	}
//
//	return timestamp, err
//}
//
//func (s serverImpl) anchorCidIfNeeded(ctx types.Context, timestamp *gogotypes.Timestamp, cid []byte) error {
//	store := ctx.KVStore(s.storeKey)
//	key := AnchorKey(cid)
//	if store.Has(key) {
//		return nil
//	}
//
//	return s.anchorCid(ctx, timestamp, cid)
//}
//
//func (s serverImpl) anchorCid(ctx types.Context, timestamp *gogotypes.Timestamp, cidBytes []byte) error {
//	bz, err := timestamp.Marshal()
//	if err != nil {
//		return err
//	}
//
//	store := ctx.KVStore(s.storeKey)
//	key := AnchorKey(cidBytes)
//	store.Set(key, bz)
//
//	return ctx.EventManager().EmitTypedEvent(&data.EventAnchorData{Cid: cidBytes})
//}

//var emptyBz = []byte{0}

func (s serverImpl) SignData(ctx types.Context, request *data.MsgSignDataRequest) (*data.MsgSignDataResponse, error) {
	return nil, fmt.Errorf("not implemented")
	//cidBz := request.Cid
	//
	//timestamp, err := blockTimestamp(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = s.anchorCidIfNeeded(ctx, timestamp, cidBz)
	//if err != nil {
	//	return nil, err
	//}
	//
	//cidStr := CIDBase64String(cidBz)
	//store := ctx.KVStore(s.storeKey)
	//
	//for _, signer := range request.Signers {
	//	key := CIDSignerKey(cidStr, signer)
	//	if store.Has(key) {
	//		continue
	//	}
	//
	//	store.Set(key, emptyBz)
	//	// set reverse lookup key
	//	store.Set(SignerCIDKey(signer, cidBz), emptyBz)
	//}
	//
	//err = ctx.EventManager().EmitTypedEvent(&data.EventSignData{
	//	Cid:     cidBz,
	//	Signers: request.Signers,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &data.MsgSignDataResponse{}, nil
}

func (s serverImpl) StoreRawData(ctx types.Context, request *data.MsgStoreRawDataRequest) (*data.MsgStoreRawDataResponse, error) {
	return nil, fmt.Errorf("not implemented")
	//cidBz := request.Cid
	//
	//timestamp, err := blockTimestamp(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = s.anchorCidIfNeeded(ctx, timestamp, cidBz)
	//if err != nil {
	//	return nil, err
	//}
	//
	//key := DataKey(cidBz)
	//store := ctx.KVStore(s.storeKey)
	//if store.Has(key) {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("CID %s already has stored data", cidBz))
	//}
	//
	//cid, err := gocid.Cast(cidBz)
	//if err != nil {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("bad CID f%x", cidBz))
	//}
	//
	//mh := cid.Hash()
	//
	//decodedMultihash, err := multihash.Decode(mh)
	//if err != nil {
	//	return nil, sdkerrors.Wrap(err, "can't retrieve multihash")
	//}
	//
	//switch decodedMultihash.Name {
	//case "sha2-256":
	//	// TODO: gas
	//case "blake2b-256":
	//	// TODO: gas
	//default:
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("unsupported hash function %s", decodedMultihash.Name))
	//}
	//
	//reqMh, err := multihash.Sum(request.Content, decodedMultihash.Code, -1)
	//if err != nil {
	//	return nil, sdkerrors.Wrap(err, "unable to perform multihash")
	//}
	//
	//if !bytes.Equal(mh, reqMh) {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "multihash verification failed, please check that cid and content are correct")
	//}
	//
	//store.Set(key, request.Content)
	//
	//err = ctx.EventManager().EmitTypedEvent(&data.EventStoreData{Cid: cidBz})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &data.MsgStoreDataResponse{}, nil
}
