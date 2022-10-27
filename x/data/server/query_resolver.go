package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/regen-network/regen-ledger/x/data"
)

// Resolver queries a resolver by its unique identifier.
func (s serverImpl) Resolver(ctx context.Context, request *data.QueryResolverRequest) (*data.QueryResolverResponse, error) {
	if request.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "ID cannot be empty")
	}

	resolver, err := s.stateStore.ResolverTable().Get(ctx, request.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "resolver with ID: %d", request.Id)
	}

	manager := sdk.AccAddress(resolver.Manager).String()

	return &data.QueryResolverResponse{
		Resolver: &data.ResolverInfo{
			Id:      request.Id,
			Url:     resolver.Url,
			Manager: manager,
		},
	}, nil
}
