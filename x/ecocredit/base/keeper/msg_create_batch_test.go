//nolint:revive,stylecheck
package keeper

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

type createBatchSuite struct {
	*baseSuite
	alice            sdk.AccAddress
	bob              sdk.AccAddress
	creditTypeAbbrev string
	classID          string
	classKey         uint64
	projectID        string
	projectKey       uint64
	tradableAmount   string
	startDate        *time.Time
	endDate          *time.Time
	originTx         *types.OriginTx
	res              *types.MsgCreateBatchResponse
	err              error
}

func TestCreateBatch(t *testing.T) {
	gocuke.NewRunner(t, &createBatchSuite{}).Path("./features/msg_create_batch.feature").Run()
}

func (s *createBatchSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.bob = s.addr2
	s.creditTypeAbbrev = "C"
	s.classID = testClassID
	s.projectID = testProjectID
	s.tradableAmount = "10"

	startDate, err := regentypes.ParseDate("start date", "2020-01-01")
	require.NoError(s.t, err)

	endDate, err := regentypes.ParseDate("end date", "2021-01-01")
	require.NoError(s.t, err)

	s.startDate = &startDate
	s.endDate = &endDate
}

func (s *createBatchSuite) ACreditTypeWithAbbreviation(a string) {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)

	s.creditTypeAbbrev = a
}

func (s *createBatchSuite) ACreditTypeWithAbbreviationAndPrecision(a string, b string) {
	precision, err := strconv.ParseUint(b, 10, 32)
	require.NoError(s.t, err)

	err = s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
		Precision:    uint32(precision),
	})
	require.NoError(s.t, err)

	s.creditTypeAbbrev = a
}

func (s *createBatchSuite) ACreditClassWithIssuerAlice() {
	cKey, err := s.k.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id:               s.classID,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	err = s.k.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
		ClassKey: cKey,
		Issuer:   s.alice,
	})
	require.NoError(s.t, err)

	s.classKey = cKey
}

func (s *createBatchSuite) ACreditClassWithClassIdAndIssuerAlice(a string) {
	creditTypeAbbrev := base.GetCreditTypeAbbrevFromClassID(a)

	cKey, err := s.k.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id:               a,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	err = s.k.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
		ClassKey: cKey,
		Issuer:   s.alice,
	})
	require.NoError(s.t, err)

	s.classKey = cKey
}

func (s *createBatchSuite) AProjectWithProjectId(a string) {
	classID := base.GetClassIDFromProjectID(a)

	class, err := s.k.stateStore.ClassTable().GetById(s.ctx, classID)
	require.NoError(s.t, err)

	pKey, err := s.k.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id:       a,
		ClassKey: class.Key,
	})
	require.NoError(s.t, err)

	seq := s.getProjectSequence(a)

	// Save because project sequence may already exist
	err = s.k.stateStore.ProjectSequenceTable().Save(s.ctx, &api.ProjectSequence{
		ClassKey:     class.Key,
		NextSequence: seq + 1,
	})
	require.NoError(s.t, err)

	s.projectKey = pKey
}

func (s *createBatchSuite) ABatchSequenceWithProjectIdAndNextSequence(a string, b string) {
	project, err := s.stateStore.ProjectTable().GetById(s.ctx, a)
	require.NoError(s.t, err)

	nextSequence, err := strconv.ParseUint(b, 10, 64)
	require.NoError(s.t, err)

	err = s.stateStore.BatchSequenceTable().Insert(s.ctx, &api.BatchSequence{
		ProjectKey:   project.Key,
		NextSequence: nextSequence,
	})
	require.NoError(s.t, err)
}

func (s *createBatchSuite) AnOriginTxIndex(a gocuke.DocString) {
	var originTxIndex api.OriginTxIndex
	err := jsonpb.UnmarshalString(a.Content, &originTxIndex)
	require.NoError(s.t, err)

	err = s.k.stateStore.OriginTxIndexTable().Insert(s.ctx, &originTxIndex)
	require.NoError(s.t, err)
}

func (s *createBatchSuite) ABatchContract(a gocuke.DocString) {
	var batchContract api.BatchContract
	err := jsonpb.UnmarshalString(a.Content, &batchContract)
	require.NoError(s.t, err)

	err = s.k.stateStore.BatchContractTable().Insert(s.ctx, &batchContract)
	require.NoError(s.t, err)
}

func (s *createBatchSuite) EcocreditModulesAddress(a string) {
	addr, err := sdk.AccAddressFromBech32(a)
	require.NoError(s.t, err)
	s.k.moduleAddress = addr
}

func (s *createBatchSuite) OriginTx(a gocuke.DocString) {
	var ot types.OriginTx
	err := json.Unmarshal([]byte(a.Content), &ot)
	require.NoError(s.t, err)
	s.originTx = &ot
}

