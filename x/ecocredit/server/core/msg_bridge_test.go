package core

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type bridgeSuite struct {
	*baseSuite
	alice            sdk.AccAddress
	creditTypeAbbrev string
	classId          string
	projectId        string
	batchDenom       string
	batchKey         uint64
	tradableAmount   string
	target           string
	contract         string
	recipient        string
	credits          []*core.Credits
	res              *core.MsgBridgeResponse
	err              error
}

func TestBridge(t *testing.T) {
	gocuke.NewRunner(t, &bridgeSuite{}).Path("./features/msg_bridge.feature").Run()
}

func (s *bridgeSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.creditTypeAbbrev = "C"
	s.classId = "C01"
	s.projectId = "C01-001"
	s.batchDenom = "C01-001-20200101-20210101-001"
	s.tradableAmount = "10"
	s.target = "polygon"
	s.contract = "0x0000000000000000000000000000000000000001"
	s.recipient = "0x0000000000000000000000000000000000000002"
	s.credits = []*core.Credits{
		{
			BatchDenom: s.batchDenom,
			Amount:     s.tradableAmount,
		},
	}
}

func (s *bridgeSuite) ACreditBatchExistsWithABatchContractEntry() {
	s.creditBatchSetup()

	err := s.k.stateStore.BatchContractTable().Insert(s.ctx, &api.BatchContract{
		BatchKey: s.batchKey,
		Contract: s.contract,
	})
	require.NoError(s.t, err)
}

func (s *bridgeSuite) ACreditBatchExistsWithoutABatchContractEntry() {
	s.creditBatchSetup()
}

func (s *bridgeSuite) AliceOwnsCreditsFromTheCreditBatch() {
	err := s.k.stateStore.BatchBalanceTable().Insert(s.ctx, &api.BatchBalance{
		BatchKey:       s.batchKey,
		Address:        s.alice,
		TradableAmount: s.tradableAmount,
	})
	require.NoError(s.t, err)
}

func (s *bridgeSuite) AliceAttemptsToBridgeCreditsFromTheCreditBatch() {
	s.res, s.err = s.k.Bridge(s.ctx, &core.MsgBridge{
		Owner:     s.alice.String(),
		Target:    s.target,
		Recipient: s.recipient,
		Credits:   s.credits,
	})
}

func (s *bridgeSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *bridgeSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

// bare minimum setup for a credit batch
func (s *bridgeSuite) creditBatchSetup() {
	// TODO: Save for now but credit type should not exist prior to unit test #893
	err := s.k.stateStore.CreditTypeTable().Save(s.ctx, &api.CreditType{
		Abbreviation: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	cKey, err := s.k.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id:               s.classId,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	pKey, err := s.k.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id:       s.projectId,
		ClassKey: cKey,
	})
	require.NoError(s.t, err)

	bKey, err := s.k.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		ProjectKey: pKey,
		Denom:      s.batchDenom,
	})
	require.NoError(s.t, err)

	err = s.k.stateStore.BatchSupplyTable().Insert(s.ctx, &api.BatchSupply{
		BatchKey:       bKey,
		TradableAmount: s.tradableAmount,
	})
	require.NoError(s.t, err)

	s.batchKey = bKey
}
