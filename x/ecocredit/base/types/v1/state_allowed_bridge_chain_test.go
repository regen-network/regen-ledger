package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type allowedBridgeChainSuite struct {
	t                  gocuke.TestingT
	allowedBridgeChain *AllowedBridgeChain
	err                error
}

func TestAllowedBridgeChain(t *testing.T) {
	gocuke.NewRunner(t, &allowedBridgeChainSuite{}).Path(
		"./features/state_allowed_bridge_chain.feature",
	).Run()
}

func (s *allowedBridgeChainSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *allowedBridgeChainSuite) TheAllowedBridgeChain(a gocuke.DocString) {
	s.allowedBridgeChain = &AllowedBridgeChain{}
	err := jsonpb.UnmarshalString(a.Content, s.allowedBridgeChain)
	require.NoError(s.t, err)
}

func (s *allowedBridgeChainSuite) TheAllowedBridgeChainIsValidated() {
	s.err = s.allowedBridgeChain.Validate()
}

func (s *allowedBridgeChainSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *allowedBridgeChainSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
