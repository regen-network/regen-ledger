package grpc_gateway

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/regen-network/regen-ledger/types/module"
)

// Module is an interface that modules should implement to register grpc-gateway routes.
type Module interface {
	module.Module

	RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux)
}
