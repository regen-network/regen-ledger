package basket

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type basketClass struct {
	t           gocuke.TestingT
	basketClass *BasketClass
	err         error
}

func TestBasketClass(t *testing.T) {
	gocuke.NewRunner(t, &basketClass{}).Path("./features/state_basket_class.feature").Run()
}

func (s *basketClass) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *basketClass) TheBasketClass(a gocuke.DocString) {
	s.basketClass = &BasketClass{}
	err := jsonpb.UnmarshalString(a.Content, s.basketClass)
	require.NoError(s.t, err)
}

func (s *basketClass) TheBasketClassIsValidated() {
	s.err = s.basketClass.Validate()
}

func (s *basketClass) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *basketClass) ExpectNoError() {
	require.NoError(s.t, s.err)
}
