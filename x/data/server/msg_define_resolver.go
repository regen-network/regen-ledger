package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/orm/types/ormerrors"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data/v3"
)

// DefineResolver defines a resolver URL and assigns it a new integer ID that can be used in calls to RegisterResolver.
func (s serverImpl) DefineResolver(ctx context.Context, msg *data.MsgDefineResolver) (*data.MsgDefineResolverResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	definerBz, err := s.ac.StringToBytes(msg.Definer)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}

	manager := definerBz
	if msg.Public {
		manager = nil
	}

	id, err := s.stateStore.ResolverTable().InsertReturningID(ctx, &api.Resolver{
		Url:     msg.ResolverUrl,
		Manager: manager,
	})
	if err != nil {
		if ormerrors.UniqueKeyViolation.Is(err) {
			return nil, ormerrors.UniqueKeyViolation.Wrap("a resolver with the same URL and manager already exists")
		}

		return nil, err
	}

	err = sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&data.EventDefineResolver{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &data.MsgDefineResolverResponse{ResolverId: id}, nil
}
