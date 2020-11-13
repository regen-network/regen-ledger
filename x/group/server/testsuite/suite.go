package testsuite

import (
	"context"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/testutil/server"
	"github.com/regen-network/regen-ledger/x/group/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	fixtureFactory server.FixtureFactory
	fixture        server.Fixture

	ctx       context.Context
	msgClient types.MsgClient
	addr1     sdk.AccAddress
	addr2     sdk.AccAddress
}

func NewIntegrationTestSuite(fixtureFactory server.FixtureFactory) *IntegrationTestSuite {
	return &IntegrationTestSuite{fixtureFactory: fixtureFactory}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.fixture = s.fixtureFactory.Setup()
	s.ctx = s.fixture.Context()
	s.msgClient = types.NewMsgClient(s.fixture.TxConn())
	if len(s.fixture.Signers()) < 2 {
		s.FailNow("expected at least 2 signers, got %d", s.fixture.Signers())
	}
	s.addr1 = s.fixture.Signers()[0]
	s.addr2 = s.fixture.Signers()[1]
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.fixture.Teardown()
}

func (s *IntegrationTestSuite) TestScenario() {
}
