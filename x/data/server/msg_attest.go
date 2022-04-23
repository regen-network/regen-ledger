package server

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/data"
)

// Attest allows for digital signing of an arbitrary piece of data on the blockchain.
func (s serverImpl) Attest(ctx context.Context, request *data.MsgAttest) (*data.MsgAttestResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	var newEntries []*data.AttestorEntry

	for _, ch := range request.ContentHashes {
		iri, id, _, err := s.anchorAndGetIRI(ctx, ch)
		if err != nil {
			return nil, err
		}

		addr, err := sdk.AccAddressFromBech32(request.Attestor)
		if err != nil {
			return nil, err
		}

		found, err := s.stateStore.DataAttestorTable().Has(ctx, id, addr)
		if err != nil {
			return nil, err
		} else if found {
			// an attestor attesting to the same piece of date is a no-op
			continue
		}

		timestamp := timestamppb.New(sdkCtx.BlockTime())

		err = s.stateStore.DataAttestorTable().Insert(ctx, &api.DataAttestor{
			Id:        id,
			Attestor:  addr,
			Timestamp: timestamp,
		})
		if err != nil {
			return nil, err
		}

		err = sdkCtx.EventManager().EmitTypedEvent(&data.EventAttest{
			Iri:      iri,
			Attestor: request.Attestor,
		})
		if err != nil {
			return nil, err
		}

		newEntries = append(newEntries, &data.AttestorEntry{
			Iri:       iri,
			Attestor:  addr.String(),
			Timestamp: types.ProtobufToGogoTimestamp(timestamp),
		})

		sdkCtx.GasMeter().ConsumeGas(data.GasCostPerIteration, "data/Attest content hash iteration")
	}

	return &data.MsgAttestResponse{NewEntries: newEntries}, nil
}
