package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

type EventReducer struct {
	ecocreditv1.StateStore
}

func (er EventReducer) Emit(ctx context.Context, evt proto.Message) error {
	err := er.reduce(ctx, evt)
	if err != nil {
		return err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return sdkCtx.EventManager().EmitTypedEvent(evt)
}

func (er EventReducer) reduce(ctx context.Context, evt proto.Message) error {
	switch evt := evt.(type) {
	case *types.EventCreateClass:
		return er.reduceEventCreateClass(ctx, evt)
	case *types.EventUpdateClassIssuers:
		return er.reduceEventUpdateClassIssuers(ctx, evt)
	}
	return nil
}
