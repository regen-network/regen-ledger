package core

import (
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type bridgeSuite struct {
	*baseSuite
	alice            sdk.AccAddress
	creditTypeAbbrev string
	classID          string
	classKey         uint64
	projectID        string
	batchDenom       string
	batchKey         uint64
	tradableAmount   string
	contract         string
	target           string
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
	s.classID = testClassID
	s.projectID = testProjectID
	s.batchDenom = testBatchDenom
	s.tradableAmount = "10"
	s.target = "polygon"
	s.contract = "0x01"
	s.recipient = "0x02"
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
		ClassKey: s.classKey,
		Contract: s.contract,
	})
	require.NoError(s.t, err)
}

func (s *bridgeSuite) ACreditBatchExistsWithoutABatchContractEntry() {
	s.creditBatchSetup()
}

func (s *bridgeSuite) AliceOwnsTradableCreditsFromTheCreditBatch() {
	err := s.k.stateStore.BatchBalanceTable().Insert(s.ctx, &api.BatchBalance{
		BatchKey:       s.batchKey,
		Address:        s.alice,
		TradableAmount: s.tradableAmount,
	})
	require.NoError(s.t, err)
}

func (s *bridgeSuite) AliceOwnsTradableCreditAmountFromTheCreditBatch(a string) {
	err := s.k.stateStore.BatchBalanceTable().Insert(s.ctx, &api.BatchBalance{
		BatchKey:       s.batchKey,
		Address:        s.alice,
		TradableAmount: a,
	})
	require.NoError(s.t, err)
}

func (s *bridgeSuite) AliceHasTheBatchBalance(a gocuke.DocString) {
	balance := &api.BatchBalance{}
	err := jsonpb.UnmarshalString(a.Content, balance)
	require.NoError(s.t, err)

	balance.BatchKey = s.batchKey
	balance.Address = s.alice

	// Save because the balance may already exist from setup
	err = s.stateStore.BatchBalanceTable().Save(s.ctx, balance)
	require.NoError(s.t, err)
}

func (s *bridgeSuite) TheBatchSupply(a gocuke.DocString) {
	supply := &api.BatchSupply{}
	err := jsonpb.UnmarshalString(a.Content, supply)
	require.NoError(s.t, err)

	supply.BatchKey = s.batchKey

	// Save because the supply may already exist from setup
	err = s.stateStore.BatchSupplyTable().Save(s.ctx, supply)
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

func (s *bridgeSuite) AliceAttemptsToBridgeCreditAmountFromTheCreditBatch(a string) {
	s.res, s.err = s.k.Bridge(s.ctx, &core.MsgBridge{
		Owner:     s.alice.String(),
		Target:    s.target,
		Recipient: s.recipient,
		Credits: []*core.Credits{
			{
				BatchDenom: s.batchDenom,
				Amount:     a,
			},
		},
	})
}

func (s *bridgeSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *bridgeSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *bridgeSuite) ExpectAliceBatchBalance(a gocuke.DocString) {
	expected := &api.BatchBalance{}
	err := jsonpb.UnmarshalString(a.Content, expected)
	require.NoError(s.t, err)

	balance, err := s.stateStore.BatchBalanceTable().Get(s.ctx, s.alice, s.batchKey)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.RetiredAmount, balance.RetiredAmount)
	require.Equal(s.t, expected.TradableAmount, balance.TradableAmount)
	require.Equal(s.t, expected.EscrowedAmount, balance.EscrowedAmount)
}

func (s *bridgeSuite) ExpectBatchSupply(a gocuke.DocString) {
	expected := &api.BatchSupply{}
	err := jsonpb.UnmarshalString(a.Content, expected)
	require.NoError(s.t, err)

	balance, err := s.stateStore.BatchSupplyTable().Get(s.ctx, s.batchKey)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.RetiredAmount, balance.RetiredAmount)
	require.Equal(s.t, expected.TradableAmount, balance.TradableAmount)
}

func (s *bridgeSuite) creditBatchSetup() {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	cKey, err := s.k.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id:               s.classID,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	s.classKey = cKey

	pKey, err := s.k.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id:       s.projectID,
		ClassKey: cKey,
	})
	require.NoError(s.t, err)

	bKey, err := s.k.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		ProjectKey: pKey,
		Denom:      s.batchDenom,
	})
	require.NoError(s.t, err)

	err = s.k.stateStore.BatchSupplyTable().Insert(s.ctx, &api.BatchSupply{
		BatchKey:        bKey,
		TradableAmount:  s.tradableAmount,
		RetiredAmount:   "0",
		CancelledAmount: "0",
	})
	require.NoError(s.t, err)

	s.batchKey = bKey
}

func (s *bridgeSuite) AliceAttemptsToBridgeCreditAmountFromTheCreditBatchTo(a, b string) {
	s.res, s.err = s.k.Bridge(s.ctx, &core.MsgBridge{
		Owner:     s.alice.String(),
		Target:    s.target,
		Recipient: b,
		Credits: []*core.Credits{
			{
				BatchDenom: s.batchDenom,
				Amount:     a,
			},
		},
	})
}

func (s *bridgeSuite) ExpectEventWithProperties(a gocuke.DocString) {
	var event core.EventBridge
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	events := s.sdkCtx.EventManager().Events()
	eventBridge := events[len(events)-1]

	require.Equal(s.t, proto.MessageName(&event), eventBridge.Type)
	err = testutil.MatchEvent(&event, eventBridge)
	require.NoError(s.t, err)
}

func (s *bridgeSuite) ACreditBatchExists() {
	s.creditBatchSetup()
}

func (s *bridgeSuite) BatchHasBatchContractEntryWithContractAddress(a string) {
	s.contract = a
	err := s.k.stateStore.BatchContractTable().Insert(s.ctx, &api.BatchContract{
		BatchKey: s.batchKey,
		ClassKey: s.classKey,
		Contract: s.contract,
	})
	require.NoError(s.t, err)
}
