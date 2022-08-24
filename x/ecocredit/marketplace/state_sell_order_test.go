package marketplace

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type sellOrder struct {
	t         gocuke.TestingT
	sellOrder *SellOrder
	err       error
}

func TestSellOrder(t *testing.T) {
	gocuke.NewRunner(t, &sellOrder{}).Path("./features/state_sell_order.feature").Run()
}

func (s *sellOrder) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *sellOrder) TheSellOrder(a gocuke.DocString) {
	s.sellOrder = &SellOrder{}
	err := jsonpb.UnmarshalString(a.Content, s.sellOrder)
	require.NoError(s.t, err)
}

func (s *sellOrder) TheSellOrderIsValidated() {
	s.err = s.sellOrder.Validate()
}

func (s *sellOrder) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *sellOrder) ExpectNoError() {
	require.NoError(s.t, s.err)
}
