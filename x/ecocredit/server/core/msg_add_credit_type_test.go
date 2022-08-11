package core

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type addCreditTypeSuite struct {
	*baseSuite
	err error
}

func TestAddCreditType(t *testing.T) {
	gocuke.NewRunner(t, &addCreditTypeSuite{}).Path("./features/msg_add_credit_type.feature").Run()
}

func (s *addCreditTypeSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *addCreditTypeSuite) AliceAttemptsToAddACreditTypeWithName(name string) {
	_, s.err = s.k.AddCreditType(s.ctx, &core.MsgAddCreditType{
		Authority: s.authority.String(),
		CreditType: &core.CreditType{
			Name: name,
		},
	})
}

func (s *addCreditTypeSuite) ACreditTypeWithProperties(a gocuke.DocString) {
	var msg *core.MsgAddCreditType

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	err = s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: msg.CreditType.Abbreviation,
		Name:         msg.CreditType.Name,
		Unit:         msg.CreditType.Unit,
		Precision:    msg.CreditType.Precision,
	})
	require.NoError(s.t, err)
}

func (s *addCreditTypeSuite) AliceAttemptsToAddACreditTypeWithProperties(a gocuke.DocString) {
	var msg *core.MsgAddCreditType

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.AddCreditType(s.ctx, msg)
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
