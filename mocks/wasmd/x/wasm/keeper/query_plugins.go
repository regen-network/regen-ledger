package keeper

import "github.com/cosmos/cosmos-sdk/baseapp"

type GRPCQueryRouter interface {
	Route(path string) baseapp.GRPCQueryHandler
}
