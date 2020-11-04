package testsuite

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"

	servertestutil "github.com/regen-network/regen-ledger/testutil/server"
	dataserver "github.com/regen-network/regen-ledger/x/data/server"
)

func TestSuite(t *testing.T) {
	key := sdk.NewKVStoreKey("ecocredit")
	_, _, addr := testdata.KeyTestPubAddr()

	configuratorFixture := servertestutil.NewConfiguratorFixture(t, []sdk.StoreKey{key}, []sdk.AccAddress{addr})
	dataserver.RegisterServices(key, configuratorFixture)
	s := NewIntegrationTestSuite(configuratorFixture)

	suite.Run(t, s)
}
