package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/regen-network/regen-ledger/x/data"
)

var _ data.QueryServer = serverImpl{}

func (s serverImpl) ByCid(goCtx context.Context, request *data.QueryByCidRequest) (*data.QueryByCidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cid := request.Cid

	var timestamp types.Timestamp
	store := ctx.KVStore(s.storeKey)
	bz := store.Get(AnchorKey(cid))
	if len(bz) == 0 {
		return nil, status.Error(codes.NotFound, "CID not found")
	}
	err := timestamp.Unmarshal(bz)
	if err != nil {
		return nil, err
	}

	var signers []string
	cidSignerPrefixKey := CIDSignerIndexPrefix(CIDBase64String(cid))
	prefixStore := prefix.NewStore(store, cidSignerPrefixKey)
	iterator := prefixStore.Iterator(nil, nil)

	for iterator.Valid() {
		signer := string(iterator.Key())
		signers = append(signers, signer)
		iterator.Next()
	}

	content := store.Get(DataKey(cid))

	return &data.QueryByCidResponse{
		Timestamp: &timestamp,
		Signers:   signers,
		Content:   content,
	}, err
}

func (s serverImpl) BySigner(goCtx context.Context, request *data.QueryBySignerRequest) (*data.QueryBySignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := prefix.NewStore(ctx.KVStore(s.storeKey), SignerCIDIndexPrefix(request.Signer))

	var cids [][]byte
	pageRes, err := query.Paginate(store, request.Pagination, func(key []byte, value []byte) error {
		cids = append(cids, key)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &data.QueryBySignerResponse{
		Cids:       cids,
		Pagination: pageRes,
	}, nil
}
