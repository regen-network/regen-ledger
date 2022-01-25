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

	before := s.startTime.Add(time.Hour * 24)
	expiration := s.startTime.Add(time.Hour * 48)
	expiration2 := s.startTime.Add(time.Hour * 50)
	after := s.startTime.Add(time.Hour * 72)
	owner := s.owner.String()

	s1 := s.createSellOrder(&expiration, owner)
	s2 := s.createSellOrder(&expiration2, owner)
	s3 := s.createSellOrder(nil, owner)

	b1 := s.createBuyOrder(&expiration, owner)
	b2 := s.createBuyOrder(&expiration2, owner)
	b3 := s.createBuyOrder(nil, owner)

	/*
	 * TEST1: set blockchain before the expiration time.
	 * prunning shouldn't remove anything */

	s.sctx = s.sctx.WithBlockTime(before)
	s.server.PruneOrders(s.sctx)

	require.True(s.server.sellOrderTable.Has(s.sctx, s1))
	require.True(s.server.sellOrderTable.Has(s.sctx, s2))
	require.True(s.server.sellOrderTable.Has(s.sctx, s3))
	require.True(s.server.buyOrderTable.Has(s.sctx, b1))
	require.True(s.server.buyOrderTable.Has(s.sctx, b2))
	require.True(s.server.buyOrderTable.Has(s.sctx, b3))

	/*
	 * TEST2: set blockchain at expiration time.
	 * prunning remove orders <= expiration */

	s.sctx = s.sctx.WithBlockTime(expiration)
	s.server.PruneOrders(s.sctx)

	require.False(s.server.sellOrderTable.Has(s.sctx, s1))
	require.True(s.server.sellOrderTable.Has(s.sctx, s2))
	require.True(s.server.sellOrderTable.Has(s.sctx, s3))
	require.False(s.server.buyOrderTable.Has(s.sctx, b1))
	require.True(s.server.buyOrderTable.Has(s.sctx, b2))
	require.True(s.server.buyOrderTable.Has(s.sctx, b3))

	s.sctx = s.sctx.WithBlockTime(after)

	/*
	 * TEST3: set blockchain at "after".
	 * prunning remove all orders those with no expiration */

	s.sctx = s.sctx.WithBlockTime(after)
	s.server.PruneOrders(s.sctx)

	require.False(s.server.sellOrderTable.Has(s.sctx, s2))
	require.True(s.server.sellOrderTable.Has(s.sctx, s3))
	require.False(s.server.buyOrderTable.Has(s.sctx, b2))
	require.True(s.server.buyOrderTable.Has(s.sctx, b3))
}

func (s *USuite) createSellOrder(expiration *time.Time, owner string) uint64 {
	o := &ecocredit.MsgSell_Order{
		BatchDenom:        "C",
		Quantity:          "1",
		AskPrice:          &sdk.Coin{"token", sdk.NewInt(1)},
		DisableAutoRetire: false,
		Expiration:        expiration,
	}
	id, err := s.server.createSellOrder(types.Context{s.sctx}, owner, o)
	s.Require().NoError(err)
	return id
}

func (s *USuite) createBuyOrder(expiration *time.Time, owner string) uint64 {
	o := &ecocredit.MsgBuy_Order{
		Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: 1}},
		Quantity:          "1",
		BidPrice:          &sdk.Coin{"token", sdk.NewInt(1)},
		DisableAutoRetire: false,
		Expiration:        expiration,
	}
	id, err := s.server.createBuyOrder(types.Context{s.sctx}, owner, o)
	s.Require().NoError(err)
	return id
}
