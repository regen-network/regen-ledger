package server

import (
	"context"

	gogotypes "github.com/gogo/protobuf/types"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data"
)

type ToIRI interface {
	ToIRI() (string, error)
}

// Anchor anchors a piece of data to the blockchain based on its secure hash.
func (s serverImpl) Anchor(ctx context.Context, request *data.MsgAnchor) (*data.MsgAnchorResponse, error) {
	iri, _, timestamp, err := s.anchorAndGetIRI(ctx, request.Hash)
	if err != nil {
		return nil, err
	}

	return &data.MsgAnchorResponse{
		Timestamp: timestamp,
		Iri:       iri,
	}, nil
}

func (s serverImpl) anchorAndGetIRI(ctx context.Context, ch ToIRI) (iri string, id []byte, timestamp *gogotypes.Timestamp, err error) {
	iri, err = ch.ToIRI()
	if err != nil {
		return "", nil, nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(s.storeKey)
	id = s.iriIDTable.GetOrCreateID(store, []byte(iri))

	timestamp, err = s.anchorAndGetTimestamp(ctx, id, iri)

	// consume additional gas whenever we verify the provided hash
	sdkCtx.GasMeter().ConsumeGas(data.GasCostPerIteration, "data hash verification")

	return iri, id, timestamp, err
}

func (s serverImpl) anchorAndGetTimestamp(ctx context.Context, id []byte, iri string) (*gogotypes.Timestamp, error) {
	dataAnchor, err := s.stateStore.DataAnchorTable().Get(ctx, id)
	if err != nil {
		if !ormerrors.IsNotFound(err) {
			return nil, err
		} else {
			sdkCtx := sdk.UnwrapSDKContext(ctx)

			timestamp, err := gogotypes.TimestampProto(sdkCtx.BlockTime())
			if err != nil {
				return nil, sdkerrors.Wrap(err, "invalid block time")
			}

			err = s.stateStore.DataAnchorTable().Insert(ctx, &api.DataAnchor{
				Id: id,
				Timestamp: &timestamppb.Timestamp{
					Seconds: timestamp.Seconds,
					Nanos:   timestamp.Nanos,
				},
			})
			if err != nil {
				return nil, err
			}

			return timestamp, sdkCtx.EventManager().EmitTypedEvent(&data.EventAnchor{Iri: iri})
		}
	}

	timestamp := &gogotypes.Timestamp{
		Seconds: dataAnchor.Timestamp.Seconds,
		Nanos:   dataAnchor.Timestamp.Nanos,
	}

	return timestamp, nil
}
