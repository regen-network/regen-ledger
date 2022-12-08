//nolint:revive,stylecheck
package tests

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmodules "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/query"

	baseapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	"github.com/regen-network/regen-ledger/types/v2/testutil/fixture"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
)

type marketSuite struct {
	t               gocuke.TestingT
	fixture         fixture.Fixture
	ctx             context.Context
	sdkCtx          sdk.Context
	ecocreditServer ecocreditServer
	marketServer    marketServer
	err             error
}

func TestMarketIntegration(t *testing.T) {
	gocuke.NewRunner(t, &marketSuite{}).Path("./features/market.feature").Run()
}

func (s *marketSuite) Before(t gocuke.TestingT) {
	s.t = t

	ff := fixture.NewFixtureFactory(t, 2)
	ff.SetModules([]sdkmodules.AppModule{
		NewEcocreditModule(ff),
	})

	s.fixture = ff.Setup()
	s.ctx = s.fixture.Context()
	s.sdkCtx = sdk.UnwrapSDKContext(s.ctx)

	s.ecocreditServer = ecocreditServer{
		MsgClient:   basetypes.NewMsgClient(s.fixture.TxConn()),
		QueryClient: basetypes.NewQueryClient(s.fixture.QueryConn()),
	}

	s.marketServer = marketServer{
		MsgClient:   markettypes.NewMsgClient(s.fixture.TxConn()),
		QueryClient: markettypes.NewQueryClient(s.fixture.QueryConn()),
	}
}

func (s *marketSuite) EcocreditState(a gocuke.DocString) {
	_, err := s.fixture.InitGenesis(s.sdkCtx, map[string]json.RawMessage{
		ecocredit.ModuleName: json.RawMessage(a.Content),
	})
	require.NoError(s.t, err)
}

func (s *marketSuite) AliceCreatesSellOrderWithMessage(a gocuke.DocString) {
	var msg markettypes.MsgSell
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	// reset context events
	s.ctx = s.fixture.Context()
	s.sdkCtx = sdk.UnwrapSDKContext(s.ctx)

	_, s.err = s.marketServer.Sell(s.ctx, &msg)
}

func (s *marketSuite) BobBuysCreditsWithMessage(a gocuke.DocString) {
	var msg markettypes.MsgBuyDirect
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	// reset context events
	s.ctx = s.fixture.Context()
	s.sdkCtx = sdk.UnwrapSDKContext(s.ctx)

	_, s.err = s.marketServer.BuyDirect(s.ctx, &msg)
}

func (s *marketSuite) AliceUpdatesSellOrderWithMessage(a gocuke.DocString) {
	var msg markettypes.MsgUpdateSellOrders
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	// reset context events
	s.ctx = s.fixture.Context()
	s.sdkCtx = sdk.UnwrapSDKContext(s.ctx)

	_, s.err = s.marketServer.UpdateSellOrders(s.ctx, &msg)
}

func (s *marketSuite) AliceCancelsSellOrderWithMessage(a gocuke.DocString) {
	var msg markettypes.MsgCancelSellOrder
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	// reset context events
	s.ctx = s.fixture.Context()
	s.sdkCtx = sdk.UnwrapSDKContext(s.ctx)

	_, s.err = s.marketServer.CancelSellOrder(s.ctx, &msg)
}

func (s *marketSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *marketSuite) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *marketSuite) ExpectEventSell(a gocuke.DocString) {
	var expected markettypes.EventSell
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&expected, sdkEvent)
	require.NoError(s.t, err)
}

func (s *marketSuite) ExpectEventBuy(a gocuke.DocString) {
	var expected markettypes.EventBuyDirect
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&expected, sdkEvent)
	require.NoError(s.t, err)
}

func (s *marketSuite) ExpectEventUpdate(a gocuke.DocString) {
	var expected markettypes.EventUpdateSellOrder
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&expected, sdkEvent)
	require.NoError(s.t, err)
}

func (s *marketSuite) ExpectEventCancel(a gocuke.DocString) {
	var expected markettypes.EventCancelSellOrder
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&expected, sdkEvent)
	require.NoError(s.t, err)
}

func (s *marketSuite) ExpectTotalSellOrders(a string) {
	expected, err := strconv.ParseUint(a, 10, 64)
	require.NoError(s.t, err)

	res, err := s.marketServer.SellOrders(s.ctx, &markettypes.QuerySellOrdersRequest{
		Pagination: &query.PageRequest{CountTotal: true},
	})
	require.NoError(s.t, err)
	require.Equal(s.t, expected, res.Pagination.Total)
}

func (s *marketSuite) ExpectQuerySellOrderWithId(a string, b gocuke.DocString) {
	id, err := strconv.ParseUint(a, 10, 64)
	require.NoError(s.t, err)

	var expected markettypes.QuerySellOrderResponse
	err = jsonpb.UnmarshalString(b.Content, &expected)
	require.NoError(s.t, err)

	req := &markettypes.QuerySellOrderRequest{SellOrderId: id}
	res, err := s.marketServer.SellOrder(s.ctx, req)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.SellOrder.Id, res.SellOrder.Id)
	require.Equal(s.t, expected.SellOrder.Seller, res.SellOrder.Seller)
	require.Equal(s.t, expected.SellOrder.BatchDenom, res.SellOrder.BatchDenom)
	require.Equal(s.t, expected.SellOrder.Quantity, res.SellOrder.Quantity)
	require.Equal(s.t, expected.SellOrder.AskDenom, res.SellOrder.AskDenom)
	require.Equal(s.t, expected.SellOrder.AskAmount, res.SellOrder.AskAmount)
	require.Equal(s.t, expected.SellOrder.DisableAutoRetire, res.SellOrder.DisableAutoRetire)
	require.Equal(s.t, expected.SellOrder.Expiration, res.SellOrder.Expiration)
}

func (s *marketSuite) ExpectQueryBalanceWithAddressAndBatchDenom(a, b string, c gocuke.DocString) {
	expected := &baseapi.QueryBalanceResponse{}
	err := jsonpb.UnmarshalString(c.Content, expected)
	require.NoError(s.t, err)

	res, err := s.ecocreditServer.Balance(s.ctx, &basetypes.QueryBalanceRequest{
		Address:    a,
		BatchDenom: b,
	})
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Balance.TradableAmount, res.Balance.TradableAmount)
	require.Equal(s.t, expected.Balance.RetiredAmount, res.Balance.RetiredAmount)
	require.Equal(s.t, expected.Balance.EscrowedAmount, res.Balance.EscrowedAmount)
}

func (s *marketSuite) ExpectQuerySupplyWithBatchDenom(a string, b gocuke.DocString) {
	expected := &baseapi.QuerySupplyResponse{}
	err := jsonpb.UnmarshalString(b.Content, expected)
	require.NoError(s.t, err)

	res, err := s.ecocreditServer.Supply(s.ctx, &basetypes.QuerySupplyRequest{
		BatchDenom: a,
	})
	require.NoError(s.t, err)

	require.Equal(s.t, expected.TradableAmount, res.TradableAmount)
	require.Equal(s.t, expected.RetiredAmount, res.RetiredAmount)
	require.Equal(s.t, expected.CancelledAmount, res.CancelledAmount)
}
