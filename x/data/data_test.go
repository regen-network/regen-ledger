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
