package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/data"
)

// ResolverInfo queries information about a resolved based on url.
func (s serverImpl) ResolverInfo(ctx context.Context, request *data.QueryResolverInfoRequest) (*data.QueryResolverInfoResponse, error) {
	res, err := s.stateStore.ResolverInfoTable().GetByUrl(ctx, request.Url)
	if err != nil {
		return nil, err
	}

	acct := sdk.AccAddress(res.Manager)

	return &data.QueryResolverInfoResponse{
		Id:      res.Id,
		Manager: acct.String(),
	}, nil
}