func (s *createBatchSuite) AliceAttemptsToCreateABatchWithProjectId(a string) {
	s.res, s.err = s.k.CreateBatch(s.ctx, &types.MsgCreateBatch{
		Issuer:    s.alice.String(),
		ProjectId: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:      s.bob.String(),
				TradableAmount: s.tradableAmount,
			},
		},
		StartDate: s.startDate,
		EndDate:   s.endDate,
	})
}

func (s *createBatchSuite) BobAttemptsToCreateABatchWithProjectId(a string) {
	s.res, s.err = s.k.CreateBatch(s.ctx, &types.MsgCreateBatch{
		Issuer:    s.bob.String(),
		ProjectId: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:      s.alice.String(),
				TradableAmount: s.tradableAmount,
			},
		},
		StartDate: s.startDate,
		EndDate:   s.endDate,
	})
}

func (s *createBatchSuite) AliceAttemptsToCreateABatchWithProjectIdStartDateAndEndDate(a, b, c string) {
	startDate, err := regentypes.ParseDate("start date", b)
	require.NoError(s.t, err)

	endDate, err := regentypes.ParseDate("end date", c)
	require.NoError(s.t, err)

	s.res, s.err = s.k.CreateBatch(s.ctx, &types.MsgCreateBatch{
		Issuer:    s.alice.String(),
		ProjectId: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:      s.bob.String(),
				TradableAmount: s.tradableAmount,
			},
		},
		StartDate: &startDate,
		EndDate:   &endDate,
	})
}

func (s *createBatchSuite) AliceAttemptsToCreateABatchWithProjectIdAndTradableAmount(a, b string) {
	s.res, s.err = s.k.CreateBatch(s.ctx, &types.MsgCreateBatch{
		Issuer:    s.alice.String(),
		ProjectId: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:      s.bob.String(),
				TradableAmount: b,
			},
		},
		StartDate: s.startDate,
		EndDate:   s.endDate,
	})
}

func (s *createBatchSuite) AliceAttemptsToCreateABatchWithProjectIdAndRetiredAmount(a, b string) {
	s.res, s.err = s.k.CreateBatch(s.ctx, &types.MsgCreateBatch{
		Issuer:    s.alice.String(),
		ProjectId: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:     s.bob.String(),
				RetiredAmount: b,
			},
		},
		StartDate: s.startDate,
		EndDate:   s.endDate,
	})
}

func (s *createBatchSuite) AliceAttemptsToCreateABatchWithProjectIdAndOriginTx(a string, b gocuke.DocString) {
	var originTx types.OriginTx
	err := jsonpb.UnmarshalString(b.Content, &originTx)
	require.NoError(s.t, err)

	s.res, s.err = s.k.CreateBatch(s.ctx, &types.MsgCreateBatch{
		Issuer:    s.alice.String(),
		ProjectId: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:      s.bob.String(),
				TradableAmount: s.tradableAmount,
			},
		},
		StartDate: s.startDate,
		EndDate:   s.endDate,
		OriginTx:  &originTx,
	})
}

func (s *createBatchSuite) AliceAttemptsToCreateABatchWithProjectIdAndIssuance(a string, b gocuke.DocString) {
	var issuance []*types.BatchIssuance
	// unmarshal with json because issuance array is not a proto message
	err := json.Unmarshal([]byte(b.Content), &issuance)
	require.NoError(s.t, err)

	s.res, s.err = s.k.CreateBatch(s.ctx, &types.MsgCreateBatch{
		Issuer:    s.alice.String(),
		ProjectId: a,
		Issuance:  issuance,
		StartDate: s.startDate,
		EndDate:   s.endDate,
	})
}

func (s *createBatchSuite) AliceAttemptsToCreateABatchWithProperties(a gocuke.DocString) {
	var msg types.MsgCreateBatch
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	msg.Issuer = s.alice.String()
	s.res, s.err = s.k.CreateBatch(s.ctx, &msg)
}

func (s *createBatchSuite) CreatesABatchFromProjectAndIssuesTradableCreditsTo(a string, b string, c string) {
	s.res, s.err = s.k.CreateBatch(s.ctx, &types.MsgCreateBatch{
		Issuer:    s.alice.String(),
		ProjectId: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:      c,
				TradableAmount: b,
			},
		},
		StartDate: s.startDate,
		EndDate:   s.endDate,
		OriginTx:  s.originTx,
	})
	require.NoError(s.t, s.err)
}

func (s *createBatchSuite) CreatesABatchFromProjectAndIssuesRetiredCreditsToFrom(a, b, c, d string) {
	s.res, s.err = s.k.CreateBatch(s.ctx, &types.MsgCreateBatch{
		Issuer:    s.alice.String(),
		ProjectId: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:              c,
				RetiredAmount:          b,
				RetirementJurisdiction: d,
			},
		},
		StartDate: s.startDate,
		EndDate:   s.endDate,
		OriginTx:  s.originTx,
	})
	require.NoError(s.t, s.err)
}

