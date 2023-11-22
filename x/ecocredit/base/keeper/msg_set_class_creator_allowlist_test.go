package keeper

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

type setClassCreatorAllowlistSuite struct {
	*baseSuite
	err error
}

func TestSetClassCreatorAllowlist(t *testing.T) {
	gocuke.NewRunner(t, &setClassCreatorAllowlistSuite{}).Path("./features/msg_set_class_creator_allowlist.feature").Run()
}

func (s *setClassCreatorAllowlistSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *setClassCreatorAllowlistSuite) AliceAttemptsToSetClassCreatorAllowlistWithProperties(a gocuke.DocString) {
	var msg *types.MsgSetClassCreatorAllowlist

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.SetClassCreatorAllowlist(s.ctx, msg)
}

func (s *setClassCreatorAllowlistSuite) ExpectClassAllowlistFlagToBe(a string) {
	isAllowed, err := strconv.ParseBool(a)
	require.NoError(s.t, err)

	res, err := s.stateStore.ClassCreatorAllowlistTable().Get(s.ctx)
	require.NoError(s.t, err)

	require.Equal(s.t, isAllowed, res.Enabled)
}

func (s *setClassCreatorAllowlistSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *setClassCreatorAllowlistSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *setClassCreatorAllowlistSuite) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}
