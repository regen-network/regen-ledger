package testsuite

import (
	"testing"

	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/types/fixture"
	datamodule "github.com/regen-network/regen-ledger/x/data/module"
)

func TestGenesis(t *testing.T) {
	ff := fixture.NewFixtureFactory(t, 2)
	ff.SetModules([]sdkmodule.AppModule{&datamodule.Module{}})
	s := NewGenesisTestSuite(ff)
	suite.Run(t, s)
}
