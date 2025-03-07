package v1

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type basketSuite struct {
	t      gocuke.TestingT
	basket *Basket
	err    error
}

func TestBasket(t *testing.T) {
	gocuke.NewRunner(t, &basketSuite{}).Path("./features/state_basket.feature").Run()
}

func (s *basketSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *basketSuite) TheBasket(a gocuke.DocString) {
	s.basket = &Basket{}
	err := jsonpb.UnmarshalString(a.Content, s.basket)
	require.NoError(s.t, err)
}

func (s *basketSuite) TheBasketIsValidated() {
	s.err = s.basket.Validate()
}

func (s *basketSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *basketSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
