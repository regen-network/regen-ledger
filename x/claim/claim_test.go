package claim_test

import (
	"bytes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/graph"
	"github.com/regen-network/regen-ledger/graph/binary"
	"github.com/regen-network/regen-ledger/graph/gen"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/claim"
	"github.com/regen-network/regen-ledger/x/data"
	schema_test "github.com/regen-network/regen-ledger/x/schema/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	schema_test.Harness
	dataKeeper data.Keeper
	Keeper     claim.Keeper
	Handler    sdk.Handler
}

func (s *Suite) SetupTest() {
	s.Setup()
	data.RegisterCodec(s.Cdc)
	claim.RegisterCodec(s.Cdc)
	dataKey := sdk.NewKVStoreKey("data")
	s.dataKeeper = data.NewKeeper(dataKey, s.Harness.Keeper, s.Cdc)
	claimKey := sdk.NewKVStoreKey("claim")
	s.Keeper = claim.NewKeeper(claimKey, s.dataKeeper, s.Cdc)
	s.Handler = claim.NewHandler(s.Keeper)
	s.Cms.MountStoreWithDB(dataKey, sdk.StoreTypeIAVL, s.Db)
	s.Cms.MountStoreWithDB(claimKey, sdk.StoreTypeIAVL, s.Db)
	_ = s.Cms.LoadLatestVersion()
	s.CreateSampleSchema()
}

func (s *Suite) randomData() types.DataAddress {
	x, ok := gen.Graph(s.Resolver).Sample()
	if !ok {
		panic("couldn't generate graph")
	}
	g := x.(graph.Graph)
	buf := new(bytes.Buffer)
	err := binary.SerializeGraph(s.Harness.Resolver, g, buf)
	s.Require().Nil(err)
	addr, err := s.dataKeeper.StoreGraph(s.Ctx, graph.Hash(g), buf.Bytes())
	s.Require().Nil(err)
	return addr
}

func (s *Suite) TestCreateClaim() {
	s.T().Logf("sign a claim")
	c := s.randomData()
	ev0 := s.randomData()
	ev1 := s.randomData()
	msg := claim.MsgSignClaim{Content: c, Evidence: []types.DataAddress{ev0, ev1}, Signers: []sdk.AccAddress{s.Addr1}}
	res := s.Handler(s.Ctx, msg)
	s.Require().Equal(sdk.CodeOK, res.Code)
	s.Require().Equal(string(res.Tags[0].Value), c.String())

	s.T().Logf("retrieve the signatures")
	sigs := s.Keeper.GetSigners(s.Ctx, c)
	s.Require().True(bytes.Equal(s.Addr1, sigs[0]))

	s.T().Logf("retrieve the evidence")
	ev := s.Keeper.GetEvidence(s.Ctx, c, s.Addr1)
	s.requireContainsData(ev, ev0)
	s.requireContainsData(ev, ev1)

	s.T().Logf("add more evidence and another signature")
	ev2 := s.randomData()
	err := s.Keeper.SignClaim(s.Ctx, c, []types.DataAddress{ev2}, []sdk.AccAddress{s.Addr1, s.Addr2})
	s.Require().Nil(err)

	s.T().Logf("retrieve the signatures")
	sigs = s.Keeper.GetSigners(s.Ctx, c)
	s.requireContainsAddr(sigs, s.Addr1)
	s.requireContainsAddr(sigs, s.Addr2)

	s.T().Logf("retrieve the evidence")
	ev = s.Keeper.GetEvidence(s.Ctx, c, s.Addr1)
	s.requireContainsData(ev, ev0)
	s.requireContainsData(ev, ev1)
	s.requireContainsData(ev, ev2)

	ev = s.Keeper.GetEvidence(s.Ctx, c, s.Addr2)
	s.requireContainsData(ev, ev2)
}

func (s *Suite) TestCreateBadClaim() {
	msg := claim.MsgSignClaim{Signers: []sdk.AccAddress{s.Addr1}}
	err := msg.ValidateBasic()
	s.Require().NotNil(err)

	msg = claim.MsgSignClaim{Content: types.GetDataAddressOnChainGraph([]byte{}), Signers: []sdk.AccAddress{s.Addr1}}
	res := s.Handler(s.Ctx, msg)
	s.Require().Equal(sdk.CodeUnknownRequest, res.Code)

	msg = claim.MsgSignClaim{Content: types.DataAddress([]byte{10, 2, 3, 4}), Signers: []sdk.AccAddress{s.Addr1}}
	res = s.Handler(s.Ctx, msg)
	s.Require().Equal(sdk.CodeUnknownRequest, res.Code)
}

func (s *Suite) requireContainsAddr(xs []sdk.AccAddress, x sdk.AccAddress) {
	for _, y := range xs {
		if bytes.Equal(x, y) {
			return
		}
	}
	s.Require().FailNow("can't find address")
}

func (s *Suite) requireContainsData(xs []types.DataAddress, x types.DataAddress) {
	for _, y := range xs {
		if bytes.Equal(x, y) {
			return
		}
	}
	s.Require().FailNow("can't find the data")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
