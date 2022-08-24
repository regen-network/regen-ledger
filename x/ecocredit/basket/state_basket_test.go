package basket

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type basket struct {
	t      gocuke.TestingT
	basket *Basket
	err    error
}

func TestBasket(t *testing.T) {
	gocuke.NewRunner(t, &basket{}).Path("./features/state_basket.feature").Run()
}

func (s *basket) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *basket) TheBasket(a gocuke.DocString) {
	s.basket = &Basket{}
	err := jsonpb.UnmarshalString(a.Content, s.basket)
	require.NoError(s.t, err)
}

func (s *basket) TheBasketIsValidated() {
	s.err = s.basket.Validate()
}

func (s *basket) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *basket) ExpectNoError() {
	require.NoError(s.t, s.err)
}
