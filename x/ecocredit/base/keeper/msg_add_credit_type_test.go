package keeper

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
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
	fmt.Println("name--------------", name)
	existing, err := s.k.stateStore.CreditTypeTable().GetByName(s.ctx, name)
	require.NoError(s.t, err)

	_, s.err = s.k.AddCreditType(s.ctx, &types.MsgAddCreditType{
		Authority: s.authority.String(),
		CreditType: &types.CreditType{
			Name:         name,
			Abbreviation: existing.Abbreviation,
			Unit:         existing.Unit,
			Precision:    existing.Precision,
		},
	})
}

func (s *addCreditTypeSuite) ACreditTypeWithProperties(a gocuke.DocString) {
	var msg *types.MsgAddCreditType

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	fmt.Println("abbrication---------", msg.CreditType.Abbreviation)
	err = s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: msg.CreditType.Abbreviation,
		Name:         msg.CreditType.Name,
		Unit:         msg.CreditType.Unit,
		Precision:    msg.CreditType.Precision,
	})
	fmt.Println("Error ----------", err)
	require.NoError(s.t, err)
}

func (s *addCreditTypeSuite) AliceAttemptsToAddACreditTypeWithProperties(a gocuke.DocString) {
	var msg *types.MsgAddCreditType

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
