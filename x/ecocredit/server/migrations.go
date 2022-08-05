package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s serverImpl) RunMigrations(ctx sdk.Context, cdc codec.Codec) error {
	return nil
}
