package data_test

import (
	"crypto/sha256"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/data"
	schematest "github.com/regen-network/regen-ledger/x/schema/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	schematest.Harness
	Keeper  data.Keeper
	Handler sdk.Handler
}

func (s *Suite) SetupTest() {
	s.Setup()
	dataKey := sdk.NewKVStoreKey("data")
	data.RegisterCodec(s.Cdc)
	s.Keeper = data.NewKeeper(dataKey, s.Cdc)
	s.Handler = data.NewHandler(s.Keeper)
	s.Cms.MountStoreWithDB(dataKey, sdk.StoreTypeIAVL, s.Db)
	_ = s.Cms.LoadLatestVersion()
	s.CreateSampleSchema()
}

//func (s *Suite) TestStoreDataGraph() {
//	params := gopter.DefaultTestParameters()
//	properties := gopter.NewProperties(params)
//	properties.Property("can round trip store/retrieve graphs",
//		prop.ForAll(func(g1 graph.Graph) (bool, error) {
//			buf := new(bytes.Buffer)
//			err := binary.SerializeGraph(s.Resolver, g1, buf)
//			if err != nil {
//				return false, err
//			}
//			hash := graph.Hash(g1)
//
//			// check if we have existing data (because the generator repeats values)
//			addr := types.GetDataAddressGraph(hash)
//			bz, err := s.Keeper.GetData(s.Ctx, addr)
//			if bz == nil {
//				res := s.Handler(s.Ctx, data.MsgStoreGraph{Hash: hash, Data: buf.Bytes(), Signer: s.Addr1})
//				if res.Code != sdk.CodeOK {
//					return false, fmt.Errorf("%+v", res)
//				}
//
//				url := res.Tags[0].Value
//				addr2 := types.MustDecodeBech32DataAddress(string(url))
//				if !bytes.Equal(addr, addr2) {
//					return false, fmt.Errorf("unexpected DataAddress %+v, %+v", []byte(addr), []byte(addr2))
//				}
//
//				// verify can't store same graph again
//				res = s.Handler(s.Ctx, data.MsgStoreGraph{Hash: hash, Data: buf.Bytes(), Signer: s.Addr1})
//				if res.Code == sdk.CodeOK {
//					return false, fmt.Errorf("shouldn't be able to store the same graph twice")
//				}
//
//				bz, err = s.Keeper.GetData(s.Ctx, addr)
//				if err != nil {
//					return false, err
//				}
//			}
//
//			g2, err := binary.DeserializeGraph(s.Resolver, bytes.NewBuffer(bz))
//			if err != nil {
//				return false, err
//			}
//
//			hash2 := graph.Hash(g2)
//			if !bytes.Equal(hash, hash2) {
//				return false, fmt.Errorf("wrong hash")
//			}
//
//			s1, err := graph.CanonicalString(g1)
//			if err != nil {
//				return false, err
//			}
//
//			s2, err := graph.CanonicalString(g2)
//			if err != nil {
//				return false, err
//			}
//
//			if s1 != s2 {
//				return false, fmt.Errorf("wrong canonical text")
//			}
//
//			return true, nil
//		}, gen.Graph(s.Resolver)))
//	properties.TestingRun(s.T())
//}

func (s *Suite) TestTrackRawData() {
	someData := "sdlkghsdg2368uysdgiuskdhfg23t69sdgkj2"
	hasher := sha256.New()
	hasher.Write([]byte(someData))
	hash := hasher.Sum(nil)
	res := s.Handler(s.Ctx, data.MsgTrackRawData{Sha256Hash: hash, Url: "", Signer: s.Addr1})
	s.Require().Equal(sdk.CodeOK, res.Code)
	s.Require().Equal(types.GetDataAddressRawData(hash).String(), string(res.Tags[0].Value))
	urls, err := s.Keeper.GetRawDataURLs(s.Ctx, hash)
	s.Require().Nil(err)
	s.Require().Empty(urls)
	someUrl := "http://example.com/nowhere"
	res = s.Handler(s.Ctx, data.MsgTrackRawData{Sha256Hash: hash, Url: someUrl, Signer: s.Addr1})
	s.Require().Equal(sdk.CodeOK, res.Code)
	s.Require().Equal(types.GetDataAddressRawData(hash).String(), string(res.Tags[0].Value))
	urls, err = s.Keeper.GetRawDataURLs(s.Ctx, hash)
	s.Require().Nil(err)
	s.Require().Len(urls, 1)
	s.Require().Equal(someUrl, urls[0])
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
