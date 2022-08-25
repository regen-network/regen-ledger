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

type removeClassClassCreator struct {
	*baseSuite
	err error
}

type classCreators struct {
	Creators []string `json:"creators"`
}

func TestRemoveClassCreators(t *testing.T) {
	gocuke.NewRunner(t, &removeClassClassCreator{}).Path("./features/msg_remove_class_creator.feature").Run()
}

func (s *removeClassClassCreator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *removeClassClassCreator) ClassCreatorsWithProperties(a gocuke.DocString) {
	var creators classCreators

	err := json.Unmarshal([]byte(a.Content), &creators)
	require.NoError(s.t, err)

	for _, creator := range creators.Creators {
		creatorAddr, err := sdk.AccAddressFromBech32(creator)
		require.NoError(s.t, err)

		err = s.stateStore.AllowedClassCreatorTable().Insert(s.ctx, &ecocreditv1.AllowedClassCreator{
			Address: creatorAddr,
		})
		require.NoError(s.t, err)
	}
}

func (s *removeClassClassCreator) AliceAttemptsToRemoveAClassCreatorWithProperties(a gocuke.DocString) {
	var msg *core.MsgRemoveClassCreator

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.RemoveClassCreator(s.ctx, msg)
}

func (s *removeClassClassCreator) ExpectClassCreatorsWithProperties(a gocuke.DocString) {
	var creators classCreators

	err := json.Unmarshal([]byte(a.Content), &creators)
	require.NoError(s.t, err)

	itr, err := s.stateStore.AllowedClassCreatorTable().List(s.ctx, ecocreditv1.AllowedClassCreatorPrimaryKey{})
	require.NoError(s.t, err)

	found := 0
	for itr.Next() {
		val, err := itr.Value()
		require.NoError(s.t, err)
		for _, creator := range creators.Creators {
			if creator == sdk.AccAddress(val.Address).String() {
				found++
			}
		}

	}
	itr.Close()

	require.Equal(s.t, len(creators.Creators), found)
}

func (s *removeClassClassCreator) ExpectClassCreatorsListToBeEmpty() {
	itr, err := s.stateStore.AllowedClassCreatorTable().List(s.ctx, ecocreditv1.AllowedClassCreatorPrimaryKey{})
	require.NoError(s.t, err)

	ok := itr.Next()
	require.False(s.t, ok)
}

func (s *removeClassClassCreator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *removeClassClassCreator) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *removeClassClassCreator) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}
