package xrb

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/schema"
	"io"
)

func SerializeGraph(schema SchemaResolver, g Graph, w io.Writer) (hash []byte, err error) {
	ctx := &szContext{schema, bufio.NewWriter(w), newHasher()}
	err = ctx.serializeGraph(g)
	if err != nil {
		return nil, err
	}
	err = ctx.w.Flush()
	if err != nil {
		return nil, err
	}
	hash = ctx.hashText.hash()
	return hash, nil
}

type szContext struct {
	resolver SchemaResolver
	w        *bufio.Writer
	hashText *hasher
}

type propInfo struct {
	id  schema.PropertyID
	def schema.PropertyDefinition
}

func (s *szContext) serializeGraph(g Graph) error {
	// Write the format version
	s.writeVarUint64(0)
	// TODO add chain height and maybe highest known property id
	r := g.RootNode()
	if r != nil {
		s.writeByte(1) // 1 for root node
		s.serializeNode(true, r)
	} else {
		s.writeByte(0) // 0 for no root node
	}
	nodes := g.Nodes()
	s.writeVarInt(len(nodes))
	var last string
	for _, n := range nodes {
		if n == nil {
			panic("node ID cannot be nil")
		}
		uri := n.URI().String()
		if last != "" {
			if uri < last {
				panic("nodes not in sorted order") // Node implementation error so panic
			}
			if last == uri {
				panic("duplicate node ID")
			}
			last = uri
		}
		s.serializeNode(false, g.GetNode(n))
	}
	return nil
}

func (s *szContext) serializeNode(root bool, n Node) {
	if !root {
		s.writeID(n.ID())
	}
	//id := n.ID()
	props := n.Properties()
	s.writeVarInt(len(props))
	for _, url := range props {
		//if root {
		//	s.hashText.write("_:_\n")
		//} else {
		//	s.hashText.writeIRI(id)
		//	s.hashText.write("\n")
		//}
		s.writeProperty(s.w, url, n.GetProperty(url))
	}
}

func (s *szContext) writeID(id types.HasURI) {
	switch id := id.(type) {
	case AccAddressID:
		s.writeByte(prefixAccAddress)
		s.writeBytes(id.AccAddress)
	case types.GeoAddress:
		s.writeByte(prefixGeoAddress)
		s.writeBytes(id)
		// TODO
	case types.DataAddress:
		s.writeByte(prefixDataAddress)
		s.writeBytes(id)
		// TODO
	case HashID:
		s.writeByte(prefixHashID)
		s.writeString(id.String())
	default:
	}
}

func (s *szContext) writeProperty(w *bufio.Writer, p Property, value interface{}) {
	// PropertyID's get prefixed with byte 0
	s.writeByte(prefixPropertyID)
	s.writeVarUint64(uint64(p.ID()))

	s.writePropertyValue(p.Arity(), p.Type(), value)
}

func (s *szContext) writeVarUint64(x uint64) {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, x)
	mustWrite(s.w, buf[:n])
}

func (s *szContext) writeVarInt(x int) {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, int64(x))
	mustWrite(s.w, buf[:n])
}

func mustWrite(w *bufio.Writer, buf []byte) {
	_, err := w.Write(buf)
	if err != nil {
		panic(err)
	}
}

func (s *szContext) writeByte(x byte) {
	mustWrite(s.w, []byte{x})
}

func (s *szContext) writeString(str string) {
	panic("TODO")
}

func (s *szContext) writeBytes(bytes []byte) {
	panic("TODO")
}

func writeBool(w *bufio.Writer, x bool) {
	panic("TODO")
}

func writeFloat64(w *bufio.Writer, x float64) {
	panic("TODO")
}

func (s *szContext) writePropertyValue(arity schema.Arity, ty schema.PropertyType, value interface{}) {
	switch arity {
	case schema.One:
		s.writePropertyOne(ty, value)
	case schema.UnorderedSet:
		s.writePropertyMany(ty, value, false)
	case schema.OrderedSet:
		s.writePropertyMany(ty, value, true)
	default:
		panic("unknown arity")

	}
}

func (s *szContext) writePropertyOne(ty schema.PropertyType, value interface{}) {
	switch ty {
	case schema.TyString:
		x, ok := value.(string)
		if !ok {
			panic(fmt.Sprintf("Expected string value, got %+v", value))
		}
		s.writeString(x)
	case schema.TyDouble:
		x, ok := value.(float64)
		if !ok {
			panic(fmt.Sprintf("Expected float64 value, got %+v", value))
		}
		writeFloat64(s.w, x)
	case schema.TyBool:
		x, ok := value.(bool)
		if !ok {
			panic(fmt.Sprintf("Expected bool value, got %+v", value))
		}
		writeBool(s.w, x)
	case schema.TyInteger:
		panic("Can't handle integer values yet")
	case schema.TyObject:
		panic("Can't handle object values yet")
	default:
	}
}

func (s *szContext) writePropertyMany(ty schema.PropertyType, value interface{}, ordered bool) {
	switch ty {
	case schema.TyString:
		arr, ok := value.([]string)
		if !ok {
			panic(fmt.Sprintf("Expected []string value, got %+v", value))
		}
		s.writeVarInt(len(arr))
		for _, x := range arr {
			s.writeString(x)
		}
	case schema.TyDouble:
		arr, ok := value.([]float64)
		if !ok {
			panic(fmt.Sprintf("Expected []float64 value, got %+v", value))
		}
		s.writeVarInt(len(arr))
		for _, x := range arr {
			writeFloat64(s.w, x)
		}
	case schema.TyBool:
		arr, ok := value.([]bool)
		if !ok {
			panic(fmt.Sprintf("Expected []bool value, got %+v", value))
		}
		s.writeVarInt(len(arr))
		for _, x := range arr {
			writeBool(s.w, x)
		}
	case schema.TyInteger:
		panic("Can't handle integer values yet")
	case schema.TyObject:
		panic("Can't handle object values yet")
	default:
	}
}
