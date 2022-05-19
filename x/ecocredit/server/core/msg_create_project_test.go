package core

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

type createProjectSuite struct {
	*baseSuite
	alice sdk.AccAddress
	err   error
}

func TestCreateProject(t *testing.T) {
	gocuke.NewRunner(t, &createProjectSuite{}).Path("./features/msg_create_project.feature").Run()
}

func (s *createProjectSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
}

func (s *createProjectSuite) ACreditTypeExistsWithAbbreviation(a string) {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)
}

func (s *createProjectSuite) AliceHasCreatedACreditClassWithCreditType(a string) {
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

func (s *createProjectSuite) AliceHasCreatedAProjectWithCreditClassId(a string) {
	_, s.err = s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Issuer:       s.alice.String(),
		ClassId:      a,
		Jurisdiction: "US",
	})
}

func (s *createProjectSuite) TheProjectSequenceForCreditClassIs(a, b string) {
	class, err := s.k.stateStore.ClassTable().GetById(s.ctx, a)
	require.NoError(s.t, err)

	nextSequence, err := strconv.ParseUint(b, 10, 32)
	require.NoError(s.t, err)

	err = s.k.stateStore.ProjectSequenceTable().Insert(s.ctx, &api.ProjectSequence{
		ClassKey:     class.Key,
		NextSequence: nextSequence,
	})
	require.NoError(s.t, err)
}

func (s *createProjectSuite) AliceCreatesAProjectWithCreditClassId(a string) {
	_, s.err = s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Issuer:       s.alice.String(),
		ClassId:      a,
		Jurisdiction: "US",
	})
}

func (s *createProjectSuite) TheProjectExistsWithProjectId(a string) {
	project, err := s.k.stateStore.ProjectTable().GetById(s.ctx, a)
	require.NoError(s.t, err)
	require.Equal(s.t, a, project.Id)
}
