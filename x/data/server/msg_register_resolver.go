package server

import (
	"bytes"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data"
)

// RegisterResolver registers data content hashes to the provided resolver.
func (s serverImpl) RegisterResolver(ctx context.Context, msg *data.MsgRegisterResolver) (*data.MsgRegisterResolverResponse, error) {
	resolverInfo, err := s.stateStore.ResolverInfoTable().Get(ctx, msg.ResolverId)
	if err != nil {
		return nil, err
	}

	manager, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(resolverInfo.Manager, manager) {
		return nil, data.ErrUnauthorizedResolverManager
	}

	for _, ch := range msg.ContentHashes {
		_, id, _, err := s.anchorAndGetIRI(ctx, ch)
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

		sdk.UnwrapSDKContext(ctx).GasMeter().ConsumeGas(data.GasCostPerIteration, "data/RegisterResolver content hash iteration")
	}

	return &data.MsgRegisterResolverResponse{}, nil
}
