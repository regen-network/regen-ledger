package tests

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmodules "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/fixture"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type bridgeSuite struct {
	t               gocuke.TestingT
	fixture         testutil.Fixture
	ctx             context.Context
	sdkCtx          sdk.Context
	ecocreditServer ecocreditServer
	err             error
}

type ecocreditServer struct {
	core.MsgClient
	core.QueryClient
}

func TestBridgeIntegration(t *testing.T) {
	gocuke.NewRunner(t, &bridgeSuite{}).Path("./features/bridge.feature").Run()
}

func (s *bridgeSuite) Before(t gocuke.TestingT) {
	s.t = t

	ff := fixture.NewFixtureFactory(t, 2)
	ff.SetModules([]sdkmodules.AppModule{
		NewEcocreditModule(ff),
	})

	s.fixture = ff.Setup()
	s.ctx = s.fixture.Context()
	s.sdkCtx = sdk.UnwrapSDKContext(s.ctx)

	s.ecocreditServer = ecocreditServer{
		MsgClient:   core.NewMsgClient(s.fixture.TxConn()),
		QueryClient: core.NewQueryClient(s.fixture.QueryConn()),
	}
}

func (s *bridgeSuite) EcocreditState(a gocuke.DocString) {
	_, err := s.fixture.InitGenesis(s.sdkCtx, map[string]json.RawMessage{
		ecocredit.ModuleName: json.RawMessage(a.Content),
	})
	require.NoError(s.t, err)
}

func (s *bridgeSuite) BridgeServiceCallsBridgeReceiveWithMessage(a gocuke.DocString) {
	var msg core.MsgBridgeReceive
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	// reset context events
	s.ctx = s.fixture.Context()
	s.sdkCtx = sdk.UnwrapSDKContext(s.ctx)

	_, s.err = s.ecocreditServer.BridgeReceive(s.ctx, &msg)
}

func (s *bridgeSuite) RecipientCallsBridgeWithMessage(a gocuke.DocString) {
	var msg core.MsgBridge
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	// reset context events
	s.ctx = s.fixture.Context()
	s.sdkCtx = sdk.UnwrapSDKContext(s.ctx)

	_, s.err = s.ecocreditServer.Bridge(s.ctx, &msg)
}

func (s *bridgeSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *bridgeSuite) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *bridgeSuite) ExpectTotalCreditBatches(a string) {
	expected, err := strconv.ParseUint(a, 10, 64)
	require.NoError(s.t, err)

	res, err := s.ecocreditServer.Batches(s.ctx, &core.QueryBatchesRequest{
		Pagination: &query.PageRequest{CountTotal: true},
	})
	require.NoError(s.t, err)
	require.Equal(s.t, expected, res.Pagination.Total)
}

func (s *bridgeSuite) ExpectTotalProjects(a string) {
	expected, err := strconv.ParseUint(a, 10, 64)
	require.NoError(s.t, err)

	res, err := s.ecocreditServer.Projects(s.ctx, &core.QueryProjectsRequest{
		Pagination: &query.PageRequest{CountTotal: true},
	})
	require.NoError(s.t, err)
	require.Equal(s.t, expected, res.Pagination.Total)
}

func (s *bridgeSuite) ExpectProjectWithProperties(a gocuke.DocString) {
	var expected core.Project
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	req := &core.QueryProjectRequest{ProjectId: expected.Id}
	project, err := s.ecocreditServer.Project(s.ctx, req)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.ReferenceId, project.Project.ReferenceId)
	require.Equal(s.t, expected.Metadata, project.Project.Metadata)
	require.Equal(s.t, expected.Jurisdiction, project.Project.Jurisdiction)
}

func (s *bridgeSuite) ExpectCreditBatchWithProperties(a gocuke.DocString) {
	var expected core.Batch
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	req := &core.QueryBatchRequest{BatchDenom: expected.Denom}
	project, err := s.ecocreditServer.Batch(s.ctx, req)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Metadata, project.Batch.Metadata)
	require.Equal(s.t, expected.StartDate, project.Batch.StartDate)
	require.Equal(s.t, expected.EndDate, project.Batch.EndDate)
	require.Equal(s.t, expected.Open, project.Batch.Open)
}

func (s *bridgeSuite) ExpectBatchSupplyWithBatchDenom(a string, b gocuke.DocString) {
	expected := &api.BatchSupply{}
	err := jsonpb.UnmarshalString(b.Content, expected)
	require.NoError(s.t, err)

	res, err := s.ecocreditServer.Supply(s.ctx, &core.QuerySupplyRequest{
		BatchDenom: a,
	})
	require.NoError(s.t, err)

	require.Equal(s.t, expected.TradableAmount, res.TradableAmount)
	require.Equal(s.t, expected.RetiredAmount, res.RetiredAmount)
	require.Equal(s.t, expected.CancelledAmount, res.CancelledAmount)
}

func (s *bridgeSuite) ExpectBatchBalanceWithAddressAndBatchDenom(a, b string, c gocuke.DocString) {
	expected := &api.BatchBalance{}
	err := jsonpb.UnmarshalString(c.Content, expected)
	require.NoError(s.t, err)

	res, err := s.ecocreditServer.Balance(s.ctx, &core.QueryBalanceRequest{
		Address:    a,
		BatchDenom: b,
	})
	require.NoError(s.t, err)

	require.Equal(s.t, expected.TradableAmount, res.Balance.TradableAmount)
	require.Equal(s.t, expected.RetiredAmount, res.Balance.RetiredAmount)
	require.Equal(s.t, expected.EscrowedAmount, res.Balance.EscrowedAmount)
}

func (s *bridgeSuite) ExpectEventBridgeReceiveWithValues(a gocuke.DocString) {
	var expected core.EventBridgeReceive
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&expected, sdkEvent)
	require.NoError(s.t, err)
}

func (s *bridgeSuite) ExpectEventBridgeWithValues(a gocuke.DocString) {
	var expected core.EventBridge
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&expected, sdkEvent)
	require.NoError(s.t, err)
}
