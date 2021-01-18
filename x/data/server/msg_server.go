package server

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/btcsuite/btcutil/base58"

	"github.com/regen-network/regen-ledger/types"

	gogotypes "github.com/gogo/protobuf/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/data"
)

var _ data.MsgServer = serverImpl{}

func idToIri(id *data.ID) string {
	hashStr := base58.Encode(id.Hash)

	switch id.Type {
	case data.IDType_ID_TYPE_RAW_UNSPECIFIED:
		ext := mediaTypeToExt(id.MediaType)
		if len(ext) > 0 {
			ext = "." + ext
		}
		return fmt.Sprintf(
			"regen:r/%s/%s%s",
			base58EncodeEnum(int32(id.DigestAlgorithm)),
			hashStr,
			ext,
		)
	case data.IDType_ID_TYPE_GRAPH:
		return fmt.Sprintf(
			"regen:g/%s/%s/%s",
			base58EncodeEnum(int32(id.GraphCanonicalizationAlgorithm)),
			base58EncodeEnum(int32(id.DigestAlgorithm)),
			hashStr,
		)
	case data.IDType_ID_TYPE_GEOGRAPHY:
		return fmt.Sprintf(
			"regen:geo/%s",
			hashStr,
		)
	default:
		panic(fmt.Errorf("unexpected IDType %d", id.Type))
	}
}

func mediaTypeToExt(mediaType data.MediaType) string {
	switch mediaType {
	default:
		return ""
	}
}

func base58EncodeEnum(algorithm int32) string {
	bz := make([]byte, 4)
	binary.LittleEndian.PutUint32(bz, uint32(algorithm))
	return base58.Encode(bz)
}

func (s serverImpl) registerIRI(ctx types.Context, iri string) (uint64, error) {
	store := ctx.KVStore(s.storeKey)
	iriIdKey := IriIdKey(iri)
	if store.Has(iriIdKey) {
		bz := store.Get(iriIdKey)
		return binary.ReadUvarint(bytes.NewBuffer(bz))
	}

	id := s.idSeq.NextVal(ctx)
	bz := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(bz, id)
	store.Set(iriIdKey, bz[:n])
	idIriKey := IdIriKey(id)
	store.Set(idIriKey, []byte(iri))
	return id, nil
}

func (s serverImpl) AnchorData(ctx types.Context, request *data.MsgAnchorDataRequest) (*data.MsgAnchorDataResponse, error) {
	contentId := request.Id
	iri := idToIri(contentId)
	id, err := s.registerIRI(ctx, iri)
	if err != nil {
		return nil, err
	}

	key := AnchorKey(id)
	store := ctx.KVStore(s.storeKey)
	if store.Has(key) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("ID %s is already anchored", iri))
	}

	timestamp, err := blockTimestamp(ctx)
	if err != nil {
		return nil, err
	}

	err = s.anchorId(ctx, timestamp, id)
	if err != nil {
		return nil, err
	}

	return &data.MsgAnchorDataResponse{Timestamp: timestamp}, nil
}

func blockTimestamp(ctx types.Context) (*gogotypes.Timestamp, error) {
	timestamp, err := gogotypes.TimestampProto(ctx.BlockTime())
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid block time")
	}

	return timestamp, err
}

func (s serverImpl) anchorCidIfNeeded(ctx types.Context, timestamp *gogotypes.Timestamp, cid []byte) error {
	store := ctx.KVStore(s.storeKey)
	key := AnchorKey(cid)
	if store.Has(key) {
		return nil
	}

	return s.anchorId(ctx, timestamp, cid)
}

func (s serverImpl) anchorId(ctx types.Context, timestamp *gogotypes.Timestamp, id uint64) error {
	bz, err := timestamp.Marshal()
	if err != nil {
		return err
	}

	store := ctx.KVStore(s.storeKey)
	key := AnchorKey(id)
	store.Set(key, bz)

	return ctx.EventManager().EmitTypedEvent(&data.EventAnchorData{})
}

var emptyBz = []byte{0}

func (s serverImpl) SignData(ctx types.Context, request *data.MsgSignDataRequest) (*data.MsgSignDataResponse, error) {
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
		store.Set(SignerCIDKey(signer, cidBz), emptyBz)
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

//func (s serverImpl) StoreData(ctx types.Context, request *data.MsgStoreDataRequest) (*data.MsgStoreDataResponse, error) {
//	cidBz := request.Cid
//
//	timestamp, err := blockTimestamp(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	err = s.anchorCidIfNeeded(ctx, timestamp, cidBz)
//	if err != nil {
//		return nil, err
//	}
//
//	key := DataKey(cidBz)
//	store := ctx.KVStore(s.storeKey)
//	if store.Has(key) {
//		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("CID %s already has stored data", cidBz))
//	}
//
//	cid, err := gocid.Cast(cidBz)
//	if err != nil {
//		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("bad CID f%x", cidBz))
//	}
//
//	mh := cid.Hash()
//
//	decodedMultihash, err := multihash.Decode(mh)
//	if err != nil {
//		return nil, sdkerrors.Wrap(err, "can't retrieve multihash")
//	}
//
//	switch decodedMultihash.Name {
//	case "sha2-256":
//		// TODO: gas
//	case "blake2b-256":
//		// TODO: gas
//	default:
//		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("unsupported hash function %s", decodedMultihash.Name))
//	}
//
//	reqMh, err := multihash.Sum(request.Content, decodedMultihash.Code, -1)
//	if err != nil {
//		return nil, sdkerrors.Wrap(err, "unable to perform multihash")
//	}
//
//	if !bytes.Equal(mh, reqMh) {
//		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "multihash verification failed, please check that cid and content are correct")
//	}
//
//	store.Set(key, request.Content)
//
//	err = ctx.EventManager().EmitTypedEvent(&data.EventStoreData{Cid: cidBz})
//	if err != nil {
//		return nil, err
//	}
//
//	return &data.MsgStoreDataResponse{}, nil
//}

func (s serverImpl) StoreRawData(context types.Context, request *data.MsgStoreRawDataRequest) (*data.MsgStoreRawDataResponse, error) {
	panic("implement me")
}
