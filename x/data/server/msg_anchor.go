package server

import (
	"context"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/data"
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

func (s serverImpl) anchorAndGetIRI(ctx context.Context, ch ToIRI) (iri string, id []byte, timestamp *gogotypes.Timestamp, err error) {
	iri, err = ch.ToIRI()
	if err != nil {
		return "", nil, nil, err
	}

	id, err = s.getOrCreateDataID(ctx, iri)
	if err != nil {
		return "", nil, nil, err
	}

	timestamp, err = s.anchorAndGetTimestamp(ctx, id, iri)
	if err != nil {
		return "", nil, nil, err
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

func (s serverImpl) anchorAndGetTimestamp(ctx context.Context, id []byte, iri string) (*gogotypes.Timestamp, error) {
	dataAnchor, err := s.stateStore.DataAnchorTable().Get(ctx, id)
	if err != nil {
		if !ormerrors.IsNotFound(err) {
			return nil, err
		}

		sdkCtx := sdk.UnwrapSDKContext(ctx)

		timestamp, err := gogotypes.TimestampProto(sdkCtx.BlockTime())
		if err != nil {
			return nil, sdkerrors.Wrap(err, "invalid block time")
		}

		err = s.stateStore.DataAnchorTable().Insert(ctx, &api.DataAnchor{
			Id:        id,
			Timestamp: types.GogoToProtobufTimestamp(timestamp),
		})
		if err != nil {
			return nil, err
		}

		return timestamp, sdkCtx.EventManager().EmitTypedEvent(&data.EventAnchor{Iri: iri})
	}

	return types.ProtobufToGogoTimestamp(dataAnchor.Timestamp), nil
}
