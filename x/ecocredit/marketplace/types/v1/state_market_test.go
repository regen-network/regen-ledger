package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type market struct {
	t      gocuke.TestingT
	market *Market
	err    error
}

func TestMarket(t *testing.T) {
	gocuke.NewRunner(t, &market{}).Path("./features/state_market.feature").Run()
}

func (s *market) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *market) TheMarket(a gocuke.DocString) {
	s.market = &Market{}
	err := jsonpb.UnmarshalString(a.Content, s.market)
	require.NoError(s.t, err)
}

func (s *market) TheMarketIsValidated() {
	s.err = s.market.Validate()
}

func (s *market) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *market) ExpectNoError() {
	require.NoError(s.t, s.err)
}
