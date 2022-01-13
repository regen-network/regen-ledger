package server

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/mem"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/stretchr/testify/suite"
)

func TestUSuite(t *testing.T) {
	s := new(USuite)
	s.sctx, s.storeKey = setupStore(t)

	suite.Run(t, s)
}

type USuite struct {
	suite.Suite
	storeKey  *sdk.KVStoreKey
	cdc       codec.Codec
	db        *mem.Store
	server    serverImpl
	sctx      sdk.Context
	startTime time.Time
	owner     sdk.AccAddress
}

func (s *USuite) SetupSuite() {
	registry := codectypes.NewInterfaceRegistry()
	ecocredit.RegisterTypes(registry)
	s.cdc = codec.NewProtoCodec(registry)
	s.server = newServer(s.storeKey, paramtypes.Subspace{} /* paramSpace types.Subspace */, nil, nil, s.cdc)

	s.startTime = time.Unix(0, 0)
	s.owner = sdk.AccAddress([]byte("owner"))
}

func (s *USuite) SetupTest() {
	s.sctx = s.sctx.WithBlockTime(s.startTime)
}

func (s *USuite) TestExpiration() {
	require := s.Require()

	t1 := s.startTime.Add(time.Hour * 24)
	t2 := s.startTime.Add(time.Hour * 72)
	texp := s.startTime.Add(time.Hour * 48)

	// create sell order with expiration (order id: 1)
	sell1 := &ecocredit.MsgSell_Order{
		BatchDenom:        "C",
		Quantity:          "1",
		AskPrice:          &sdk.Coin{"token", sdk.NewInt(1)},
		DisableAutoRetire: false,
		Expiration:        &texp,
	}
	_, err := s.server.createSellOrder(types.Context{s.sctx}, s.owner.String(), sell1)
	require.NoError(err)

	// create sell order without expiration (order id: 2)
	sell2 := &ecocredit.MsgSell_Order{
		BatchDenom:        "C",
		Quantity:          "1",
		AskPrice:          &sdk.Coin{"token", sdk.NewInt(1)},
		DisableAutoRetire: false,
	}
	_, err = s.server.createSellOrder(types.Context{s.sctx}, s.owner.String(), sell2)
	require.NoError(err)

	// create buy order with expiration (order id: 1)
	buy1 := &ecocredit.MsgBuy_Order{
		Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: 1}},
		Quantity:          "1",
		BidPrice:          &sdk.Coin{"token", sdk.NewInt(1)},
		DisableAutoRetire: false,
		Expiration:        &texp,
	}
	_, err = s.server.createBuyOrder(types.Context{s.sctx}, s.owner.String(), buy1)
	require.NoError(err)

	// create buy order without expiration (order id: 2)
	buy2 := &ecocredit.MsgBuy_Order{
		Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: 1}},
		Quantity:          "1",
		BidPrice:          &sdk.Coin{"token", sdk.NewInt(1)},
		DisableAutoRetire: false,
	}
	_, err = s.server.createBuyOrder(types.Context{s.sctx}, s.owner.String(), buy2)
	require.NoError(err)

	// set block time before expiration
	s.sctx = s.sctx.WithBlockTime(t1)

	s.server.PruneOrders(s.sctx)

	require.True(s.server.sellOrderTable.Has(s.sctx, 1))
	require.True(s.server.sellOrderTable.Has(s.sctx, 2))
	require.True(s.server.buyOrderTable.Has(s.sctx, 1))
	require.True(s.server.buyOrderTable.Has(s.sctx, 2))

	// set block time after expiration
	s.sctx = s.sctx.WithBlockTime(t2)

	s.server.PruneOrders(s.sctx)

	require.False(s.server.sellOrderTable.Has(s.sctx, 1))
	require.True(s.server.sellOrderTable.Has(s.sctx, 2))
	require.False(s.server.buyOrderTable.Has(s.sctx, 1))
	require.True(s.server.buyOrderTable.Has(s.sctx, 2))
}
