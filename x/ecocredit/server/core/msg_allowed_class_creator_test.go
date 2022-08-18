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

type allowedClassCreator struct {
	*baseSuite
	err error
}

func TestAllowedClassCreator(t *testing.T) {
	gocuke.NewRunner(t, &allowedClassCreator{}).Path("./features/msg_allowed_class_creator.feature").Run()
}

func (s *allowedClassCreator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *allowedClassCreator) ClassCreatorsWithProperties(a gocuke.DocString) {
	var msg *core.MsgAllowedClassCreator

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

func (s *allowedClassCreator) AliceAttemptsToAddClassCreatorsWithProperties(a gocuke.DocString) {
	var msg *core.MsgAllowedClassCreator

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.AllowedClassCreator(s.ctx, msg)
}

func (s *allowedClassCreator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *allowedClassCreator) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *allowedClassCreator) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}
