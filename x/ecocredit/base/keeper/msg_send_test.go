package keeper

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

type send struct {
	*baseSuite
	alice            sdk.AccAddress
	bob              sdk.AccAddress
	creditTypeAbbrev string
	classID          string
	classKey         uint64
	projectID        string
	projectKey       uint64
	batchDenom       string
	batchKey         uint64
	tradableAmount   string
	res              *types.MsgSendResponse
	err              error
}

func TestSend(t *testing.T) {
	gocuke.NewRunner(t, &send{}).Path("./features/msg_send.feature").Run()
}

func (s *send) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.bob = s.addr2
	s.creditTypeAbbrev = "C"
	s.classID = testClassID
	s.projectID = testProjectID
	s.batchDenom = testBatchDenom
	s.tradableAmount = "10"
}

func (s *send) ACreditTypeWithAbbreviationAndPrecision(a, b string) {
	precision, err := strconv.ParseUint(b, 10, 32)
	require.NoError(s.t, err)

	err = s.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
		Precision:    uint32(precision),
	})
	require.NoError(s.t, err)

	s.creditTypeAbbrev = a
}

func (s *send) ACreditBatch() {
	s.creditBatchSetup()
}

func (s *send) ACreditBatchWithDenom(a string) {
	s.projectSetup()

	bKey, err := s.k.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		ProjectKey: s.projectKey,
		Denom:      a,
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

func (s *send) ACreditBatchFromCreditClassWithCreditType(a string) {
	cKey, err := s.k.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id:               s.classID,
		CreditTypeAbbrev: a,
	})
	require.NoError(s.t, err)

	s.classKey = cKey

	pKey, err := s.k.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id:       s.projectID,
		ClassKey: cKey,
	})
	require.NoError(s.t, err)

	s.projectKey = pKey

	bKey, err := s.k.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		ProjectKey: s.projectKey,
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

func (s *send) AliceHasTheBatchBalance(a gocuke.DocString) {
	balance := &api.BatchBalance{}
	err := jsonpb.UnmarshalString(a.Content, balance)
	require.NoError(s.t, err)

	balance.BatchKey = s.batchKey
	balance.Address = s.alice

	// Save because the balance may already exist from setup
	err = s.stateStore.BatchBalanceTable().Save(s.ctx, balance)
	require.NoError(s.t, err)
}

func (s *send) BobHasTheBatchBalance(a gocuke.DocString) {
	balance := &api.BatchBalance{}
	err := jsonpb.UnmarshalString(a.Content, balance)
	require.NoError(s.t, err)

	balance.BatchKey = s.batchKey
	balance.Address = s.bob

	// Save because the balance may already exist from setup
	err = s.stateStore.BatchBalanceTable().Save(s.ctx, balance)
	require.NoError(s.t, err)
}

func (s *send) AliceOwnsTradableCreditAmount(a string) {
	err := s.k.stateStore.BatchBalanceTable().Insert(s.ctx, &api.BatchBalance{
		BatchKey:       s.batchKey,
		Address:        s.alice,
		TradableAmount: a,
	})
	require.NoError(s.t, err)
}

func (s *send) AliceOwnsTradableCreditsWithBatchDenom(a string) {
	batch, err := s.k.stateStore.BatchTable().GetByDenom(s.ctx, a)
	require.NoError(s.t, err)

	err = s.k.stateStore.BatchBalanceTable().Insert(s.ctx, &api.BatchBalance{
		BatchKey:       batch.Key,
		Address:        s.alice,
		TradableAmount: s.tradableAmount,
	})
	require.NoError(s.t, err)
}

func (s *send) TheBatchSupply(a gocuke.DocString) {
	supply := &api.BatchSupply{}
	err := jsonpb.UnmarshalString(a.Content, supply)
	require.NoError(s.t, err)

	supply.BatchKey = s.batchKey

	// Save because the supply may already exist from setup
	err = s.stateStore.BatchSupplyTable().Save(s.ctx, supply)
	require.NoError(s.t, err)
}

func (s *send) AliceAttemptsToSendCreditsToBobWithTradableAmount(a string) {
	s.res, s.err = s.k.Send(s.ctx, &types.MsgSend{
		Sender:    s.alice.String(),
		Recipient: s.bob.String(),
		Credits: []*types.MsgSend_SendCredits{
			{
				BatchDenom:     s.batchDenom,
				TradableAmount: a,
			},
		},
	})
}

