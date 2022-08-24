package basket

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type basketBalance struct {
	t             gocuke.TestingT
	basketBalance *BasketBalance
	err           error
}

func TestBasketBalance(t *testing.T) {
	gocuke.NewRunner(t, &basketBalance{}).Path("./features/state_basket_balance.feature").Run()
}

func (s *basketBalance) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *basketBalance) TheBasketBalance(a gocuke.DocString) {
	s.basketBalance = &BasketBalance{}
	err := jsonpb.UnmarshalString(a.Content, s.basketBalance)
	require.NoError(s.t, err)
}

func (s *basketBalance) TheBasketBalanceIsValidated() {
	s.err = s.basketBalance.Validate()
}

func (s *basketBalance) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *basketBalance) ExpectNoError() {
	require.NoError(s.t, s.err)
}
