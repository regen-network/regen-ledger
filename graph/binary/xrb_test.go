package binary

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
	"github.com/regen-network/regen-ledger/graph"
	"github.com/regen-network/regen-ledger/graph/impl"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/util"
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
		PropertyType: graph.TyString,
	})
	s.Require().Nil(err)
	p2, _, err := s.keeper.DefineProperty(s.ctx, schema.PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "x",
		PropertyType: graph.TyDouble,
	})
	s.Require().Nil(err)
	p3, _, err := s.keeper.DefineProperty(s.ctx, schema.PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "b",
		PropertyType: graph.TyBool,
	})
	s.Require().Nil(err)
	p4, _, err := s.keeper.DefineProperty(s.ctx, schema.PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "names",
		PropertyType: graph.TyString,
		Arity:        graph.UnorderedSet,
	})
	s.Require().Nil(err)
	p5, _, err := s.keeper.DefineProperty(s.ctx, schema.PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "xs",
		PropertyType: graph.TyDouble,
		Arity:        graph.UnorderedSet,
	})
	s.Require().Nil(err)
	s.Require().Nil(err)
	p6, _, err := s.keeper.DefineProperty(s.ctx, schema.PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "nameList",
		PropertyType: graph.TyString,
		Arity:        graph.OrderedSet,
	})
	s.Require().Nil(err)
	p7, _, err := s.keeper.DefineProperty(s.ctx, schema.PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "xList",
		PropertyType: graph.TyDouble,
		Arity:        graph.OrderedSet,
	})
	s.Require().Nil(err)
	p8, _, err := s.keeper.DefineProperty(s.ctx, schema.PropertyDefinition{
		Creator:      s.anAddr,
		Name:         "bList",
		PropertyType: graph.TyBool,
		Arity:        graph.OrderedSet,
	})
	s.Require().Nil(err)
	s.props = []schema.PropertyID{p1, p2, p3, p4, p5, p6, p7, p8}
}

func (s *TestSuite) TestGenGraph() {
	gs, ok := gen.SliceOfN(3, s.GenGraph()).Sample()
	if ok {
		for _, g := range gs.([]graph.Graph) {
			s.T().Logf("Graph %s:\n %s",
				util.MustEncodeBech32(types.Bech32DataAddressPrefix, graph.Hash(g)),
				g.String(),
			)
		}
	}
}

func (s *TestSuite) TestProperties() {
	params := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(params)
	properties.Property("can round trip serialize/deserialize graphs with same hashes",
		prop.ForAll(func(g1 graph.Graph) (bool, error) {
			txt1, err := graph.CanonicalString(g1)
			if err != nil {
				return false, err
			}
			hash1 := graph.Hash(g1)
			w := new(bytes.Buffer)
			err = SerializeGraph(s.resolver, g1, w)
			if err != nil {
				return false, err
			}
			g2, err := DeserializeGraph(s.resolver, w)
			txt2, err := graph.CanonicalString(g1)
			if err != nil {
				return false, err
			}
			hash2 := graph.Hash(g2)
			s.T().Logf("%s %s", txt1, txt2)
			if txt1 != txt2 {
				return false, fmt.Errorf("canonical strings do not match")
			}
			if !bytes.Equal(hash1, hash2) {
				return false, fmt.Errorf("hashes do not match")
			}

			// TODO actually compare the contents of the graphs, not just their hashes
			return true, nil
		}, s.GenGraph()))

	properties.TestingRun(s.T())
}

func (s *TestSuite) GenGraph() gopter.Gen {
	return GenSlice(0, 20, s.GenNode(), reflect.TypeOf([]graph.Node{})).
		Map(func(xs []graph.Node) graph.Graph {
			g := impl.NewGraph()
			n := len(xs)
			if n > 0 && n%2 == 0 {
				root := xs[0]
				root.SetID(nil)
				g.WithNode(root)
				xs = xs[1:n]
			}
			for _, node := range xs {
				g.WithNode(node)
			}
			return g
		})
}

func (s *TestSuite) GenNode() gopter.Gen {
	return s.GenProps().FlatMap(func(x interface{}) gopter.Gen {
		pvs := x.([]propVal)
		node := impl.NewNode(nil)
		for _, pv := range pvs {
			node.SetProperty(pv.prop, pv.val)
		}
		return s.GenID().Map(func(id types.HasURI) graph.Node {
			node.SetID(id)
			return node
		})
	}, reflect.TypeOf((*graph.Node)(nil)))
}

func (s *TestSuite) GenID() gopter.Gen {
	return gen.UInt8Range(prefixIDMin, prefixIDMax).FlatMap(
		func(i interface{}) gopter.Gen {
			switch i.(byte) {
			case prefixGeoAddress:
				return GenGeoAddress()
			case prefixAccAddress:
				return GenAccAddressID()
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
	prop graph.Property
	val  interface{}
}

func (s *TestSuite) GenProps() gopter.Gen {
	nProps := len(s.props)
	return GenSlice(0, nProps,
		gen.UInt64Range(1, uint64(nProps)).
			FlatMap(func(x interface{}) gopter.Gen {
				p := s.resolver.GetPropertyByID(schema.PropertyID(x.(uint64)))
				if p == nil {
					panic("can't resolve property")
				}
				return GenValue(p.Arity(), p.Type()).Map(
					func(v *gopter.GenResult) *gopter.GenResult {
						return gopter.NewGenResult(
							propVal{p, v.Result},
							v.Shrinker)
					})
			}, reflect.TypeOf(propVal{})),
		reflect.TypeOf([]propVal{}),
	)
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
		}, ty)
}

func GenValue(arity graph.Arity, propertyType graph.PropertyType) gopter.Gen {
	switch arity {
	case graph.One:
		return GenOneValue(propertyType)
	case graph.UnorderedSet:
		return GenUnorderedSet(propertyType)
	case graph.OrderedSet:
		return GenOrderedSet(propertyType)
	default:
		panic("unknown arity")
	}
}

func GenUnorderedSet(propertyType graph.PropertyType) gopter.Gen {
	switch propertyType {
	case graph.TyString:
		return GenSlice(0, 50, gen.AnyString(), reflect.TypeOf([]string{})).
			Map(func(xs []string) []string {
				return UniqueSortedStrings(xs)
			})
	case graph.TyDouble:
		return GenSlice(0, 50, gen.Float64(), reflect.TypeOf([]float64{})).
			Map(func(xs []float64) []float64 {
				return UniqueSortedFloat64s(xs)
			})
	default:
		panic(fmt.Sprintf("don't know how to handle PropertyType %s", propertyType.String()))
	}
}
func GenOrderedSet(propertyType graph.PropertyType) gopter.Gen {
	switch propertyType {
	case graph.TyString:
		return GenSlice(0, 50, gen.AnyString(), reflect.TypeOf([]string{}))
	case graph.TyDouble:
		return GenSlice(0, 50, gen.Float64(), reflect.TypeOf([]float64{}))
	case graph.TyBool:
		return GenSlice(0, 50, gen.Bool(), reflect.TypeOf([]bool{}))
	default:
		panic(fmt.Sprintf("don't know how to handle PropertyType %s", propertyType.String()))
	}
}

func GenOneValue(propertyType graph.PropertyType) gopter.Gen {
	switch propertyType {
	case graph.TyString:
		return gen.AnyString()
	case graph.TyDouble:
		return gen.Float64()
	case graph.TyBool:
		return gen.Bool()
	default:
		panic(fmt.Sprintf("don't know how to handle PropertyType %s", propertyType.String()))
	}
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
