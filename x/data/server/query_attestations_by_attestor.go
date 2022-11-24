package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/types"
	regenerrors "github.com/regen-network/regen-ledger/types/errors"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/data"
)

// AttestationsByAttestor queries data attestations by an attestor.
func (s serverImpl) AttestationsByAttestor(ctx context.Context, request *data.QueryAttestationsByAttestorRequest) (*data.QueryAttestationsByAttestorResponse, error) {
	if len(request.Attestor) == 0 {
		return nil, regenerrors.ErrInvalidArgument.Wrap("attestor cannot be empty")
	}

	addr, err := sdk.AccAddressFromBech32(request.Attestor)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrapf("attestor: %s", err.Error())
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	it, err := s.stateStore.DataAttestorTable().List(
		ctx,
		api.DataAttestorAttestorIndexKey{}.WithAttestor(addr),
		ormlist.Paginate(pg),
	)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	var attestations []*data.AttestationInfo
	for it.Next() {
		dataAttestor, err := it.Value()
		if err != nil {
			return nil, err
		}

		dataID, err := s.stateStore.DataIDTable().Get(ctx, dataAttestor.Id)
		if err != nil {
			return nil, regenerrors.ErrNotFound.Wrap(err.Error())
		}

		attestations = append(attestations, &data.AttestationInfo{
			Iri:       dataID.Iri,
			Attestor:  request.Attestor,
			Timestamp: types.ProtobufToGogoTimestamp(dataAttestor.Timestamp),
		})
	}

	pageRes, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}

	return &data.QueryAttestationsByAttestorResponse{
		Attestations: attestations,
		Pagination:   pageRes,
	}, nil
}
