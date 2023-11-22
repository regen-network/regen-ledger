package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type basketBalanceSuite struct {
	t                  gocuke.TestingT
	basketBalanceSuite *BasketBalance
	err                error
}

func TestBasketBalance(t *testing.T) {
	gocuke.NewRunner(t, &basketBalanceSuite{}).Path("./features/state_basket_balance.feature").Run()
}

func (s *basketBalanceSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *basketBalanceSuite) TheBasketBalance(a gocuke.DocString) {
	s.basketBalanceSuite = &BasketBalance{}
	err := jsonpb.UnmarshalString(a.Content, s.basketBalanceSuite)
	require.NoError(s.t, err)
}

func (s *basketBalanceSuite) TheBasketBalanceIsValidated() {
	s.err = s.basketBalanceSuite.Validate()
}

func (s *basketBalanceSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *basketBalanceSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
