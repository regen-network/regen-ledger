package core

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func (k Keeper) BatchesByIssuer(ctx context.Context, request *core.QueryBatchesByIssuerRequest) (*core.QueryBatchesByIssuerResponse, error) {
	pg, err := GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}
	issuerAcc, err := sdk.AccAddressFromBech32(request.Issuer)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.ClassIssuerTable().List(ctx, ecocreditv1.ClassIssuerIssuerIndexKey{}.WithIssuer(issuerAcc))
	if err != nil {
		return nil, err
	}
	classIdSet := make(map[uint64]struct{})
	for it.Next() {
		classIssuer, err := it.Value()
		if err != nil {
			return nil, err
		}
		classIdSet[classIssuer.ClassId] = struct{}{}
	}
}
