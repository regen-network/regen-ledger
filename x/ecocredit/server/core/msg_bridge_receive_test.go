package core

import (
	"strconv"
	"testing"
	"time"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type bridgeReceiveSuite struct {
	*baseSuite
	alice            sdk.AccAddress
	bob              sdk.AccAddress
	creditTypeAbbrev string
	classId          string
	projectId        string
	referenceId      string
	batchDenom       string
	batchKey         uint64
	tradableAmount   string
	startDate        *time.Time
	endDate          *time.Time
	originTx         *core.OriginTx
	metadata         string
	contract         string
	res              *core.MsgBridgeReceiveResponse
	err              error
}

func TestBridgeReceive(t *testing.T) {
	gocuke.NewRunner(t, &bridgeReceiveSuite{}).Path("./features/msg_bridge_receive.feature").Run()
}

func (s *bridgeReceiveSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.bob = s.addr // TODO: #893
	s.creditTypeAbbrev = "C"
	s.classId = "C01"
	s.projectId = "C01-001"
	s.referenceId = "VCS-001"
	s.batchDenom = "C01-001-20200101-20210101-001"
	s.tradableAmount = "10"
	s.contract = "0x0000000000000000000000000000000000000001"

	startDate, err := types.ParseDate("start date", "2020-01-01")
	require.NoError(s.t, err)

	endDate, err := types.ParseDate("end date", "2021-01-01")
	require.NoError(s.t, err)

	s.startDate = &startDate
	s.endDate = &endDate

	s.originTx = &core.OriginTx{
		Id:     "0x0",
		Source: "polygon",
		// Contract: s.contract, // TODO: #1225
	}
}

func (s *bridgeReceiveSuite) ACreditBatchExistsWithNoContract() {
	s.creditBatchSetup()
}

func (s *bridgeReceiveSuite) ACreditBatchExistsWithContract(a string) {
	s.creditBatchSetup()

	err := s.k.stateStore.BatchContractTable().Insert(s.ctx, &api.BatchContract{
		BatchKey: s.batchKey,
		Contract: a,
	})
	require.NoError(s.t, err)
}

func (s *bridgeReceiveSuite) AliceAttemptsToBridgeCreditsFromContract(a string) {
	//s.originTx.Contract = a // TODO: #1225

	s.res, s.err = s.k.BridgeReceive(s.ctx, &core.MsgBridgeReceive{
		Issuer:  s.alice.String(),
		ClassId: s.classId,
		Project: &core.MsgBridgeReceive_Project{
			ReferenceId:  s.referenceId,
			Jurisdiction: s.tradableAmount,
			Metadata:     s.metadata,
		},
		Batch: &core.MsgBridgeReceive_Batch{
			Recipient: s.bob.String(),
			Amount:    s.tradableAmount,
			StartDate: s.startDate,
			EndDate:   s.endDate,
			Metadata:  s.metadata,
		},
		OriginTx: s.originTx,
	})
}

func (s *bridgeReceiveSuite) ExpectTotalCreditBatches(a string) {
	expTotal, err := strconv.ParseUint(a, 10, 64)
	require.NoError(s.t, err)

	it, err := s.k.stateStore.BatchTable().List(s.ctx, api.BatchPrimaryKey{})
	require.NoError(s.t, err)

	var total uint64
	for it.Next() {
		total++
	}
	it.Close()

	require.Equal(s.t, expTotal, total)
}

// bare minimum setup for a credit batch
func (s *bridgeReceiveSuite) creditBatchSetup() {
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

	err = s.k.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
		ClassKey: cKey,
		Issuer:   s.alice,
	})
	require.NoError(s.t, err)

	pKey, err := s.k.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id:       s.projectId,
		ClassKey: cKey,
	})
	require.NoError(s.t, err)

	err = s.k.stateStore.ProjectSequenceTable().Insert(s.ctx, &api.ProjectSequence{
		ClassKey:     cKey,
		NextSequence: 2,
	})
	require.NoError(s.t, err)

	bKey, err := s.k.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		Issuer:     s.alice,
		ProjectKey: pKey,
		Denom:      s.batchDenom,
		Open:       true,
	})
	require.NoError(s.t, err)

	err = s.k.stateStore.BatchSupplyTable().Insert(s.ctx, &api.BatchSupply{
		BatchKey:       bKey,
		TradableAmount: "0",
	})
	require.NoError(s.t, err)

	err = s.k.stateStore.BatchSequenceTable().Insert(s.ctx, &api.BatchSequence{
		ProjectKey:   pKey,
		NextSequence: 2,
	})
	require.NoError(s.t, err)

	s.batchKey = bKey
}
