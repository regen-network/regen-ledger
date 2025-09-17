package server

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	"github.com/regen-network/regen-ledger/types/v2"
	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	"github.com/regen-network/regen-ledger/x/data/v3"
)

// AttestationsByAttestor queries data attestations by an attestor.
func (s serverImpl) AttestationsByAttestor(ctx context.Context, request *data.QueryAttestationsByAttestorRequest) (*data.QueryAttestationsByAttestorResponse, error) {
	if len(request.Attestor) == 0 {
		return nil, regenerrors.ErrInvalidArgument.Wrap("attestor cannot be empty")
	}

	addr, err := s.addressCodec.StringToBytes(request.Attestor)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrapf("attestor: %s", err.Error())
	}

	it, err := s.stateStore.DataAttestorTable().List(
		ctx,
		api.DataAttestorAttestorIndexKey{}.WithAttestor(addr),
		ormutil.PageReqToOrmPaginate(request.Pagination),
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

	return &data.QueryAttestationsByAttestorResponse{
		Attestations: attestations,
		Pagination:   ormutil.PageResToCosmosTypes(it.PageResponse()),
	}, nil
}
