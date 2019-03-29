package binary

import (
	"encoding/binary"
	"fmt"
	"github.com/regen-network/regen-ledger/graph"
	"github.com/regen-network/regen-ledger/graph/impl"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/schema"
	"io"
	"math"
)
import sdk "github.com/cosmos/cosmos-sdk/types"

// DeserializeGraph deserializes a Graph from a binary reader
func DeserializeGraph(resolver SchemaResolver, r io.ByteScanner) (g graph.Graph, err error) {
	ctx := &dszContext{resolver, r, 0}
	g, err = ctx.readGraph()
	if err != nil {
		return nil, err
	}
	return g, nil
}

type dszContext struct {
	resolver SchemaResolver
	r        io.ByteScanner
	version  uint64
}

func (ctx *dszContext) readGraph() (g graph.Graph, err error) {
	ctx.version = ctx.mustReadUvarint64()
	haveRootNode := ctx.mustReadByte()
	g = impl.NewGraph()
	if haveRootNode == 1 {
		r, err := ctx.readNodeProperties()
		if err != nil {
			return nil, err
		}
		g.WithNode(r)
	}
	nNodes := ctx.mustReadVarint()
	for i := 0; i < nNodes; i++ {
		n, err := ctx.readNode()
		if err != nil {
			return nil, err
		}
		g.WithNode(n)
	}
	return g, nil
}

func (ctx *dszContext) readNode() (n graph.Node, err error) {
	id, err := ctx.readID()
	if err != nil {
		return nil, err
	}
	n, err = ctx.readNodeProperties()
	if err != nil {
		return nil, err
	}
	n.SetID(id)
	return n, nil
}

func (ctx *dszContext) readID() (id types.HasURI, err error) {
	prefix := ctx.mustReadByte()
	switch prefix {
	case prefixGeoAddress:
		return types.GeoAddress(ctx.readByteSlice()), nil
	case prefixAccAddress:
		return AccAddressID{sdk.AccAddress(ctx.readByteSlice())}, nil
	case prefixHashID:
		return HashID{ctx.readString()}, nil
	default:
		return nil, fmt.Errorf("unexpected ID prefix %d", prefix)
	}
}

func (ctx *dszContext) readNodeProperties() (n graph.Node, err error) {
	n = impl.NewNode(nil)
	nProps := ctx.mustReadVarint()
	for i := 0; i < nProps; i++ {
		prop, value, err := ctx.readProperty()
		// TODO verify ordering
		// hash
		if err != nil {
			return nil, err
		}
		n.SetProperty(prop, value)
	}
	return n, nil
}

func (ctx *dszContext) readProperty() (prop graph.Property, value interface{}, err error) {
	prefix := ctx.mustReadByte()
	if prefix != prefixPropertyID {
		return nil, nil, fmt.Errorf("unexpected property ID prefix %d", prefix)
	}
	id := ctx.mustReadVarint64()
	prop = ctx.resolver.GetPropertyByID(schema.PropertyID(id))
	if prop == nil {
		return nil, nil, fmt.Errorf("can't resolve property with ID %d", id)
	}
	val, err := ctx.readValue(prop)
	if err != nil {
		return nil, nil, err
	}
	return prop, val, nil
}

func (ctx *dszContext) readValue(prop graph.Property) (x interface{}, err error) {
	switch prop.Arity() {
	case graph.One:
		return ctx.readOneValue(prop.Type())
	case graph.UnorderedSet:
		return ctx.readValues(prop.Type())
	case graph.OrderedSet:
		return ctx.readValues(prop.Type())
	default:
		panic("unknown arity")
	}
}

func (ctx *dszContext) readOneValue(propertyType graph.PropertyType) (interface{}, error) {
	switch propertyType {
	case graph.TyString:
		return ctx.readString(), nil
	case graph.TyDouble:
		return ctx.readFloat64(), nil
	case graph.TyBool:
		return ctx.readBool(), nil
	case graph.TyInteger:
		panic("Can't handle integer values yet")
	case graph.TyObject:
		panic("Can't handle object values yet")
	default:
		panic("Unknown PropertyType")
	}
}

func (ctx *dszContext) readValues(propertyType graph.PropertyType) (interface{}, error) {
	switch propertyType {
	case graph.TyString:
		n := ctx.mustReadVarint()
		res := make([]string, n)
		for i := 0; i < n; i++ {
			res[i] = ctx.readString()
		}
		return res, nil
	case graph.TyDouble:
		n := ctx.mustReadVarint()
		res := make([]float64, n)
		for i := 0; i < n; i++ {
			res[i] = ctx.readFloat64()
		}
		return res, nil
	case graph.TyBool:
		n := ctx.mustReadVarint()
		res := make([]bool, n)
		for i := 0; i < n; i++ {
			res[i] = ctx.readBool()
		}
		return res, nil
	case graph.TyInteger:
		panic("Can't handle integer values yet")
	case graph.TyObject:
		panic("Can't handle object values yet")
	default:
		panic("Unknown PropertyType")
	}
}

func (ctx *dszContext) mustReadByte() byte {
	b, err := ctx.r.ReadByte()
	if err != nil {
		panic(err)
	}
	return b
}

func (ctx *dszContext) mustReadUvarint64() uint64 {
	x, err := binary.ReadUvarint(ctx.r)
	if err != nil {
		panic(err)
	}
	return x
}

func (ctx *dszContext) mustReadVarint64() int64 {
	x, err := binary.ReadVarint(ctx.r)
	if err != nil {
		panic(err)
	}
	return x
}

func (ctx *dszContext) mustReadVarint() int {
	return int(ctx.mustReadVarint64())
}

func (ctx *dszContext) readString() string {
	return string(ctx.readByteSlice())
}

func (ctx *dszContext) readByteSlice() []byte {
	n := ctx.mustReadVarint()
	return ctx.readNBytes(n)
}

func (ctx *dszContext) readNBytes(n int) []byte {
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = ctx.mustReadByte()
	}
	return res
}

func (ctx *dszContext) readFloat64() float64 {
	bytes := ctx.readNBytes(8)
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func (ctx *dszContext) readBool() bool {
	x := ctx.mustReadByte()
	switch x {
	case 0:
		return false
	case 1:
		return true
	default:
		panic("unexpected value for bool")
	}
}
