package core

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type createBatchSuite struct {
	*baseSuite
	alice   sdk.AccAddress
	classId string
	err     error
}

func TestCreateBatch(t *testing.T) {
	runner := gocuke.NewRunner(t, &createBatchSuite{}).Path("./features/msg_create_batch.feature")
	runner.Run()
}

func (s *createBatchSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
}

func (s *createBatchSuite) AliceHasCreatedACreditClass() {
	gmAny := gomock.Any()

	fee := sdk.Coin{
		Denom:  "regen",
		Amount: sdk.NewInt(20),
	}

	creditType := &core.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "metric ton CO2 equivalent",
		Precision:    6,
	}

	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(ctx interface{}, p *core.Params) {
		p.CreditClassFee = sdk.Coins{fee}
		p.CreditTypes = []*core.CreditType{creditType}
	}).AnyTimes()

	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gmAny, gmAny, gmAny, gmAny).Return(nil).AnyTimes()

	s.bankKeeper.EXPECT().BurnCoins(gmAny, gmAny, gmAny).Return(nil).AnyTimes()

	res, err := s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.alice.String(),
		Issuers:          []string{s.alice.String()},
		CreditTypeAbbrev: creditType.Abbreviation,
		Fee:              &fee,
	})
	require.NoError(s.t, err)

	s.classId = res.ClassId
}

func (s *createBatchSuite) AliceHasCreatedAProjectWithId(a string) {
	_, err := s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Issuer:              s.alice.String(),
		ClassId:             s.classId,
		Metadata:            "metadata",
		ProjectJurisdiction: "US",
		ProjectId:           a,
	})
	require.NoError(s.t, err)
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
				Recipient:              s.alice.String(),
				TradableAmount:         "50",
				RetiredAmount:          "0",
				RetirementJurisdiction: "US",
			},
		},
		Metadata:  "metadata",
		StartDate: &startDate,
		EndDate:   &endDate,
	})
}

func (s *createBatchSuite) TheCreditBatchExistsWithDenom(a string) {
	found, err := s.k.stateStore.BatchInfoTable().HasByBatchDenom(s.ctx, a)
	require.NoError(s.t, err)
	require.True(s.t, found)

}

func (s *createBatchSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}
