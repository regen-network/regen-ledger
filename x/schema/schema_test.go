package schema

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	"testing"
)

type TestSuite struct {
	suite.Suite
	keeper  Keeper
	handler sdk.Handler
	ctx     sdk.Context
	cms     store.CommitMultiStore
	anAddr  sdk.AccAddress
}

func (s *TestSuite) SetupTest() {
	db := dbm.NewMemDB()
	s.cms = store.NewCommitMultiStore(db)
	key := sdk.NewKVStoreKey("schema")
	cdc := codec.New()
	RegisterCodec(cdc)
	s.keeper = NewKeeper(key, cdc)
	s.cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	_ = s.cms.LoadLatestVersion()
	s.ctx = sdk.NewContext(s.cms, abci.Header{}, false, log.NewNopLogger())
	s.anAddr = sdk.AccAddress{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s.handler = NewHandler(s.keeper)
}

func (s *TestSuite) TestCreatorCantBeEmpty() {
	s.T().Log("define property")
	prop1 := PropertyDefinition{
		Name:         "test1",
		PropertyType: TyBool,
	}
	_, _, err := s.keeper.DefineProperty(s.ctx, prop1)
	s.Require().NotNil(err)
}

func (s *TestSuite) TestNameCantBeEmpty() {
	s.T().Log("define property")
	prop1 := PropertyDefinition{
		Creator:      s.anAddr,
		PropertyType: TyBool,
	}
	_, _, err := s.keeper.DefineProperty(s.ctx, prop1)
	s.Require().NotNil(err)
}

func (s *TestSuite) TestPropertyCanOnlyBeDefinedOnce() {
	s.T().Log("define property")
	prop1 := PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "test1",
		PropertyType: TyBool,
	}
	_, _, err := s.keeper.DefineProperty(s.ctx, prop1)
	s.Require().Nil(err)

	s.T().Log("try to define property with same name")
	prop2 := PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "test1",
		PropertyType: TyInteger,
	}
	_, _, err = s.keeper.DefineProperty(s.ctx, prop2)
	s.Require().NotNil(err)
}

func (s *TestSuite) TestCheckPropertyType() {
	s.T().Log("valid property type should be accepted")
	prop1 := PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "test1",
		PropertyType: TyBool,
	}
	err := prop1.ValidateBasic()
	s.Require().Nil(err)
	_, _, err = s.keeper.DefineProperty(s.ctx, prop1)
	s.Require().Nil(err)

	s.T().Log("invalid property type should be rejected")
	prop2 := PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "test2",
		PropertyType: PropertyType(12345678),
	}
	err = prop2.ValidateBasic()
	s.Require().NotNil(err)
	_, _, err = s.keeper.DefineProperty(s.ctx, prop2)
	s.Require().NotNil(err)
}

func (s *TestSuite) TestCanRetrieveProperty() {
	s.T().Log("define property")
	prop := PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "test1",
		PropertyType: TyBool,
	}
	id, url, err := s.keeper.DefineProperty(s.ctx, prop)
	s.Require().Nil(err)

	s.T().Log("try retrieve property")
	propCopy, found := s.keeper.GetProperty(s.ctx, id)
	s.Require().True(found)
	s.Require().True(bytes.Equal(prop.Creator, propCopy.Creator))
	s.Require().Equal(prop.Name, propCopy.Name)
	s.Require().Equal(prop.PropertyType, propCopy.PropertyType)
	s.Require().Equal(prop.Many, propCopy.Many)

	s.T().Log("try retrieve property id from URL")
	idCopy := s.keeper.GetPropertyID(s.ctx, url)
	s.Require().Equal(id, idCopy)
}

func (s *TestSuite) TestIncrementPropertyID() {
	s.T().Log("create one property")
	prop1 := PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "test1",
		PropertyType: TyBool,
	}
	id, url, err := s.keeper.DefineProperty(s.ctx, prop1)
	s.Require().Nil(err)
	s.Require().Equal(PropertyID(1), id)

	s.T().Log("create another property")
	prop2 := PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "test2",
		PropertyType: TyString,
		Many:         true,
	}
	id2, url2, err := s.keeper.DefineProperty(s.ctx, prop2)
	s.Require().Nil(err)
	s.Require().Equal(PropertyID(2), id2)
	s.Require().NotEqual(url, url2)
}

func (s *TestSuite) TestPropertyNotFound() {
	_, found := s.keeper.GetProperty(s.ctx, 0)
	s.Require().False(found)

	id := s.keeper.GetPropertyID(s.ctx, "")
	s.Require().Equal(PropertyID(0), id)
}

func (s *TestSuite) TestPropertyNameRegex() {
	prop1 := PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "TestCamelCase",
		PropertyType: TyString,
		Many:         true,
	}
	err := prop1.ValidateBasic()
	s.Require().NotNil(err)
}

func (s *TestSuite) TestDefinePropertyHandler() {
	s.T().Log("create one property")
	prop1 := PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "test1",
		PropertyType: TyBool,
	}
	res := s.handler(s.ctx, prop1)
	s.Require().Equal(sdk.CodeOK, res.Code)
	s.Require().Equal(prop1.URL(), string(res.Tags[0].Value))
	s.Require().Equal("1", string(res.Tags[1].Value))
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
