package core

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type createBatchSuite struct {
	*baseSuite
	alice            sdk.AccAddress
	creditTypeAbbrev string
	classId          string
	classKey         uint64
	projectId        string
	projectKey       uint64
	startDate        *time.Time
	endDate          *time.Time
	tradableAmount   string
	res              *core.MsgCreateBatchResponse
	err              error
}

func TestCreateBatch(t *testing.T) {
	gocuke.NewRunner(t, &createBatchSuite{}).Path("./features/msg_create_batch.feature").Run()
}

func (s *createBatchSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.creditTypeAbbrev = "C"
	s.classId = "C01"
	s.projectId = "C01-001"

	startDate, err := types.ParseDate("start date", "2020-01-01")
	require.NoError(s.t, err)

	endDate, err := types.ParseDate("end date", "2021-01-01")
	require.NoError(s.t, err)

	s.startDate = &startDate
	s.endDate = &endDate
	s.tradableAmount = "50"
}

func (s *createBatchSuite) ACreditType() {
	// TODO: Save for now but credit type should not exist prior to unit test #893
	err := s.k.stateStore.CreditTypeTable().Save(s.ctx, &api.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Name:         s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *createBatchSuite) ACreditTypeWithAbbreviation(a string) {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)

	s.creditTypeAbbrev = a
}

func (s *createBatchSuite) ACreditClassWithIssuerAlice() {
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

	s.classKey = cKey
}

func (s *createBatchSuite) ACreditClassWithClassIdAndIssuerAlice(a string) {
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassId(a)

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

func (s *createBatchSuite) AProject() {
	err := s.k.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:       s.projectId,
		ClassKey: s.classKey,
	})
	require.NoError(s.t, err)
}

func (s *createBatchSuite) AProjectWithProjectId(a string) {
	classId := core.GetClassIdFromProjectId(a)

	class, err := s.k.stateStore.ClassTable().GetById(s.ctx, classId)
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

func (s *createBatchSuite) ACreditBatchWithDenom(a string) {
	err := s.k.stateStore.BatchTable().Insert(s.ctx, &api.Batch{
		Denom:      a,
		ProjectKey: s.projectKey,
	})
	require.NoError(s.t, err)

	seq := s.getBatchSequence(a)

	// Save because batch sequence may already exist
	err = s.k.stateStore.BatchSequenceTable().Save(s.ctx, &api.BatchSequence{
		ProjectKey:   s.projectKey,
		NextSequence: seq + 1,
	})
	require.NoError(s.t, err)
}

func (s *createBatchSuite) AnOriginTxIndex(a gocuke.DocString) {
	originTxIndex := &api.OriginTxIndex{}
	err := jsonpb.UnmarshalString(a.Content, originTxIndex)
	require.NoError(s.t, err)

	err = s.k.stateStore.OriginTxIndexTable().Insert(s.ctx, originTxIndex)
	require.NoError(s.t, err)
}

func (s *createBatchSuite) AliceAttemptsToCreateACreditBatchWithProjectIdStartDateAndEndDate(a, b, c string) {
	startDate, err := types.ParseDate("start date", b)
	require.NoError(s.t, err)

	endDate, err := types.ParseDate("end date", c)
	require.NoError(s.t, err)

	s.res, s.err = s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    s.alice.String(),
		ProjectId: a,
		Issuance: []*core.BatchIssuance{
			{
				Recipient:      s.alice.String(),
				TradableAmount: s.tradableAmount,
			},
		},
		StartDate: &startDate,
		EndDate:   &endDate,
	})
}

func (s *createBatchSuite) AliceAttemptsToCreateACreditBatchWithOriginTx(a gocuke.DocString) {
	originTx := &core.OriginTx{}
	err := jsonpb.UnmarshalString(a.Content, originTx)
	require.NoError(s.t, err)

	s.res, s.err = s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    s.alice.String(),
		ProjectId: s.projectId,
		Issuance: []*core.BatchIssuance{
			{
				Recipient:      s.alice.String(),
				TradableAmount: s.tradableAmount,
			},
		},
		StartDate: s.startDate,
		EndDate:   s.endDate,
		OriginTx:  originTx,
	})
}

func (s *createBatchSuite) AliceAttemptsToCreateACreditBatchWithTheIssuance(a gocuke.DocString) {
	var issuance []*core.BatchIssuance
	// unmarshal with json because issuance array is not a proto message
	err := json.Unmarshal([]byte(a.Content), &issuance)
	require.NoError(s.t, err)

	s.res, s.err = s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    s.alice.String(),
		ProjectId: s.projectId,
		Issuance:  issuance,
		StartDate: s.startDate,
		EndDate:   s.endDate,
	})
}

func (s *createBatchSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *createBatchSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *createBatchSuite) ExpectCreditBatchWithDenom(a string) {
	batch, err := s.k.stateStore.BatchTable().GetByDenom(s.ctx, a)
	require.NoError(s.t, err)
	require.Equal(s.t, a, batch.Denom)
}

func (s *createBatchSuite) ExpectRecipientBatchBalanceWithAddress(a string, b gocuke.DocString) {
	expected := &api.BatchBalance{}
	err := jsonpb.UnmarshalString(b.Content, expected)
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

func (s *createBatchSuite) getProjectSequence(projectId string) uint64 {
	str := strings.Split(projectId, "-")
	seq, err := strconv.ParseUint(str[1], 10, 32)
	require.NoError(s.t, err)
	return seq
}

func (s *createBatchSuite) getBatchSequence(batchDenom string) uint64 {
	str := strings.Split(batchDenom, "-")
	seq, err := strconv.ParseUint(str[4], 10, 32)
	require.NoError(s.t, err)
	return seq
}
