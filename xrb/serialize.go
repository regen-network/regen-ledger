package xrb

import (
	"encoding/binary"
	"fmt"
	"github.com/regen-network/regen-ledger/x/schema"
	"io"
)

type Serializer struct {
	properties map[string]propInfo
}

func (sz *Serializer) SerializeGraph(g Graph, w io.Writer) (hash []byte, err error) {
	panic("TODO")
}

type szContext struct {
	*Serializer
	w io.Writer
	//resGraph graph
	//curResNode node
	hashText string
}

type propInfo struct {
	id  schema.PropertyID
	def schema.PropertyDefinition
}

//type graph struct {
//	rootNode *node
//	nodes map[string]*node
//}
//
//type node struct {
//	id         string
//	// TODO classes    []string
//	properties map[string]interface{}
//}

func (s *szContext) serializeGraph(g Graph) {
	// Write the format version
	writeVarUint64(s.w, 0)
	// TODO add chain height and maybe highest known property id
	r := g.RootNode()
	if r != nil {
		writeByte(s.w, 1) // 1 for root node
		s.serializeNode(true, r)
	} else {
		writeByte(s.w, 0) // 0 for no root node
	}
	nodes := g.Nodes()
	writeVarInt(s.w, len(nodes))
	for _, n := range nodes {
		s.serializeNode(false, g.GetNode(n))
	}
}

func (s *szContext) serializeNode(root bool, n Node) {
	if !root {
		writeString(s.w, n.ID())
	}
	for _, url := range n.Properties() {
		s.writeProperty(s.w, url, n.GetProperty(url))
	}
}

func writeVarUint64(w io.Writer, x uint64) {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, x)
	mustWrite(w, buf[:n])
}

func writeVarInt(w io.Writer, x int) {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, int64(x))
	mustWrite(w, buf[:n])
}

func mustWrite(w io.Writer, buf []byte) {
	_, err := w.Write(buf)
	if err != nil {
		panic(err)
	}
}

func writeByte(w io.Writer, x byte) {
	mustWrite(w, []byte{x})
}

func writeString(w io.Writer, str string) {
	panic("TODO")
}

func writeBool(w io.Writer, x bool) {
	panic("TODO")
}

func writeFloat64(w io.Writer, x float64) {
	panic("TODO")
}

func (s *szContext) writeProperty(w io.Writer, name string, value interface{}) {
	info, found := s.properties[name]
	if !found {
		panic("TODO")
	}
	// PropertyID's get prefixed with byte 0
	writeByte(w, 0)
	writeVarUint64(w, uint64(info.id))

	s.writePropertyValue(w, info.def.Many, info.def.PropertyType, value)
}

func (s *szContext) writePropertyValue(w io.Writer, many bool, ty schema.PropertyType, value interface{}) {
	if !many {
		s.writePropertyOne(ty, value)
	} else {
		s.writePropertyMany(ty, value)
	}
}

func (s *szContext) writePropertyOne(ty schema.PropertyType, value interface{}) {
	switch ty {
	case schema.TyString:
		x, ok := value.(string)
		if !ok {
			panic(fmt.Sprintf("Expected string value, got %+v", value))
		}
		writeString(s.w, x)
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

func (s *szContext) writePropertyMany(ty schema.PropertyType, value interface{}) {
	switch ty {
	case schema.TyString:
		arr, ok := value.([]string)
		if !ok {
			panic(fmt.Sprintf("Expected []string value, got %+v", value))
		}
		writeVarInt(s.w, len(arr))
		for _, x := range arr {
			writeString(s.w, x)
		}
	case schema.TyDouble:
		arr, ok := value.([]float64)
		if !ok {
			panic(fmt.Sprintf("Expected []float64 value, got %+v", value))
		}
		writeVarInt(s.w, len(arr))
		for _, x := range arr {
			writeFloat64(s.w, x)
		}
	case schema.TyBool:
		arr, ok := value.([]bool)
		if !ok {
			panic(fmt.Sprintf("Expected []bool value, got %+v", value))
		}
		writeVarInt(s.w, len(arr))
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
