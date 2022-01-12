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
	abci "github.com/tendermint/tendermint/abci/types"
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

	owner sdk.AccAddress
	acc1  sdk.AccAddress
}

func (s *USuite) SetupSuite() {
	registry := codectypes.NewInterfaceRegistry()
	ecocredit.RegisterTypes(registry)
	s.cdc = codec.NewProtoCodec(registry)
	s.server = newServer(s.storeKey, paramtypes.Subspace{} /* paramSpace types.Subspace */, nil, nil, s.cdc)

	s.startTime = time.Unix(0, 0)
	s.owner = sdk.AccAddress([]byte("owner"))
	s.acc1 = sdk.AccAddress([]byte("acc1"))
}

func (s *USuite) SetupTest() {
	s.sctx = s.sctx.WithBlockTime(s.startTime)
}

func (s *USuite) TestPrune1() {
	require := s.Require()
	t1 := s.startTime.Add(time.Hour * 1)
	texp := s.startTime.Add(time.Hour * 2)

	denom := "carbon1"
	o1 := &ecocredit.MsgSell_Order{denom, "1", &sdk.Coin{"atom", sdk.NewInt(1)}, false, &texp}
	_, err := s.server.createSellOrder(types.Context{s.sctx}, s.owner.String(), o1)
	require.NoError(err)

	s.sctx = s.sctx.WithBlockTime(t1)

	s.server.PruneOrders(s.sctx)

	// todo: verify

	// test cases:
	// * orders with no expire time are not prunned
	// * ...
}
