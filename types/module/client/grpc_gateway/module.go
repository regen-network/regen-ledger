package grpc_gateway

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/regen-network/regen-ledger/types/module"
)

type Module interface {
	module.ModuleBase

	RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux)
}
