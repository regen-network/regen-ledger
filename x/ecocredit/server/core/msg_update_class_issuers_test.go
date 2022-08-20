//nolint:revive,stylecheck
package core

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type updateClassIssuers struct {
	*baseSuite
	alice sdk.AccAddress
	bob   sdk.AccAddress
	res   *core.MsgUpdateClassIssuersResponse
	err   error
}

func TestUpdateClassIssuers(t *testing.T) {
	gocuke.NewRunner(t, &updateClassIssuers{}).Path("./features/msg_update_class_issuers.feature").Run()
}

func (s *updateClassIssuers) Before(t gocuke.TestingT) {
	// TODO: move to init function in the root directory of the module #1243
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")

	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.bob = s.addr2
}

func (s *updateClassIssuers) ACreditTypeWithAbbreviation(a string) {
	// TODO: Save for now but credit type should not exist prior to unit test #893
	err := s.k.stateStore.CreditTypeTable().Save(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)
}

func (s *updateClassIssuers) ACreditClassWithClassIdAndAdminAlice(a string) {
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassID(a)

	err := s.k.stateStore.ClassTable().Insert(s.ctx, &api.Class{
		Id:               a,
		Admin:            s.alice,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *updateClassIssuers) ACreditClassWithClassIdAdminAliceAndIssuers(a string, b gocuke.DocString) {
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassID(a)

	cKey, err := s.k.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id:               a,
		Admin:            s.alice,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	var issuers []string
	err = json.Unmarshal([]byte(b.Content), &issuers)
	require.NoError(s.t, err)

	for _, issuer := range issuers {
		bz, err := sdk.AccAddressFromBech32(issuer)
		require.NoError(s.t, err)

		err = s.k.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
			ClassKey: cKey,
			Issuer:   bz,
		})
		require.NoError(s.t, err)
	}
}

func (s *updateClassIssuers) AliceAttemptsToUpdateClassIssuersWithClassId(a string) {
	s.res, s.err = s.k.UpdateClassIssuers(s.ctx, &core.MsgUpdateClassIssuers{
		Admin:   s.alice.String(),
		ClassId: a,
	})
}

func (s *updateClassIssuers) BobAttemptsToUpdateClassIssuersWithClassId(a string) {
	s.res, s.err = s.k.UpdateClassIssuers(s.ctx, &core.MsgUpdateClassIssuers{
		Admin:   s.bob.String(),
		ClassId: a,
	})
}

func (s *updateClassIssuers) AliceAttemptsToUpdateClassIssuersWithClassIdAndAddIssuers(a string, b gocuke.DocString) {
	var addIssuers []string
	err := json.Unmarshal([]byte(b.Content), &addIssuers)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateClassIssuers(s.ctx, &core.MsgUpdateClassIssuers{
		Admin:      s.alice.String(),
		ClassId:    a,
		AddIssuers: addIssuers,
	})
}

func (s *updateClassIssuers) AliceAttemptsToUpdateClassIssuersWithClassIdAndRemoveIssuers(a string, b gocuke.DocString) {
	var removeIssuers []string
	err := json.Unmarshal([]byte(b.Content), &removeIssuers)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateClassIssuers(s.ctx, &core.MsgUpdateClassIssuers{
		Admin:         s.alice.String(),
		ClassId:       a,
		RemoveIssuers: removeIssuers,
	})
}

func (s *updateClassIssuers) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *updateClassIssuers) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *updateClassIssuers) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *updateClassIssuers) ExpectCreditClassWithClassIdAndIssuers(a string, b gocuke.DocString) {
	class, err := s.stateStore.ClassTable().GetById(s.ctx, a)
	require.NoError(s.t, err)

	var issuers []string
	err = json.Unmarshal([]byte(b.Content), &issuers)
	require.NoError(s.t, err)

	for _, issuer := range issuers {
		bx, err := sdk.AccAddressFromBech32(issuer)
		require.NoError(s.t, err)

		found, err := s.stateStore.ClassIssuerTable().Has(s.ctx, class.Key, bx)
		require.NoError(s.t, err)

		require.True(s.t, found)
	}
}
