package keeper

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

type updateClassFeeSuite struct {
	*baseSuite
	msg *types.MsgUpdateClassFee
	err error
}

func TestUpdateClassFee(t *testing.T) {
	gocuke.NewRunner(t, &updateClassFeeSuite{}).Path("./features/msg_update_class_fee.feature").Run()
}

func (s *updateClassFeeSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *updateClassFeeSuite) TheMessage(a gocuke.DocString) {
	s.msg = &types.MsgUpdateClassFee{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *updateClassFeeSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *updateClassFeeSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *updateClassFeeSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *updateClassFeeSuite) AliceAttemptsToUpdateClassFeeWithProperties(a gocuke.DocString) {
	var msg *types.MsgUpdateClassFee

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.UpdateClassFee(s.ctx, msg)
}

func (s *updateClassFeeSuite) ExpectClassFeeWithProperties(a gocuke.DocString) {
	var expected *api.ClassFee
	err := json.Unmarshal([]byte(a.Content), &expected)
	require.NoError(s.t, err)

	actual, err := s.stateStore.ClassFeeTable().Get(s.ctx)
	require.NoError(s.t, err)

	if expected.Fee.Denom != "" {
		require.Equal(s.t, expected.Fee.Amount, actual.Fee.Amount)
		require.Equal(s.t, expected.Fee.Denom, actual.Fee.Denom)
	}
}

func (s *updateClassFeeSuite) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}
