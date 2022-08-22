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

type removeClassClassCreators struct {
	*baseSuite
	err error
}

func TestRemoveClassCreators(t *testing.T) {
	gocuke.NewRunner(t, &removeClassClassCreators{}).Path("./features/msg_remove_class_creators.feature").Run()
}

func (s *removeClassClassCreators) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *removeClassClassCreators) ClassCreatorsWithProperties(a gocuke.DocString) {
	var msg *core.MsgAddClassCreators

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	for _, creator := range msg.Creators {
		creatorAddr, err := sdk.AccAddressFromBech32(creator)
		require.NoError(s.t, err)

		err = s.stateStore.AllowedClassCreatorTable().Save(s.ctx, &ecocreditv1.AllowedClassCreator{
			Address: creatorAddr,
		})
		require.NoError(s.t, err)
	}

}

func (s *removeClassClassCreators) AliceAttemptsToRemoveClassCreatorsWithProperties(a gocuke.DocString) {
	var msg *core.MsgRemoveClassCreators

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.RemoveClassCreators(s.ctx, msg)
}

func (s *removeClassClassCreators) ExpectClassCreatorsWithProperties(a gocuke.DocString) {
	var msg *core.MsgRemoveClassCreators

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	found := 0

	params, err := s.k.Params(s.ctx, &core.QueryParamsRequest{})
	require.NoError(s.t, err)

	require.Equal(s.t, len(msg.Creators), len(params.Params.AllowedClassCreators))
	for _, creator := range params.Params.AllowedClassCreators {
		for _, creator1 := range msg.Creators {
			if creator == creator1 {
				found++
			}
		}
	}
	require.Equal(s.t, len(msg.Creators), found)
}

func (s *removeClassClassCreators) ExpectClassCreatorsListToBeEmpty() {
	params, err := s.k.Params(s.ctx, &core.QueryParamsRequest{})
	require.NoError(s.t, err)

	require.Zero(s.t, len(params.Params.AllowedClassCreators))
}

func (s *removeClassClassCreators) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *removeClassClassCreators) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *removeClassClassCreators) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}
