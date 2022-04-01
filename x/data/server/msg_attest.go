package server

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data"
)

// Attest allows for digital signing of an arbitrary piece of data on the blockchain.
func (s serverImpl) Attest(ctx context.Context, request *data.MsgAttest) (*data.MsgAttestResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	iri, id, timestamp, err := s.anchorAndGetIRI(ctx, request.Hash)
	if err != nil {
		return nil, err
	}

	for _, attestor := range request.Attestors {
		addr, err := sdk.AccAddressFromBech32(attestor)
		if err != nil {
			return nil, err
		}

		exists, err := s.stateStore.DataAttestorTable().Has(ctx, id, addr)
		if err != nil {
			return nil, err
		} else if exists {
			continue
		}

		err = s.stateStore.DataAttestorTable().Insert(ctx, &api.DataAttestor{
			Id:       id,
			Attestor: addr,
			Timestamp: &timestamppb.Timestamp{
				Seconds: timestamp.Seconds,
				Nanos:   timestamp.Nanos,
			},
		})
		if err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(data.GasCostPerIteration, "data/Attest attestor iteration")
	}

	err = sdkCtx.EventManager().EmitTypedEvent(&data.EventAttest{
		Iri:       iri,
		Attestors: request.Attestors,
	})
	if err != nil {
		return nil, err
	}

	return &data.MsgAttestResponse{}, nil
}
