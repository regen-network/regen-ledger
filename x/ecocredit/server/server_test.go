package server

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/testutil/server/configurator"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/testsuite"
)

func TestServer(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	key := sdk.NewKVStoreKey(ecocredit.ModuleName)
	addrs := configurator.MakeTestAddresses(6)
	cfg := configurator.NewFixture(t, []sdk.StoreKey{key}, addrs, cdc)
	RegisterServices(key, cfg)
	s := testsuite.NewIntegrationTestSuite(cfg)

	suite.Run(t, s)
}
