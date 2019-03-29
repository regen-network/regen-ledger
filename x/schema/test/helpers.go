package test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/graph"
	"github.com/regen-network/regen-ledger/x/schema"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
)

type Harness struct {
	suite.Suite
	Keeper   schema.Keeper
	Handler  sdk.Handler
	Ctx      sdk.Context
	Cms      store.CommitMultiStore
	AnAddr   sdk.AccAddress
	Resolver graph.SchemaResolver
}

func (s *Harness) Setup() {
	db := dbm.NewMemDB()
	s.Cms = store.NewCommitMultiStore(db)
	key := sdk.NewKVStoreKey("schema")
	cdc := codec.New()
	schema.RegisterCodec(cdc)
	s.Keeper = schema.NewKeeper(key, cdc)
	s.Cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	_ = s.Cms.LoadLatestVersion()
	s.Ctx = sdk.NewContext(s.Cms, abci.Header{}, false, log.NewNopLogger())
	s.AnAddr = sdk.AccAddress{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s.Handler = schema.NewHandler(s.Keeper)
	s.Resolver = schema.NewOnChainSchemaResolver(s.Keeper, s.Ctx)
}

func (s *Harness) CreateSampleSchema() {
	CreateSampleSchema(s.Suite, s.Keeper, s.Ctx, s.AnAddr)
}

func CreateSampleSchema(s suite.Suite, keeper schema.Keeper, ctx sdk.Context, anAddr sdk.AccAddress) {
	anAddr = sdk.AccAddress{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// create a mock schema
	_, _, err := keeper.DefineProperty(ctx, schema.PropertyDefinition{
		Creator:      anAddr,
		Name:         "name",
		PropertyType: graph.TyString,
	})
	s.Require().Nil(err)
	_, _, err = keeper.DefineProperty(ctx, schema.PropertyDefinition{
		Creator:      anAddr,
		Name:         "x",
		PropertyType: graph.TyDouble,
	})
	s.Require().Nil(err)
	_, _, err = keeper.DefineProperty(ctx, schema.PropertyDefinition{
		Creator:      anAddr,
		Name:         "b",
		PropertyType: graph.TyBool,
	})
	s.Require().Nil(err)
	_, _, err = keeper.DefineProperty(ctx, schema.PropertyDefinition{
		Creator:      anAddr,
		Name:         "names",
		PropertyType: graph.TyString,
		Arity:        graph.UnorderedSet,
	})
	s.Require().Nil(err)
	_, _, err = keeper.DefineProperty(ctx, schema.PropertyDefinition{
		Creator:      anAddr,
		Name:         "xs",
		PropertyType: graph.TyDouble,
		Arity:        graph.UnorderedSet,
	})
	s.Require().Nil(err)
	s.Require().Nil(err)
	_, _, err = keeper.DefineProperty(ctx, schema.PropertyDefinition{
		Creator:      anAddr,
		Name:         "nameList",
		PropertyType: graph.TyString,
		Arity:        graph.OrderedSet,
	})
	s.Require().Nil(err)
	_, _, err = keeper.DefineProperty(ctx, schema.PropertyDefinition{
		Creator:      anAddr,
		Name:         "xList",
		PropertyType: graph.TyDouble,
		Arity:        graph.OrderedSet,
	})
	s.Require().Nil(err)
	_, _, err = keeper.DefineProperty(ctx, schema.PropertyDefinition{
		Creator:      anAddr,
		Name:         "bList",
		PropertyType: graph.TyBool,
		Arity:        graph.OrderedSet,
	})
	s.Require().Nil(err)
}
