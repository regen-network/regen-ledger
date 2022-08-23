//nolint:revive,stylecheck
package core

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type updateClassMetadata struct {
	*baseSuite
	alice sdk.AccAddress
	bob   sdk.AccAddress
	res   *core.MsgUpdateClassMetadataResponse
	err   error
}

func TestUpdateClassMetadata(t *testing.T) {
	gocuke.NewRunner(t, &updateClassMetadata{}).Path("./features/msg_update_class_metadata.feature").Run()
}

func (s *updateClassMetadata) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.bob = s.addr2
}

func (s *updateClassMetadata) ACreditTypeWithAbbreviation(a string) {
	// TODO: Save for now but credit type should not exist prior to unit test #893
	err := s.k.stateStore.CreditTypeTable().Save(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)
}

func (s *updateClassMetadata) ACreditClassWithClassIdAndAdminAlice(a string) {
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassID(a)

	err := s.k.stateStore.ClassTable().Insert(s.ctx, &api.Class{
		Id:               a,
		Admin:            s.alice,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *updateClassMetadata) AliceAttemptsToUpdateClassMetadataWithClassId(a string) {
	s.res, s.err = s.k.UpdateClassMetadata(s.ctx, &core.MsgUpdateClassMetadata{
		Admin:   s.alice.String(),
		ClassId: a,
	})
}

func (s *updateClassMetadata) BobAttemptsToUpdateClassMetadataWithClassId(a string) {
	s.res, s.err = s.k.UpdateClassMetadata(s.ctx, &core.MsgUpdateClassMetadata{
		Admin:   s.bob.String(),
		ClassId: a,
	})
}

func (s *updateClassMetadata) AliceAttemptsToUpdateClassMetadataWithClassIdAndNewMetadata(a string, b gocuke.DocString) {
	s.res, s.err = s.k.UpdateClassMetadata(s.ctx, &core.MsgUpdateClassMetadata{
		Admin:       s.alice.String(),
		ClassId:     a,
		NewMetadata: b.Content,
	})
}

func (s *updateClassMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *updateClassMetadata) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *updateClassMetadata) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *updateClassMetadata) ExpectCreditClassWithClassIdAndMetadata(a string, b gocuke.DocString) {
	class, err := s.stateStore.ClassTable().GetById(s.ctx, a)
	require.NoError(s.t, err)

	require.Equal(s.t, b.Content, class.Metadata)
}
