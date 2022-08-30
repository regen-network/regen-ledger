package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type basketClassSuite struct {
	t                gocuke.TestingT
	basketClassSuite *BasketClass
	err              error
}

func TestBasketClass(t *testing.T) {
	gocuke.NewRunner(t, &basketClassSuite{}).Path("./features/state_basket_class.feature").Run()
}

func (s *basketClassSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *basketClassSuite) TheBasketClass(a gocuke.DocString) {
	s.basketClassSuite = &BasketClass{}
	err := jsonpb.UnmarshalString(a.Content, s.basketClassSuite)
	require.NoError(s.t, err)
}

func (s *basketClassSuite) TheBasketClassIsValidated() {
	s.err = s.basketClassSuite.Validate()
}

func (s *basketClassSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *basketClassSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
