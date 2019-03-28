package xrb

import (
	"bytes"
	"fmt"
	"github.com/campoy/unique"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/schema"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	"reflect"
	"sort"
	"testing"
)

// TODO graph generator
// TODO verify graph can be serialized and deserialized and is equivalent and has same hash

type TestSuite struct {
	suite.Suite
	keeper   schema.Keeper
	ctx      sdk.Context
	cms      store.CommitMultiStore
	anAddr   sdk.AccAddress
	props    []schema.PropertyID
	resolver SchemaResolver
}

func (s *TestSuite) SetupSuite() {
	db := dbm.NewMemDB()
	s.cms = store.NewCommitMultiStore(db)
	key := sdk.NewKVStoreKey("schema")
	cdc := codec.New()
	schema.RegisterCodec(cdc)
	s.keeper = schema.NewKeeper(key, cdc)
	s.cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	_ = s.cms.LoadLatestVersion()
	s.ctx = sdk.NewContext(s.cms, abci.Header{}, false, log.NewNopLogger())
	s.anAddr = sdk.AccAddress{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s.resolver = onChainSchemaResolver{s.keeper, s.ctx}

	// create a mock schema
	p1, _, err := s.keeper.DefineProperty(s.ctx, schema.PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "name",
		PropertyType: schema.TyString,
	})
	s.Require().Nil(err)
	p2, _, err := s.keeper.DefineProperty(s.ctx, schema.PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "x",
		PropertyType: schema.TyDouble,
	})
	s.Require().Nil(err)
	p3, _, err := s.keeper.DefineProperty(s.ctx, schema.PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "b",
		PropertyType: schema.TyBool,
	})
	s.Require().Nil(err)
	p4, _, err := s.keeper.DefineProperty(s.ctx, schema.PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "names",
		PropertyType: schema.TyString,
		Arity:        schema.UnorderedSet,
	})
	s.Require().Nil(err)
	p5, _, err := s.keeper.DefineProperty(s.ctx, schema.PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "xs",
		PropertyType: schema.TyDouble,
		Arity:        schema.UnorderedSet,
	})
	s.Require().Nil(err)
	s.Require().Nil(err)
	p6, _, err := s.keeper.DefineProperty(s.ctx, schema.PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "nameList",
		PropertyType: schema.TyString,
		Arity:        schema.OrderedSet,
	})
	s.Require().Nil(err)
	p7, _, err := s.keeper.DefineProperty(s.ctx, schema.PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "xList",
		PropertyType: schema.TyDouble,
		Arity:        schema.OrderedSet,
	})
	s.Require().Nil(err)
	p8, _, err := s.keeper.DefineProperty(s.ctx, schema.PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "bList",
		PropertyType: schema.TyBool,
		Arity:        schema.OrderedSet,
	})
	s.Require().Nil(err)
	s.props = []schema.PropertyID{p1, p2, p3, p4, p5, p6, p7, p8}
}

func (s *TestSuite) TestGenGraph() {
	gs, ok := gen.SliceOfN(3, s.GenGraph()).Sample()
	if ok {
		for _, g := range gs.([]*graph) {
			s.T().Log(g.String())
		}
	}
}

func (s *TestSuite) TestProperties() {
	params := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(params)
	properties.Property("can round trip serialize/deserialize graphs and calculate the same hashes",
		prop.ForAll(func(g Graph) (bool, error) {
			w := new(bytes.Buffer)
			hash, err := SerializeGraph(s.resolver, g, w)
			if err != nil {
				return false, err
			}
			_, hashCopy, err := DeserializeGraph(s.resolver, w)
			if err != nil {
				return false, err
			}
			if !bytes.Equal(hash, hashCopy) {
				return false, fmt.Errorf("hashes do not match")
			}

			// TODO actually compare the contents of the graphs, not just their hashes
			return true, nil
		}, s.GenGraph()))

	properties.TestingRun(s.T())
}

func (s *TestSuite) GenGraph() gopter.Gen {
	return GenSlice(0, 20, s.GenNode(), reflect.TypeOf([]*node{})).
		Map(func(xs []*node) *graph {
			g := &graph{}
			n := len(xs)
			if n > 0 && n%2 == 0 {
				root := xs[0]
				root.id = nil
				g.rootNode = root
				xs = xs[1:n]
			}
			g.nodeNames = make([]types.HasURI, len(xs))
			g.nodes = make(map[string]*node)
			for i, node := range xs {
				id := node.id
				g.nodeNames[i] = id
				g.nodes[id.String()] = node
			}
			return g
		})
}

func (s *TestSuite) GenNode() gopter.Gen {
	return s.GenProps().FlatMap(func(x interface{}) gopter.Gen {
		pvs := x.([]propVal)
		propNames := make([]Property, len(pvs))
		props := make(map[schema.PropertyID]interface{})
		for i, pv := range pvs {
			propNames[i] = pv.prop
			props[pv.prop.ID()] = pv.val
		}
		return s.GenID().Map(func(id types.HasURI) *node {
			return &node{
				id,
				propNames,
				props,
			}
		})
	}, reflect.TypeOf(&node{}))
}

func (s *TestSuite) GenID() gopter.Gen {
	return gen.UInt8Range(prefixIDMin, prefixIDMax).FlatMap(
		func(i interface{}) gopter.Gen {
			switch i.(byte) {
			case prefixGeoAddress:
				return GenGeoAddress()
			case prefixAccAddress:
				return GenAccAddressID()
			case prefixDataAddress:
				return GenDataAddress()
			case prefixHashID:
				return GenHashID()
			default:
				panic("unknown ID prefix")
			}

		},
		reflect.TypeOf((*types.HasURI)(nil)),
	)
}