func (s *createBatchSuite) CreatesABatchFromProjectAndIssuesRetiredCreditsToFromWithReason(a, b, c, d, e string) {
	s.res, s.err = s.k.CreateBatch(s.ctx, &types.MsgCreateBatch{
		Issuer:    s.alice.String(),
		ProjectId: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:              c,
				RetiredAmount:          b,
				RetirementJurisdiction: d,
				RetirementReason:       e,
			},
		},
		StartDate: s.startDate,
		EndDate:   s.endDate,
		OriginTx:  s.originTx,
	})
	require.NoError(s.t, s.err)
}

func (s *createBatchSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *createBatchSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *createBatchSuite) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *createBatchSuite) ExpectRecipientBatchBalanceWithAddress(a string, b gocuke.DocString) {
	var expected api.BatchBalance
	err := jsonpb.UnmarshalString(b.Content, &expected)
	require.NoError(s.t, err)

	batch, err := s.stateStore.BatchTable().GetByDenom(s.ctx, s.res.BatchDenom)
	require.NoError(s.t, err)

	recipient, err := sdk.AccAddressFromBech32(a)
	require.NoError(s.t, err)

	balance, err := s.stateStore.BatchBalanceTable().Get(s.ctx, recipient, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.RetiredAmount, balance.RetiredAmount)
	require.Equal(s.t, expected.TradableAmount, balance.TradableAmount)
	require.Equal(s.t, expected.EscrowedAmount, balance.EscrowedAmount)
}

func (s *createBatchSuite) ExpectBatchSupply(a gocuke.DocString) {
	var expected api.BatchSupply
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	batch, err := s.stateStore.BatchTable().GetByDenom(s.ctx, s.res.BatchDenom)
	require.NoError(s.t, err)

	balance, err := s.stateStore.BatchSupplyTable().Get(s.ctx, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.RetiredAmount, balance.RetiredAmount)
	require.Equal(s.t, expected.TradableAmount, balance.TradableAmount)
	require.Equal(s.t, expected.CancelledAmount, balance.CancelledAmount)
}

func (s *createBatchSuite) ExpectBatchSequenceWithProjectIdAndNextSequence(a string, b string) {
	nextSequence, err := strconv.ParseUint(b, 10, 64)
	require.NoError(s.t, err)

	project, err := s.stateStore.ProjectTable().GetById(s.ctx, a)
	require.NoError(s.t, err)

	batchSequence, err := s.stateStore.BatchSequenceTable().Get(s.ctx, project.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, nextSequence, batchSequence.NextSequence)
}

func (s *createBatchSuite) ExpectBatchProperties(a gocuke.DocString) {
	var expected types.Batch
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	batch, err := s.stateStore.BatchTable().GetByDenom(s.ctx, expected.Denom)
	require.NoError(s.t, err)

	coreBatch := new(types.Batch)
	require.NoError(s.t, ormutil.PulsarToGogoSlow(batch, coreBatch))

	// set the properties that get set during state machine execution.
	expected.Key = coreBatch.Key
	expected.ProjectKey = coreBatch.ProjectKey
	expected.Issuer = coreBatch.Issuer
	expected.IssuanceDate = coreBatch.IssuanceDate

	require.Equal(s.t, expected, *coreBatch)
}

func (s *createBatchSuite) ExpectBatchContract(a gocuke.DocString) {
	var expected types.BatchContract
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	batchContract, err := s.stateStore.BatchContractTable().GetByClassKeyContract(s.ctx, s.classKey, expected.Contract)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.BatchKey, batchContract.BatchKey)
}

func (s *createBatchSuite) ExpectTheResponse(a gocuke.DocString) {
	var res types.MsgCreateBatchResponse
	err := jsonpb.UnmarshalString(a.Content, &res)
	require.NoError(s.t, err)

	require.Equal(s.t, &res, s.res)
}

func (s *createBatchSuite) ExpectEventRetireWithProperties(a gocuke.DocString) {
	var event types.EventRetire
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	eventRetire, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, eventRetire)
	require.NoError(s.t, err)
}

func (s *createBatchSuite) ExpectEventMintWithProperties(a gocuke.DocString) {
	var event types.EventMint
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)
	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *createBatchSuite) ExpectEventTransferWithProperties(a gocuke.DocString) {
	var event types.EventTransfer
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)
	event.Sender = s.k.moduleAddress.String()

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *createBatchSuite) ExpectEventCreateBatchWithProperties(a gocuke.DocString) {
	var event types.EventCreateBatch
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *createBatchSuite) getProjectSequence(projectID string) uint64 {
	str := strings.Split(projectID, "-")
	seq, err := strconv.ParseUint(str[1], 10, 32)
	require.NoError(s.t, err)
	return seq
}
