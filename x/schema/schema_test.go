package schema_test

import (
	"bytes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/graph"
	"github.com/regen-network/regen-ledger/x/schema"
	"github.com/regen-network/regen-ledger/x/schema/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	test.Harness
}

func (s *TestSuite) SetupTest() {
	s.Harness.Setup()
}

func (s *TestSuite) TestCreatorCantBeEmpty() {
	s.T().Log("define property")
	prop1 := schema.PropertyDefinition{
		Name:         "test1",
		PropertyType: graph.TyBool,
	}
	_, _, err := s.Keeper.DefineProperty(s.Ctx, prop1)
	s.Require().NotNil(err)
}

func (s *TestSuite) TestNameCantBeEmpty() {
	s.T().Log("define property")
	prop1 := schema.PropertyDefinition{
		Creator:      s.AnAddr,
		PropertyType: graph.TyBool,
	}
	_, _, err := s.Keeper.DefineProperty(s.Ctx, prop1)
	s.Require().NotNil(err)
}

func (s *TestSuite) TestPropertyCanOnlyBeDefinedOnce() {
	s.T().Log("define property")
	prop1 := schema.PropertyDefinition{
		Creator:      s.AnAddr,
		Name:         "test1",
		PropertyType: graph.TyBool,
	}
	_, _, err := s.Keeper.DefineProperty(s.Ctx, prop1)
	s.Require().Nil(err)

	s.T().Log("try to define property with same name")
	prop2 := schema.PropertyDefinition{
		Creator:      s.AnAddr,
		Name:         "test1",
		PropertyType: graph.TyInteger,
	}
	_, _, err = s.Keeper.DefineProperty(s.Ctx, prop2)
	s.Require().NotNil(err)
}

func (s *TestSuite) TestCheckPropertyType() {
	s.T().Log("invalid property type should be rejected")
	prop1 := schema.PropertyDefinition{
		Creator:      s.AnAddr,
		Name:         "test1",
		PropertyType: graph.PropertyType(12345678),
	}
	err := prop1.ValidateBasic()
	s.Require().NotNil(err)
	_, _, err = s.Keeper.DefineProperty(s.Ctx, prop1)
	s.Require().NotNil(err)
}

func (s *TestSuite) TestCheckArity() {
	s.T().Log("invalid arity should be rejected")
	prop1 := schema.PropertyDefinition{
		Creator:      s.AnAddr,
		Name:         "test1",
		PropertyType: graph.TyObject,
		Arity:        graph.Arity(513848),
	}
	err := prop1.ValidateBasic()
	s.Require().NotNil(err)
	_, _, err = s.Keeper.DefineProperty(s.Ctx, prop1)
	s.Require().NotNil(err)
}

func (s *TestSuite) TestCanRetrieveProperty() {
	s.T().Log("define property")
	prop := schema.PropertyDefinition{
		Creator:      s.AnAddr,
		Name:         "test1",
		PropertyType: graph.TyBool,
	}
	id, url, err := s.Keeper.DefineProperty(s.Ctx, prop)
	s.Require().Nil(err)

	s.T().Log("try retrieve property")
	propCopy, found := s.Keeper.GetPropertyDefinition(s.Ctx, id)
	s.Require().True(found)
	s.Require().True(bytes.Equal(prop.Creator, propCopy.Creator))
	s.Require().Equal(prop.Name, propCopy.Name)
	s.Require().Equal(prop.PropertyType, propCopy.PropertyType)
	s.Require().Equal(prop.Arity, propCopy.Arity)

	s.T().Log("try retrieve property id from URL")
	idCopy := s.Keeper.GetPropertyID(s.Ctx, url.String())
	s.Require().Equal(id, idCopy)
}

func (s *TestSuite) TestIncrementPropertyID() {
	s.T().Log("create one property")
	prop1 := schema.PropertyDefinition{
		Creator:      s.AnAddr,
		Name:         "test1",
		PropertyType: graph.TyBool,
	}
	id, url, err := s.Keeper.DefineProperty(s.Ctx, prop1)
	s.Require().Nil(err)
	s.Require().Equal(graph.PropertyID(1), id)

	s.T().Log("create another property")
	prop2 := schema.PropertyDefinition{
		Creator:      s.AnAddr,
		Name:         "test2",
		PropertyType: graph.TyString,
		Arity:        graph.UnorderedSet,
	}
	id2, url2, err := s.Keeper.DefineProperty(s.Ctx, prop2)
	s.Require().Nil(err)
	s.Require().Equal(graph.PropertyID(2), id2)
	s.Require().NotEqual(url, url2)
}

func (s *TestSuite) TestPropertyNotFound() {
	_, found := s.Keeper.GetPropertyDefinition(s.Ctx, 0)
	s.Require().False(found)

	id := s.Keeper.GetPropertyID(s.Ctx, "")
	s.Require().Equal(graph.PropertyID(0), id)
}

func (s *TestSuite) TestPropertyNameRegex() {
	prop1 := schema.PropertyDefinition{
		Creator:      s.AnAddr,
		Name:         "TestCamelCase",
		PropertyType: graph.TyString,
		Arity:        graph.OrderedSet,
	}
	err := prop1.ValidateBasic()
	s.Require().NotNil(err)
}

func (s *TestSuite) TestDefinePropertyHandler() {
	s.T().Log("create one property")
	prop1 := schema.PropertyDefinition{
		Creator:      s.AnAddr,
		Name:         "test1",
		PropertyType: graph.TyBool,
	}
	res := s.Handler(s.Ctx, prop1)
	s.Require().Equal(sdk.CodeOK, res.Code)
	s.Require().Equal(prop1.URI().String(), string(res.Tags[0].Value))
	s.Require().Equal("1", string(res.Tags[1].Value))
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
