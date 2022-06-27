package server

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
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

	id, err := s.stateStore.ResolverTable().InsertReturningID(ctx, &api.Resolver{
		Url:     msg.ResolverUrl,
		Manager: manager,
	})
	if err != nil {
		if err == ormerrors.UniqueKeyViolation {
			return nil, fmt.Errorf("a resolver with the same URL and manager already exists: %s", err)
		} else {
			return nil, err
		}
	}

	err = sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&data.EventDefineResolver{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &data.MsgDefineResolverResponse{ResolverId: id}, nil
}
