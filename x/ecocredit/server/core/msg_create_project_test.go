package core

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type createProjectSuite struct {
	*baseSuite
	alice   sdk.AccAddress
	classId string
	err     error
}

func TestCreateProject(t *testing.T) {
	runner := gocuke.NewRunner(t, &createProjectSuite{}).Path("./features/msg_create_project.feature")
	runner.Run()
}

func (s *createProjectSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
}

func (s *createProjectSuite) AliceHasCreatedACreditClass() {
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

func (s *createProjectSuite) AliceHasCreatedAProjectWithId(a string) {
	_, err := s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Issuer:              s.alice.String(),
		ClassId:             s.classId,
		Metadata:            "metadata",
		ProjectJurisdiction: "US",
		ProjectId:           a,
	})
	require.NoError(s.t, err)
}

func (s *createProjectSuite) TheProjectSequenceNumberIs(a string) {
	nextSequence, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	class, err := s.k.stateStore.ClassInfoTable().GetById(s.ctx, s.classId)
	require.NoError(s.t, err)

	err = s.k.stateStore.ProjectSequenceTable().Insert(s.ctx, &api.ProjectSequence{
		ClassKey:     class.Key,
		NextSequence: nextSequence,
	})
	require.NoError(s.t, err)
}

func (s *createProjectSuite) AliceCreatesAProject() {
	_, s.err = s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Issuer:              s.alice.String(),
		ClassId:             s.classId,
		Metadata:            "metadata",
		ProjectJurisdiction: "US",
	})
}

func (s *createProjectSuite) AliceCreatesAProjectWithId(a string) {
	_, s.err = s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Issuer:              s.alice.String(),
		ClassId:             s.classId,
		Metadata:            "metadata",
		ProjectJurisdiction: "US",
		ProjectId:           a,
	})
}

func (s *createProjectSuite) TheProjectExistsWithId(a string) {
	found, err := s.k.stateStore.ProjectInfoTable().HasById(s.ctx, a)
	require.NoError(s.t, err)
	require.True(s.t, found)

}

func (s *createProjectSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}
