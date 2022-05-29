package core

import (
	"testing"

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
	alice sdk.AccAddress
	err   error
}

func TestCreateBatch(t *testing.T) {
	gocuke.NewRunner(t, &createBatchSuite{}).Path("./features/msg_create_batch.feature").Run()
}

func (s *createBatchSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
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
