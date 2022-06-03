package core

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

type createBatchSuite struct {
	*baseSuite
	alice            sdk.AccAddress
	creditTypeAbbrev string
	classId          string
	classKey         uint64
	projectId        string
	startDate        *time.Time
	endDate          *time.Time
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
}

func (s *createBatchSuite) ACreditType() {
	// TODO: Save for now but credit type should not exist prior to unit test #893
	err := s.k.stateStore.CreditTypeTable().Save(s.ctx, &api.CreditType{
		Abbreviation: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *createBatchSuite) AliceCreatedACreditClass() {
	classKey, err := s.k.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id:               s.classId,
		Admin:            s.alice,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	s.classKey = classKey
}

func (s *createBatchSuite) AliceIsACreditClassIssuer() {
	err := s.k.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
		ClassKey: s.classKey,
		Issuer:   s.alice,
	})
	require.NoError(s.t, err)
}

func (s *createBatchSuite) AliceCreatedAProject() {
	err := s.k.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:       s.projectId,
		Admin:    s.alice,
		ClassKey: s.classKey,
	})
	require.NoError(s.t, err)
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

func (s *createBatchSuite) ACreditTypeExistsWithAbbreviation(a string) {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)
}

func (s *createBatchSuite) AliceHasCreatedACreditClassWithCreditType(a string) {
	gmAny := gomock.Any()

	fee := sdk.Coin{
		Denom:  "regen",
		Amount: sdk.NewInt(20),
	}

	allowListEnabled := false
	utils.ExpectParamGet(&allowListEnabled, s.paramsKeeper, core.KeyAllowlistEnabled, 1)
	coinFee := sdk.Coins{fee}
	utils.ExpectParamGet(&coinFee, s.paramsKeeper, core.KeyCreditClassFee, 1)

	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gmAny, gmAny, gmAny, gmAny).Return(nil).AnyTimes()

	s.bankKeeper.EXPECT().BurnCoins(gmAny, gmAny, gmAny).Return(nil).AnyTimes()

	_, err := s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.alice.String(),
		Issuers:          []string{s.alice.String()},
		CreditTypeAbbrev: a,
		Fee:              &fee,
	})
	require.NoError(s.t, err)
}

func (s *createBatchSuite) AliceHasCreatedAProjectWithCreditClassId(a string) {
	_, err := s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Issuer:       s.alice.String(),
		ClassId:      a,
		Jurisdiction: "US",
	})
	require.NoError(s.t, err)
}

func (s *createBatchSuite) AliceHasCreatedACreditBatchWithProjectId(a string) {
	startDate, err := types.ParseDate("start date", "2020-01-01")
	require.NoError(s.t, err)

	endDate, err := types.ParseDate("end date", "2021-01-01")
	require.NoError(s.t, err)

	_, s.err = s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    s.alice.String(),
		ProjectId: a,
		Issuance: []*core.BatchIssuance{
			{
				Recipient:      s.alice.String(),
				TradableAmount: "50",
			},
		},
		StartDate: &startDate,
		EndDate:   &endDate,
	})
}

func (s *createBatchSuite) AliceCreatesACreditBatchWithProjectId(a string) {
	startDate, err := types.ParseDate("start date", "2020-01-01")
	require.NoError(s.t, err)

	endDate, err := types.ParseDate("end date", "2021-01-01")
	require.NoError(s.t, err)

	_, s.err = s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    s.alice.String(),
		ProjectId: a,
		Issuance: []*core.BatchIssuance{
			{
				Recipient:      s.alice.String(),
				TradableAmount: "50",
			},
		},
		StartDate: &startDate,
		EndDate:   &endDate,
	})
}

func (s *createBatchSuite) TheCreditBatchExistsWithDenom(a string) {
	batch, err := s.k.stateStore.BatchTable().GetByDenom(s.ctx, a)
	require.NoError(s.t, err)
	require.Equal(s.t, a, batch.Denom)

}

func (s *createBatchSuite) ExpectBatchBalanceForRecipientWithAddress(a string, b gocuke.DocString) {
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
