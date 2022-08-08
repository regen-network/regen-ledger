package core

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type addCreditTypeSuite struct {
	*baseSuite
	alice            sdk.AccAddress
	params           core.Params
	creditTypeAbbrev string
	classId          string
	res              *core.MsgAddCreditTypeResponse
	err              error
}

func TestAddCreditType(t *testing.T) {
	gocuke.NewRunner(t, &addCreditTypeSuite{}).Path("./features/msg_add_credit_type.feature").Run()
}

func (s *addCreditTypeSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.creditTypeAbbrev = "C"
	s.classId = "C01"
}

func (s *addCreditTypeSuite) ACreditTypeWithAbbreviation() {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name: ,
	})
	require.NoError(s.t, err)
}

func (s *addCreditTypeSuite) AliceAttemptsToAddACerditTypeWithName(a string) {
	s.params.AllowedClassCreators = append(s.params.AllowedClassCreators, s.alice.String())
}

func (s *addCreditTypeSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *addCreditTypeSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *addCreditTypeSuite) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *addCreditTypeSuite) ExpectTheResponse(a gocuke.DocString) {
	var res *core.MsgCreateClassResponse

	err := json.Unmarshal([]byte(a.Content), &res)
	require.NoError(s.t, err)

	require.Equal(s.t, res, s.res)
}