func GenGeoAddress() gopter.Gen {
	return gen.SliceOfN(20, gen.UInt8()).Map(
		func(xs []byte) types.GeoAddress {
			return types.GeoAddress(xs)
		})
}

func GenDataAddress() gopter.Gen {
	return gen.SliceOfN(20, gen.UInt8()).Map(
		func(xs []byte) types.DataAddress {
			return types.DataAddress(xs)
		})
}

func GenAccAddressID() gopter.Gen {
	return gen.SliceOfN(20, gen.UInt8()).Map(
		func(xs []byte) AccAddressID {
			return AccAddressID{sdk.AccAddress(xs)}
		})
}

func GenHashID() gopter.Gen {
	return gen.Identifier().Map(func(x string) HashID {
		return HashID{x}
	})
}

type propVal struct {
	prop Property
	val  interface{}
}

type propVals []propVal

func (pv propVals) Len() int { return len(pv) }

func (pv propVals) Swap(i, j int) {
	pv[i], pv[j] = pv[j], pv[i]
}

func (pv propVals) Less(i, j int) bool {
	return pv[i].prop.URI().String() < pv[j].prop.URI().String()
}

func (s *TestSuite) GenProps() gopter.Gen {
	nProps := len(s.props)
	return GenSlice(0, nProps,
		gen.UInt64Range(1, uint64(nProps)).
			FlatMap(func(x interface{}) gopter.Gen {
				prop := s.resolver.GetPropertyByID(schema.PropertyID(x.(uint64)))
				if prop == nil {
					panic("can't resolve property")
				}
				return GenValue(prop.Arity(), prop.Type()).Map(
					func(v *gopter.GenResult) *gopter.GenResult {
						return gopter.NewGenResult(
							propVal{prop, v.Result},
							v.Shrinker)
					})
			}, reflect.TypeOf(propVal{})),
		reflect.TypeOf([]propVal{}),
	).Map(func(xs []propVal) []propVal {
		props := make([]propVal, len(xs))
		copy(props, xs)
		unique.Slice(&props, func(i, j int) bool {
			return xs[i].prop.URI().String() < xs[j].prop.URI().String()
		})
		sort.Sort(propVals(props))
		return props
	})
}

func UniqueStrings(xs []string) []string {
	xsCopy := make([]string, len(xs))
	copy(xsCopy, xs)
	unique.Slice(&xsCopy, func(i, j int) bool {
		return xs[i] < xs[j]
	})
	return xsCopy
}

func UniqueSortedStrings(xs []string) []string {
	xs = UniqueStrings(xs)
	sort.Strings(xs)
	return xs
}

func UniqueFloat64s(xs []float64) []float64 {
	xsCopy := make([]float64, len(xs))
	copy(xsCopy, xs)
	unique.Slice(&xsCopy, func(i, j int) bool {
		return xs[i] < xs[j]
	})
	return xsCopy
}

func UniqueSortedFloat64s(xs []float64) []float64 {
	xs = UniqueFloat64s(xs)
	sort.Float64s(xs)
	return xs
}

func GenSlice(min int, max int, g gopter.Gen, ty reflect.Type) gopter.Gen {
	return gen.IntRange(min, max).FlatMap(
		func(n interface{}) gopter.Gen {
			return gen.SliceOfN(n.(int), g)
		},
		reflect.TypeOf([]schema.PropertyID{}),
	)
}

func GenValue(arity schema.Arity, propertyType schema.PropertyType) gopter.Gen {
	switch arity {
	case schema.One:
		return GenOneValue(propertyType)
	case schema.UnorderedSet:
		return GenUnorderedSet(propertyType)
	case schema.OrderedSet:
		return GenOrderedSet(propertyType)
	default:
		panic("unknown arity")
	}
}

func GenUnorderedSet(propertyType schema.PropertyType) gopter.Gen {
	switch propertyType {
	case schema.TyString:
		return GenSlice(0, 50, gen.AnyString(), reflect.TypeOf([]string{})).
			Map(func(xs []string) []string {
				return UniqueSortedStrings(xs)
			})
	case schema.TyDouble:
		return GenSlice(0, 50, gen.Float64(), reflect.TypeOf([]float64{})).
			Map(func(xs []float64) []float64 {
				return UniqueSortedFloat64s(xs)
			})
	default:
		panic(fmt.Sprintf("don't know how to handle PropertyType %s", propertyType.String()))
	}
}
func GenOrderedSet(propertyType schema.PropertyType) gopter.Gen {
	switch propertyType {
	case schema.TyString:
		return GenSlice(0, 50, gen.AnyString(), reflect.TypeOf([]string{}))
	case schema.TyDouble:
		return GenSlice(0, 50, gen.Float64(), reflect.TypeOf([]float64{}))
	case schema.TyBool:
		return GenSlice(0, 50, gen.Bool(), reflect.TypeOf([]bool{}))
	default:
		panic(fmt.Sprintf("don't know how to handle PropertyType %s", propertyType.String()))
	}
}

func GenOneValue(propertyType schema.PropertyType) gopter.Gen {
	switch propertyType {
	case schema.TyString:
		return gen.AnyString()
	case schema.TyDouble:
		return gen.Float64()
	case schema.TyBool:
		return gen.Bool()
	default:
		panic(fmt.Sprintf("don't know how to handle PropertyType %s", propertyType.String()))
	}
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
