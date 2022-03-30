package server

import (
	"bytes"
	"context"

	cosmossdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	sdk "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/data"
)

// RegisterResolver registers data content hashes to the provided resolver.
func (s serverImpl) RegisterResolver(ctx context.Context, msg *data.MsgRegisterResolver) (*data.MsgRegisterResolverResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	resolverInfo, err := s.stateStore.ResolverInfoTable().Get(ctx, msg.ResolverId)
	if err != nil {
		return nil, err
	}

	manager, err := cosmossdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(resolverInfo.Manager, manager.Bytes()) {
		return nil, data.ErrUnauthorizedResolverManager
	}

	for _, datum := range msg.Data {
		_, id, _, err := s.anchorAndGetIRI(sdkCtx, datum)
		if err != nil {
			return nil, err
		}

		err = s.stateStore.DataResolverTable().Save(
			ctx,
			&api.DataResolver{
				ResolverId: msg.ResolverId,
				Id:         id,
			},
		)
		if err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(data.GasCostPerIteration, "data/RegisterResolver datum iteration")
	}

	return &data.MsgRegisterResolverResponse{}, nil
}
