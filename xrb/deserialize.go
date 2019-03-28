package xrb

import (
	"encoding/binary"
	"fmt"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/schema"
	"io"
)

func DeserializeGraph(resolver SchemaResolver, r io.ByteScanner) (g Graph, hash []byte, err error) {
	ctx := &dszContext{resolver, r, newHasher(), 0}
	g, err = ctx.readGraph()
	if err != nil {
		return nil, nil, err
	}
	hash = ctx.h.hash()
	return g, hash, nil
}

type dszContext struct {
	resolver SchemaResolver
	r        io.ByteScanner
	h        *hasher
	version  uint64
}

func (ctx *dszContext) readGraph() (g *graph, err error) {
	ctx.version = ctx.mustReadUvarint()
	haveRootNode := ctx.mustReadByte()
	g = &graph{nodeNames: []types.HasURI{}, nodes: make(map[string]*node)}
	if haveRootNode == 1 {
		r, err := ctx.readNodeProperties()
		if err != nil {
			return nil, err
		}
		g.rootNode = r
	}
	nNodes := ctx.mustReadVarint()
	for i := int64(0); i < nNodes; i++ {
		n, err := ctx.readNode()
		// TODO verify ordering
		// hash
		if err != nil {
			return nil, err
		}
		g.nodeNames = append(g.nodeNames, n.id)
		g.nodes[n.id.String()] = n
	}
	return g, nil
}

func (ctx *dszContext) readNode() (n *node, err error) {
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

func (ctx *dszContext) readNodeProperties() (n *node, err error) {
	n = &node{properties: make(map[schema.PropertyID]interface{}), propertyNames: []Property{}}
	nProps := ctx.mustReadVarint()
	for i := int64(0); i < nProps; i++ {
		url, value, err := ctx.readProperty()
		// TODO verify ordering
		// hash
		if err != nil {
			return nil, err
		}
		n.propertyNames = append(n.propertyNames, url)
		n.properties[url.ID()] = value
	}
	return n, nil
}

func (ctx *dszContext) readProperty() (prop Property, value interface{}, err error) {
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

func (ctx *dszContext) readValue(prop Property) (x interface{}, err error) {
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