func (s *send) AliceAttemptsToSendCreditsToBobWithRetiredAmount(a string) {
	s.res, s.err = s.k.Send(s.ctx, &types.MsgSend{
		Sender:    s.alice.String(),
		Recipient: s.bob.String(),
		Credits: []*types.MsgSend_SendCredits{
			{
				BatchDenom:    s.batchDenom,
				RetiredAmount: a,
			},
		},
	})
}

func (s *send) AliceAttemptsToSendCreditsToBobWithBatchDenom(a string) {
	s.res, s.err = s.k.Send(s.ctx, &types.MsgSend{
		Sender:    s.alice.String(),
		Recipient: s.bob.String(),
		Credits: []*types.MsgSend_SendCredits{
			{
				BatchDenom:     a,
				TradableAmount: s.tradableAmount,
			},
		},
	})
}

func (s *send) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *send) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *send) ExpectAliceBatchBalance(a gocuke.DocString) {
	expected := &api.BatchBalance{}
	err := jsonpb.UnmarshalString(a.Content, expected)
	require.NoError(s.t, err)

	balance, err := s.stateStore.BatchBalanceTable().Get(s.ctx, s.alice, s.batchKey)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.RetiredAmount, balance.RetiredAmount)
	require.Equal(s.t, expected.TradableAmount, balance.TradableAmount)
	require.Equal(s.t, expected.EscrowedAmount, balance.EscrowedAmount)
}

func (s *send) ExpectBobBatchBalance(a gocuke.DocString) {
	expected := &api.BatchBalance{}
	err := jsonpb.UnmarshalString(a.Content, expected)
	require.NoError(s.t, err)

	balance, err := s.stateStore.BatchBalanceTable().Get(s.ctx, s.bob, s.batchKey)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.RetiredAmount, balance.RetiredAmount)
	require.Equal(s.t, expected.TradableAmount, balance.TradableAmount)
	require.Equal(s.t, expected.EscrowedAmount, balance.EscrowedAmount)
}

func (s *send) ExpectBatchSupply(a gocuke.DocString) {
	expected := &api.BatchSupply{}
	err := jsonpb.UnmarshalString(a.Content, expected)
	require.NoError(s.t, err)

	balance, err := s.stateStore.BatchSupplyTable().Get(s.ctx, s.batchKey)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.RetiredAmount, balance.RetiredAmount)
	require.Equal(s.t, expected.TradableAmount, balance.TradableAmount)
}

func (s *send) BobsAddress(a string) {
	addr, err := sdk.AccAddressFromBech32(a)
	require.NoError(s.t, err)
	s.bob = addr
}

func (s *send) AliceAttemptsToSendCreditsToBobWithRetiredAmountAndJurisdiction(a, b string) {
	s.res, s.err = s.k.Send(s.ctx, &types.MsgSend{
		Sender:    s.alice.String(),
		Recipient: s.bob.String(),
		Credits: []*types.MsgSend_SendCredits{
			{
				BatchDenom:             s.batchDenom,
				RetiredAmount:          a,
				RetirementJurisdiction: b,
			},
		},
	})
	require.NoError(s.t, s.err)
}

func (s *send) AliceAttemptsToSendCreditsToBobWithRetiredAmountJurisdictionAndReason(a, b, c string) {
	s.res, s.err = s.k.Send(s.ctx, &types.MsgSend{
		Sender:    s.alice.String(),
		Recipient: s.bob.String(),
		Credits: []*types.MsgSend_SendCredits{
			{
				BatchDenom:             s.batchDenom,
				RetiredAmount:          a,
				RetirementJurisdiction: b,
				RetirementReason:       c,
			},
		},
	})
	require.NoError(s.t, s.err)
}

func (s *send) ExpectEventRetireWithProperties(a gocuke.DocString) {
	var event types.EventRetire
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *send) ExpectEventTransferWithProperties(a gocuke.DocString) {
	var event types.EventTransfer
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *send) AlicesAddress(a string) {
	addr, err := sdk.AccAddressFromBech32(a)
	require.NoError(s.t, err)
	s.alice = addr
}

func (s *send) projectSetup() {
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

	s.projectKey = pKey
}

func (s *send) creditBatchSetup() {
	s.projectSetup()

	bKey, err := s.k.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		ProjectKey: s.projectKey,
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
