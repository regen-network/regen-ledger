package core

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type updateClassFeesSuite struct {
	*baseSuite
	err error
}

func TestUpdateClassFees(t *testing.T) {
	gocuke.NewRunner(t, &updateClassFeesSuite{}).Path("./features/msg_update_class_fee.feature").Run()
}

func (s *updateClassFeesSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *updateClassFeesSuite) AliceAttemptsToUpdateClassFeeWithProperties(a gocuke.DocString) {
	var msg *core.MsgUpdateClassFees

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.UpdateClassFees(s.ctx, msg)
}

func (s *updateClassFeesSuite) ExpectClassFeesWithProperties(a gocuke.DocString) {
	var expected *api.ClassFees
	err := json.Unmarshal([]byte(a.Content), &expected)
	require.NoError(s.t, err)

	actual, err := s.stateStore.ClassFeesTable().Get(s.ctx)
	require.NoError(s.t, err)

	for i, fee := range expected.Fees {
		require.Equal(s.t, fee.Amount, actual.Fees[i].Amount)
		require.Equal(s.t, fee.Denom, actual.Fees[i].Denom)
	}
}

func (s *updateClassFeesSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *updateClassFeesSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *updateClassFeesSuite) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}