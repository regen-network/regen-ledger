package server

import (
	"context"

	"github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/data"
)

var _ data.QueryServer = serverImpl{}

func (s serverImpl) Data(goCtx context.Context, request *data.QueryDataRequest) (*data.QueryDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cid := request.Cid

	var timestamp types.Timestamp
	err := s.anchorTable.GetOne(ctx, cid, &timestamp)
	if err != nil {
		return nil, err
	}

	var signers data.Signers
	// ignore error because we at least have the timestamp
	_ = s.signersTable.GetOne(ctx, cid, &signers)

	store := ctx.KVStore(s.storeKey)
	content := store.Get(DataKey(cid))

	return &data.QueryDataResponse{
		Timestamp: &timestamp,
		Signers:   signers.Signers,
		Content:   content,
	}, err
}
