//nolint:revive,stylecheck
package core

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type updateClassAdmin struct {
	*baseSuite
	alice sdk.AccAddress
	bob   sdk.AccAddress
	res   *core.MsgUpdateClassAdminResponse
	err   error
}

func TestUpdateClassAdmin(t *testing.T) {
	gocuke.NewRunner(t, &updateClassAdmin{}).Path("./features/msg_update_class_admin.feature").Run()
}

func (s *updateClassAdmin) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.bob = s.addr2
}

func (s *updateClassAdmin) ACreditTypeWithAbbreviation(a string) {
	// TODO: Save for now but credit type should not exist prior to unit test #893
	err := s.k.stateStore.CreditTypeTable().Save(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)
}

func (s *updateClassAdmin) ACreditClassWithClassIdAndAdminAlice(a string) {
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassID(a)

	err := s.k.stateStore.ClassTable().Insert(s.ctx, &api.Class{
		Id:               a,
		Admin:            s.alice,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *updateClassAdmin) AliceAttemptsToUpdateClassAdminWithClassId(a string) {
	s.res, s.err = s.k.UpdateClassAdmin(s.ctx, &core.MsgUpdateClassAdmin{
		Admin:    s.alice.String(),
		ClassId:  a,
		NewAdmin: s.bob.String(),
	})
}

func (s *updateClassAdmin) BobAttemptsToUpdateClassAdminWithClassId(a string) {
	s.res, s.err = s.k.UpdateClassAdmin(s.ctx, &core.MsgUpdateClassAdmin{
		Admin:    s.bob.String(),
		ClassId:  a,
		NewAdmin: s.bob.String(),
	})
}

func (s *updateClassAdmin) AliceAttemptsToUpdateClassAdminWithClassIdAndNewAdminBob(a string) {
	s.res, s.err = s.k.UpdateClassAdmin(s.ctx, &core.MsgUpdateClassAdmin{
		Admin:    s.alice.String(),
		ClassId:  a,
		NewAdmin: s.bob.String(),
	})
}

func (s *updateClassAdmin) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *updateClassAdmin) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *updateClassAdmin) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *updateClassAdmin) ExpectCreditClassWithClassIdAndAdminBob(a string) {
	class, err := s.stateStore.ClassTable().GetById(s.ctx, a)
	require.NoError(s.t, err)

	require.Equal(s.t, s.bob.Bytes(), class.Admin)
}

func (s *updateClassAdmin) ExpectEventWithProperties(a gocuke.DocString) {
	var event core.EventUpdateClassAdmin
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}
