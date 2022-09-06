package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type basketFeeSuite struct {
	t         gocuke.TestingT
	basketFee *BasketFee
	err       error
}

func TestBasketFee(t *testing.T) {
	gocuke.NewRunner(t, &basketFeeSuite{}).Path("./features/state_basket_fee.feature").Run()
}

func (s *basketFeeSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *basketFeeSuite) TheBasketFee(a gocuke.DocString) {
	s.basketFee = &BasketFee{}
	err := jsonpb.UnmarshalString(a.Content, s.basketFee)
	require.NoError(s.t, err)
}

func (s *basketFeeSuite) TheBasketFeeIsValidated() {
	s.err = s.basketFee.Validate()
}

func (s *basketFeeSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *basketFeeSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
