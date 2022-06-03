package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/data"
)

// Resolver queries information about a resolved based on url.
func (s serverImpl) Resolver(ctx context.Context, request *data.QueryResolverRequest) (*data.QueryResolverResponse, error) {
	res, err := s.stateStore.ResolverTable().Get(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	manager := sdk.AccAddress(res.Manager).String()

	return &data.QueryResolverResponse{
		Resolver: &data.ResolverInfo{
			Url:     res.Url,
			Manager: manager,
		},
	}, nil
}
