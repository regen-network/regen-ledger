package binary

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/regen-network/regen-ledger/graph"
	"github.com/regen-network/regen-ledger/types"
	"io"
	"math"
)

func SerializeGraph(schema SchemaResolver, g graph.Graph, w io.Writer) error {
	ctx := &szContext{schema, bufio.NewWriter(w)}
	err := ctx.serializeGraph(g)
	if err != nil {
		return err
	}
	err = ctx.w.Flush()
	if err != nil {
		return err
	}
	return nil
}

type szContext struct {
	resolver SchemaResolver
	w        *bufio.Writer
}

func (s *szContext) serializeGraph(g graph.Graph) error {
	// Write the format version
	s.writeVarUint64(0)
	// TODO add chain height and maybe highest known property id
	r := g.RootNode()
	if r != nil {
		s.writeByte(1) // 1 for root node
		err := s.serializeNode(true, r)
		if err != nil {
			return err
		}
	} else {
		s.writeByte(0) // 0 for no root node
	}
	nodes := g.Nodes()
	s.writeVarInt(len(nodes))
	for _, n := range nodes {
		if n == nil {
			panic("node ID cannot be nil")
		}
		err := s.serializeNode(false, g.GetNode(n))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *szContext) serializeNode(root bool, n graph.Node) error {
	if !root {
		err := s.writeID(n.ID())
		if err != nil {
			return err
		}
	}
	//id := n.ID()
	props := n.Properties()
	s.writeVarInt(len(props))
	for _, url := range props {
		err := s.writeProperty(s.w, url, n.GetProperty(url))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *szContext) writeID(id types.HasURI) error {
	switch id := id.(type) {
	case AccAddressID:
		s.writeByte(prefixAccAddress)
		s.writeByteSlice(id.AccAddress)
	case types.GeoAddress:
		s.writeByte(prefixGeoAddress)
		s.writeByteSlice(id)
		// TODO
	case HashID:
		s.writeByte(prefixHashID)
		s.writeString(id.fragment)
	default:
		return fmt.Errorf("unexpected ID %s", id.String())
	}
	return nil
}

func (s *szContext) writeProperty(w *bufio.Writer, p graph.Property, value interface{}) error {
	// PropertyID's get prefixed with byte 0
	s.writeByte(prefixPropertyID)
	id := s.resolver.GetPropertyID(p)
	if id == 0 {
		return fmt.Errorf("can't resolve property %s in schema", p.URI())
	}
	s.writeVarInt64(int64(id))
	s.writePropertyValue(p.Arity(), p.Type(), value)
	return nil
}

func (s *szContext) writeVarUint64(x uint64) {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, x)
	mustWrite(s.w, buf[:n])
}

func (s *szContext) writeVarInt64(x int64) {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, x)
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
	s.writeByteSlice([]byte(str))
}

func (s *szContext) writeByteSlice(bytes []byte) {
	s.writeVarInt(len(bytes))
	mustWrite(s.w, bytes)
}

func (s *szContext) writeBool(x bool) {
	if !x {
		s.writeByte(0)
	} else {
		s.writeByte(1)
	}
}

func (s *szContext) writeFloat64(x float64) {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(x))
	mustWrite(s.w, buf[:])
}

func (s *szContext) writePropertyValue(arity graph.Arity, ty graph.PropertyType, value interface{}) {
	switch arity {
	case graph.One:
		s.writePropertyOne(ty, value)
	case graph.UnorderedSet:
		s.writePropertyMany(ty, value, false)
	case graph.OrderedSet:
		s.writePropertyMany(ty, value, true)
	default:
		panic("unknown arity")

	}
}

func (s *szContext) writePropertyOne(ty graph.PropertyType, value interface{}) {
	switch ty {
	case graph.TyString:
		x, ok := value.(string)
		if !ok {
			panic(fmt.Sprintf("Expected string value, got %+v", value))
		}
		s.writeString(x)
	case graph.TyDouble:
		x, ok := value.(float64)
		if !ok {
			panic(fmt.Sprintf("Expected float64 value, got %+v", value))
		}
		s.writeFloat64(x)
	case graph.TyBool:
		x, ok := value.(bool)
		if !ok {
			panic(fmt.Sprintf("Expected bool value, got %+v", value))
		}
		s.writeBool(x)
	case graph.TyInteger:
		panic("Can't handle integer values yet")
	case graph.TyObject:
		panic("Can't handle object values yet")
	default:
		panic("Unknown PropertyType")
	}
}

func (s *szContext) writePropertyMany(ty graph.PropertyType, value interface{}, ordered bool) {
	switch ty {
	case graph.TyString:
		arr, ok := value.([]string)
		if !ok {
			panic(fmt.Sprintf("Expected []string value, got %+v", value))
		}
		s.writeVarInt(len(arr))
		for _, x := range arr {
			s.writeString(x)
		}
	case graph.TyDouble:
		arr, ok := value.([]float64)
		if !ok {
			panic(fmt.Sprintf("Expected []float64 value, got %+v", value))
		}
		s.writeVarInt(len(arr))
		for _, x := range arr {
			s.writeFloat64(x)
		}
	case graph.TyBool:
		arr, ok := value.([]bool)
		if !ok {
			panic(fmt.Sprintf("Expected []bool value, got %+v", value))
		}
		s.writeVarInt(len(arr))
		for _, x := range arr {
			s.writeBool(x)
		}
	case graph.TyInteger:
		panic("Can't handle integer values yet")
	case graph.TyObject:
		panic("Can't handle object values yet")
	default:
		panic("Unknown PropertyType")
	}
}
