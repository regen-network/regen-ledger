package testsuite

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/testutil/server/configurator"
	dataserver "github.com/regen-network/regen-ledger/x/data/server"
)

func TestSuite(t *testing.T) {
	key := sdk.NewKVStoreKey("ecocredit")
	_, _, addr := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()

	configuratorFixture := configurator.NewFixture(t, []sdk.StoreKey{key}, []sdk.AccAddress{addr, addr2})
	dataserver.RegisterServices(key, configuratorFixture)
	s := NewIntegrationTestSuite(configuratorFixture)

	suite.Run(t, s)
}
