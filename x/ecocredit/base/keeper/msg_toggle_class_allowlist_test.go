package keeper

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

type toggleClassCreatorAllowlist struct {
	*baseSuite
	err error
}

func TestToggleClassCreatorAllowlist(t *testing.T) {
	gocuke.NewRunner(t, &toggleClassCreatorAllowlist{}).Path("./features/msg_toggle_class_allowlist.feature").Run()
}

func (s *toggleClassCreatorAllowlist) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *toggleClassCreatorAllowlist) AliceAttemptsToToggleClassAllowlistWithProperties(a gocuke.DocString) {
	var msg *types.MsgToggleCreditClassAllowlist

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.ToggleCreditClassAllowlist(s.ctx, msg)
}

func (s *toggleClassCreatorAllowlist) ExpectClassAllowlistFlagToBe(a string) {
	isAllowed, err := strconv.ParseBool(a)
	require.NoError(s.t, err)

	res, err := s.stateStore.AllowListEnabledTable().Get(s.ctx)
	require.NoError(s.t, err)

	require.Equal(s.t, isAllowed, res.Enabled)
}

func (s *toggleClassCreatorAllowlist) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *toggleClassCreatorAllowlist) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *toggleClassCreatorAllowlist) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}
