/*Package gen provides gopter.Gen random data generator helpers for the graph package.
 */
package gen

import (
	"fmt"
	"github.com/campoy/unique"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/regen-network/regen-ledger/graph"
	"github.com/regen-network/regen-ledger/graph/binary/consts"
	"github.com/regen-network/regen-ledger/graph/impl"
	"github.com/regen-network/regen-ledger/types"
	"reflect"
	"sort"
)

func Graph(resolver graph.SchemaResolver) gopter.Gen {
	return slice(0, 20, Node(resolver), reflect.TypeOf([]*impl.NodeImpl{})).
		Map(func(xs []*impl.NodeImpl) graph.Graph {
			g := impl.NewGraph()
			n := len(xs)
			if n > 0 && n%2 == 1 {
				root := xs[0]
				root.SetID(nil)
				g.WithNode(root)
				xs = xs[1:]
			}
			for _, node := range xs {
				g.WithNode(node)
			}
			return g
		})
}

func Node(resolver graph.SchemaResolver) gopter.Gen {
	return Props(resolver).FlatMap(func(x interface{}) gopter.Gen {
		pvs := x.([]propVal)
		node := impl.NewNode(nil)
		for _, pv := range pvs {
			node.SetProperty(pv.prop, pv.val)
		}
		return ID().Map(func(id types.HasURI) *impl.NodeImpl {
			node.SetID(id)
			return node.(*impl.NodeImpl)
		})
	}, reflect.TypeOf(impl.NewNode(nil)))
}

func ID() gopter.Gen {
	return gen.UInt8Range(0, consts.PrefixIDMax).FlatMap(
		func(i interface{}) gopter.Gen {
			switch i.(byte) {
			case consts.PrefixGeoAddress:
				return GeoAddress()
			case consts.PrefixAccAddress:
				return AccAddressID()
			case consts.PrefixHashID:
				return HashID()
			default:
				panic("unknown ID prefix")
			}

		},
		reflect.TypeOf((*types.HasURI)(nil)),
	)
}

func GeoAddress() gopter.Gen {
	return gen.SliceOfN(20, gen.UInt8()).Map(
		func(xs []byte) types.GeoAddress {
			return types.GeoAddress(xs)
		})
}

func AccAddressID() gopter.Gen {
	return gen.SliceOfN(20, gen.UInt8()).Map(
		func(xs []byte) graph.AccAddressID {
			return graph.AccAddressID{sdk.AccAddress(xs)}
		})
}

func HashID() gopter.Gen {
	return gen.Identifier().Map(func(x string) graph.HashID {
		return graph.HashID{x}
	})
}

type propVal struct {
	prop graph.Property
	val  interface{}
}

func Props(resolver graph.SchemaResolver) gopter.Gen {
	return slice(0, 10,
		gen.Int64Range(1, int64(resolver.MaxPropertyID())).
			FlatMap(func(x interface{}) gopter.Gen {
				p := resolver.GetPropertyByID(graph.PropertyID(x.(int64)))
				if p == nil {
					panic("can't resolve property")
				}
				return Value(p.Arity(), p.Type()).Map(
					func(v *gopter.GenResult) *gopter.GenResult {
						return gopter.NewGenResult(
							propVal{p, v.Result},
							v.Shrinker)
					})
			}, reflect.TypeOf(propVal{})),
		reflect.TypeOf([]propVal{}),
	)
}

func uniqueStrings(xs []string) []string {
	xsCopy := make([]string, len(xs))
	copy(xsCopy, xs)
	unique.Slice(&xsCopy, func(i, j int) bool {
		return xs[i] < xs[j]
	})
	return xsCopy
}

func uniqueSortedStrings(xs []string) []string {
	xs = uniqueStrings(xs)
	sort.Strings(xs)
	return xs
}

func uniqueFloat64s(xs []float64) []float64 {
	xsCopy := make([]float64, len(xs))
	copy(xsCopy, xs)
	unique.Slice(&xsCopy, func(i, j int) bool {
		return xs[i] < xs[j]
	})
	return xsCopy
}

func uniqueSortedFloat64s(xs []float64) []float64 {
	xs = uniqueFloat64s(xs)
	sort.Float64s(xs)
	return xs
}

func slice(min int, max int, g gopter.Gen, ty reflect.Type) gopter.Gen {
	return gen.IntRange(min, max).FlatMap(
		func(n interface{}) gopter.Gen {
			return gen.SliceOfN(n.(int), g)
		}, ty)
}

func Value(arity graph.Arity, propertyType graph.PropertyType) gopter.Gen {
	switch arity {
	case graph.One:
		return oneValue(propertyType)
	case graph.UnorderedSet:
		return unorderedSet(propertyType)
	case graph.OrderedSet:
		return orderedSet(propertyType)
	default:
		panic("unknown arity")
	}
}

func unorderedSet(propertyType graph.PropertyType) gopter.Gen {
	switch propertyType {
	case graph.TyString:
		return slice(1, 50, gen.AnyString(), reflect.TypeOf([]string{})).
			Map(func(xs []string) []string {
				return uniqueSortedStrings(xs)
			})
	case graph.TyDouble:
		return slice(1, 50, gen.Float64(), reflect.TypeOf([]float64{})).
			Map(func(xs []float64) []float64 {
				return uniqueSortedFloat64s(xs)
			})
	default:
		panic(fmt.Sprintf("don't know how to handle PropertyType %s", propertyType.String()))
	}
}
func orderedSet(propertyType graph.PropertyType) gopter.Gen {
	switch propertyType {
	case graph.TyString:
		return slice(1, 50, gen.AnyString(), reflect.TypeOf([]string{}))
	case graph.TyDouble:
		return slice(1, 50, gen.Float64(), reflect.TypeOf([]float64{}))
	case graph.TyBool:
		return slice(1, 50, gen.Bool(), reflect.TypeOf([]bool{}))
	default:
		panic(fmt.Sprintf("don't know how to handle PropertyType %s", propertyType.String()))
	}
}

func oneValue(propertyType graph.PropertyType) gopter.Gen {
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
