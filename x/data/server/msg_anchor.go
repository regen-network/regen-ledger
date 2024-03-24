package server

import (
	"context"
	"time"

	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/protobuf/types/known/timestamppb"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	"github.com/regen-network/regen-ledger/types/v2"
	"github.com/regen-network/regen-ledger/x/data/v3"
)

type ToIRI interface {
	ToIRI() (string, error)
}

// Anchor anchors a piece of data to the blockchain based on its secure hash.
func (s serverImpl) Anchor(ctx context.Context, request *data.MsgAnchor) (*data.MsgAnchorResponse, error) {
	iri, _, timestamp, err := s.anchorAndGetIRI(ctx, request.ContentHash)
	if err != nil {
		return nil, err
	}

	return &data.MsgAnchorResponse{
		Iri:       iri,
		Timestamp: timestamp,
	}, nil
}

func (s serverImpl) anchorAndGetIRI(ctx context.Context, ch ToIRI) (iri string, id []byte, timestamp time.Time, err error) {
	iri, err = ch.ToIRI()
	if err != nil {
		return "", nil, time.Time{}, err
	}

	id, err = s.getOrCreateDataID(ctx, iri)
	if err != nil {
		return "", nil, time.Time{}, err
	}

	timestamp, err = s.anchorAndGetTimestamp(ctx, id, iri)
	if err != nil {
		return "", nil, timestamp, err
	}

	return iri, id, timestamp, err
}

func (s serverImpl) getOrCreateDataID(ctx context.Context, iri string) (id []byte, err error) {
	dataID := &api.DataID{Iri: ""}

	for collisions := 0; dataID.Iri != iri; collisions++ {
		id = s.iriHasher.CreateID([]byte(iri), collisions)

		dataID, err = s.stateStore.DataIDTable().Get(ctx, id)
		if err != nil {
			if !ormerrors.IsNotFound(err) {
				return nil, err
			}

			dataID = &api.DataID{
				Id:  id,
				Iri: iri,
			}
			err = s.stateStore.DataIDTable().Insert(ctx, dataID)
			if err != nil {
				return nil, err
			}
		}

		// consume additional gas whenever we create or verify the data ID
		sdk.UnwrapSDKContext(ctx).GasMeter().ConsumeGas(data.GasCostPerIteration, "create/verify data id")
	}

	return id, nil
}

func (s serverImpl) anchorAndGetTimestamp(ctx context.Context, id []byte, iri string) (time.Time, error) {
	dataAnchor, err := s.stateStore.DataAnchorTable().Get(ctx, id)
	if err != nil {
		if !ormerrors.IsNotFound(err) {
			return time.Time{}, err
		}

		sdkCtx := sdk.UnwrapSDKContext(ctx)
		timestamp := sdkCtx.BlockTime()
		err = s.stateStore.DataAnchorTable().Insert(ctx, &api.DataAnchor{
			Id:        id,
			Timestamp: timestamppb.New(timestamp),
		})
		if err != nil {
			return timestamp, err
		}

		return timestamp, sdkCtx.EventManager().EmitTypedEvent(&data.EventAnchor{Iri: iri})
	}

	return dataAnchor.Timestamp.AsTime(), nil
}
