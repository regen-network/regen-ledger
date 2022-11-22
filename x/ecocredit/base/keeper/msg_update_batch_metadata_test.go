//nolint:revive,stylecheck
package keeper

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/libs/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

type updateBatchMetadata struct {
	*baseSuite
	alice            sdk.AccAddress
	bob              sdk.AccAddress
	creditTypeAbbrev string
	classKey         uint64
	projectKey       uint64
	batchDenom       string
	res              *types.MsgUpdateBatchMetadataResponse
	err              error
}

func TestUpdateBatchMetadata(t *testing.T) {
	gocuke.NewRunner(t, &updateBatchMetadata{}).Path("./features/msg_update_batch_metadata.feature").Run()
}

func (s *updateBatchMetadata) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.bob = s.addr2
}

func (s *updateBatchMetadata) ACreditTypeWithAbbreviation(a string) {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)
}

func (s *updateBatchMetadata) ACreditClassWithIdAndIssuerAlice(a string) {
	cKey, err := s.k.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id:               a,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	err = s.k.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
		ClassKey: cKey,
		Issuer:   s.alice,
	})
	require.NoError(s.t, err)

	s.classKey = cKey
}

func (s *updateBatchMetadata) AProjectWithId(a string) {
	pKey, err := s.k.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id:       a,
		ClassKey: s.classKey,
	})
	require.NoError(s.t, err)

	s.projectKey = pKey
}

func (s *updateBatchMetadata) ACreditBatchWithBatchDenomAndIssuerAlice(a string) {
	projectID := base.GetProjectIDFromBatchDenom(a)

	project, err := s.k.stateStore.ProjectTable().GetById(s.ctx, projectID)
	require.NoError(s.t, err)

	_, err = s.k.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		ProjectKey: project.Key,
		Issuer:     s.alice,
		Denom:      a,
		Open:       true, // true unless specified
	})
	require.NoError(s.t, err)

	s.batchDenom = a
}

func (s *updateBatchMetadata) ACreditBatchWithBatchDenomIssuerAliceAndOpen(a, b string) {
	projectID := base.GetProjectIDFromBatchDenom(a)

	project, err := s.k.stateStore.ProjectTable().GetById(s.ctx, projectID)
	require.NoError(s.t, err)

	open, err := strconv.ParseBool(b)
	require.NoError(s.t, err)

	_, err = s.k.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		ProjectKey: project.Key,
		Issuer:     s.alice,
		Denom:      a,
		Open:       open,
	})
	require.NoError(s.t, err)

	s.batchDenom = a
}

func (s *updateBatchMetadata) AliceAttemptsToUpdateBatchMetadataWithBatchDenom(a string) {
	s.res, s.err = s.k.UpdateBatchMetadata(s.ctx, &types.MsgUpdateBatchMetadata{
		Issuer:     s.alice.String(),
		BatchDenom: a,
	})
}

func (s *updateBatchMetadata) BobAttemptsToUpdateBatchMetadataWithBatchDenom(a string) {
	s.res, s.err = s.k.UpdateBatchMetadata(s.ctx, &types.MsgUpdateBatchMetadata{
		Issuer:     s.bob.String(),
		BatchDenom: a,
	})
}

func (s *updateBatchMetadata) AliceAttemptsToUpdateBatchMetadataWithBatchDenomAndNewMetadata(a string, b gocuke.DocString) {
	s.res, s.err = s.k.UpdateBatchMetadata(s.ctx, &types.MsgUpdateBatchMetadata{
		Issuer:      s.alice.String(),
		BatchDenom:  a,
		NewMetadata: b.Content,
	})
}

func (s *updateBatchMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *updateBatchMetadata) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *updateBatchMetadata) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *updateBatchMetadata) ExpectCreditBatchWithBatchDenomAndMetadata(a string, b gocuke.DocString) {
	batch, err := s.stateStore.BatchTable().GetByDenom(s.ctx, a)
	require.NoError(s.t, err)

	require.Equal(s.t, b.Content, batch.Metadata)
}

func (s *updateBatchMetadata) AliceUpdatesTheBatchMetadata() {
	newMetadata := rand.Str(5)
	_, err := s.k.UpdateBatchMetadata(s.ctx, &types.MsgUpdateBatchMetadata{
		Issuer:      s.alice.String(),
		BatchDenom:  s.batchDenom,
		NewMetadata: newMetadata,
	})
	require.NoError(s.t, err)
}

func (s *updateBatchMetadata) ExpectEventWithProperties(a gocuke.DocString) {
	var event types.EventUpdateBatchMetadata
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	events := s.sdkCtx.EventManager().Events()
	lastEvent := events[len(events)-1]

	require.Equal(s.t, proto.MessageName(&event), lastEvent.Type)
	require.Len(s.t, lastEvent.Attributes, 1)

	batchDenom := strings.Trim(string(lastEvent.Attributes[0].Value), `"`)
	require.Equal(s.t, event.BatchDenom, batchDenom)
}
