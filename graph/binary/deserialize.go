package binary

import (
	"encoding/binary"
	"fmt"
	"github.com/regen-network/regen-ledger/graph"
	"github.com/regen-network/regen-ledger/graph/impl"
	"github.com/regen-network/regen-ledger/x/schema"
	"io"
)

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
	ctx.version = ctx.mustReadUvarint()
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
	for i := int64(0); i < nNodes; i++ {
		n, err := ctx.readNode()
		if err != nil {
			return nil, err
		}
		g.WithNode(n)
	}
	return g, nil
}

func (ctx *dszContext) readNode() (n graph.Node, err error) {
	panic("TODO")
	//id, err := ctx.readString()
	//if err != nil {
	//	return nil, err
	//}
	//n, err = ctx.readNodeProperties()
	//if err != nil {
	//	return nil, err
	//}
	//n.id = id
	//return n, nil
}

func (ctx *dszContext) readNodeProperties() (n graph.Node, err error) {
	n = impl.NewNode(nil)
	nProps := ctx.mustReadVarint()
	for i := int64(0); i < nProps; i++ {
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
	if prefix != 0 {
		return nil, nil, fmt.Errorf("unexpected property ID prefix %d", prefix)
	}
	id, err := binary.ReadUvarint(ctx.r)
	if err != nil {
		return nil, nil, err
	}
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
	panic("TODO")
}

func (ctx *dszContext) readString() (x string, err error) {
	panic("TODO")
}

func (ctx *dszContext) mustReadByte() byte {
	b, err := ctx.r.ReadByte()
	if err != nil {
		panic(err)
	}
	return b
}

func (ctx *dszContext) mustReadUvarint() uint64 {
	x, err := binary.ReadUvarint(ctx.r)
	if err != nil {
		panic(err)
	}
	return x
}

func (ctx *dszContext) mustReadVarint() int64 {
	x, err := binary.ReadVarint(ctx.r)
	if err != nil {
		panic(err)
	}
	return x
}
