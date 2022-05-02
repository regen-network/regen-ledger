package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data"
)

// DefineResolver defines a resolver URL and assigns it a new integer ID that can be used in calls to RegisterResolver.
func (s serverImpl) DefineResolver(ctx context.Context, msg *data.MsgDefineResolver) (*data.MsgDefineResolverResponse, error) {
	manager, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		return nil, err
	}

	id, err := s.stateStore.ResolverInfoTable().InsertReturningID(ctx, &api.ResolverInfo{
		Url:     msg.ResolverUrl,
		Manager: manager,
	})
	if err != nil {
		return nil, data.ErrResolverURLExists
	}

	err = sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&data.EventDefineResolver{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &data.MsgDefineResolverResponse{ResolverId: id}, nil
}
