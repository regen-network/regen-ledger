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

type adddClassCreators struct {
	*baseSuite
	err error
}

func TestAddClassCreators(t *testing.T) {
	gocuke.NewRunner(t, &adddClassCreators{}).Path("./features/msg_add_class_creators.feature").Run()
}

func (s *adddClassCreators) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *adddClassCreators) ClassCreatorsWithProperties(a gocuke.DocString) {
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

func (s *adddClassCreators) AliceAttemptsToAddClassCreatorsWithProperties(a gocuke.DocString) {
	var msg *core.MsgAddClassCreators

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.AddClassCreators(s.ctx, msg)
}

func (s *adddClassCreators) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *adddClassCreators) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *adddClassCreators) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}
