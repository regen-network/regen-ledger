package server

import (
	"context"
	"fmt"

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

	return &data.MsgSignDataResponse{}, nil
}

func (s serverImpl) StoreData(goCtx context.Context, request *data.MsgStoreDataRequest) (*data.MsgStoreDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cid, cidBytes, err := s.anchorCidIfNeeded(ctx, request.Cid)
	if err != nil {
		return nil, err
	}

	return &data.MsgStoreDataResponse{}
}
