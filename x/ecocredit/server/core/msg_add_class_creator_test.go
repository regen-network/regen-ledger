package core

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type addClassCreator struct {
	*baseSuite
	err error
}

func TestAddClassCreator(t *testing.T) {
	gocuke.NewRunner(t, &addClassCreator{}).Path("./features/msg_add_class_creators.feature").Run()
}

func (s *addClassCreator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *addClassCreator) ClassCreatorWithProperties(a gocuke.DocString) {
	var msg *core.MsgAddClassCreator

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	require.NoError(s.t, err)

	err = s.stateStore.AllowedClassCreatorTable().Save(s.ctx, &ecocreditv1.AllowedClassCreator{
		Address: creatorAddr,
	})
	require.NoError(s.t, err)

}

func (s *addClassCreator) AliceAttemptsToAddClassCreatorWithProperties(a gocuke.DocString) {
	var msg *core.MsgAddClassCreator

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.AddClassCreator(s.ctx, msg)
}

func (s *addClassCreator) ExpectClassCreatorWithProperties(a gocuke.DocString) {
	var msg *core.MsgAddClassCreator

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	params, err := s.k.Params(s.ctx, &core.QueryParamsRequest{})
	require.NoError(s.t, err)

	require.Equal(s.t, len(params.Params.AllowedClassCreators), 1)
	require.Equal(s.t, msg.Creator, params.Params.AllowedClassCreators[0])
}

func (s *addClassCreator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *addClassCreator) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *addClassCreator) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}
