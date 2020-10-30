package server

import (
	"bytes"
	"context"
	"fmt"
	"github.com/multiformats/go-multihash"

	gocid "github.com/ipfs/go-cid"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/regen-network/regen-ledger/x/data"
)

var _ data.MsgServer = serverImpl{}

func (s serverImpl) AnchorData(goCtx context.Context, request *data.MsgAnchorDataRequest) (*data.MsgAnchorDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cid, err := gocid.Decode(request.Cid)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid CID %s", request.Cid))
	}

	cidBytes := cid.Bytes()

	if s.anchorTable.Has(ctx, cidBytes) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("CID %s is already anchored", request.Cid))
	}

	err = s.anchorCid(ctx, cidBytes)
	if err != nil {
		return nil, err
	}

	return &data.MsgAnchorDataResponse{}, nil
}

func (s serverImpl) anchorCidIfNeeded(ctx sdk.Context, cidStr string) (gocid.Cid, []byte, error) {
	cid, err := gocid.Decode(cidStr)
	if err != nil {
		return cid, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid CID %s", cidStr))
	}

	cidBytes := cid.Bytes()
	if s.anchorTable.Has(ctx, cidBytes) {
		return cid, cidBytes, nil
	}

	return cid, cidBytes, s.anchorCid(ctx, cidBytes)
}

func (s serverImpl) anchorCid(ctx sdk.Context, cidBytes []byte) error {
	timestamp, err := gogotypes.TimestampProto(ctx.BlockTime())
	if err != nil {
		return sdkerrors.Wrap(err, "invalid block time")
	}

	err = s.anchorTable.Create(ctx, cidBytes, timestamp)
	if err != nil {
		return sdkerrors.Wrap(err, "error anchoring data")
	}

	return nil
}

func (s serverImpl) SignData(goCtx context.Context, request *data.MsgSignDataRequest) (*data.MsgSignDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, cidBytes, err := s.anchorCidIfNeeded(ctx, request.Cid)
	if err != nil {
		return nil, err
	}

	var signers data.Signers
	if s.signersTable.Has(ctx, cidBytes) {
		err = s.signersTable.GetOne(ctx, cidBytes, &signers)
		if err != nil {
			return nil, err
		}

		// merge signers
		seen := map[string]bool{}
		for _, signer := range signers.Signers {
			seen[signer] = true
		}

		for _, signer := range request.Signers {
			_, found := seen[signer]
			if !found {
				signers.Signers = append(signers.Signers, signer)
			}
		}
	}

	err = s.signersTable.Save(ctx, cidBytes, &signers)
	if err != nil {
		return nil, err
	}

	return &data.MsgSignDataResponse{}, nil
}

func (s serverImpl) StoreData(goCtx context.Context, request *data.MsgStoreDataRequest) (*data.MsgStoreDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cid, cidBytes, err := s.anchorCidIfNeeded(ctx, request.Cid)
	if err != nil {
		return nil, err
	}

	store := ctx.KVStore(s.storeKey)
	if store.Has(cidBytes) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("CID %s already has stored data", request.Cid))
	}

	mh := cid.Hash()

	decodedMultihash, err := multihash.Decode(mh)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "can't retrieve multihash")
	}

	switch decodedMultihash.Name {
	case "sha2-256":
	case "blake2b-256":
	default:
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("unsupported hash function %s", decodedMultihash.Name))
	}

	reqMh, err := multihash.Sum(request.Content, decodedMultihash.Code, -1)
	if err != nil {
		return nil, sdkerrors.Wrap(err, fmt.Sprintf("unable to perform multihash"))
	}

	if !bytes.Equal(mh, reqMh) {
		return nil, sdkerrors.Wrap(err, fmt.Sprintf("unable to perform multihash"))
	}

	store.Set(cidBytes, request.Content)

	return &data.MsgStoreDataResponse{}, nil
}
