package xrb

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/regen-network/regen-ledger/x/schema"
	"io"
	"strconv"
)

type Serializer struct {
	properties map[string]propInfo
}

func (sz *Serializer) SerializeGraph(g Graph, w io.Writer) (hash []byte, err error) {
	panic("TODO")
}

type szContext struct {
	*Serializer
	w bufio.Writer
	//resGraph graph
	//curResNode node
	hashText hashAcc
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
	var last string
	for _, n := range nodes {
		if n == "" {
			panic("node ID cannot be nil")
		}
		if last != "" {
			if n < last {
				panic("nodes not in sorted order") // Node implementation error so panic
			}
			if last == n {
				panic("duplicate node ID")
			}
			last = n
		}
		s.serializeNode(false, g.GetNode(n))
	}
}

func (s *szContext) serializeNode(root bool, n Node) {
	if !root {
		writeString(s.w, n.ID())
	}
	id := n.ID()
	for _, url := range n.Properties() {
		if root {
			s.hashText.write("_:_\n")
		} else {
			s.hashText.writeIRI(id)
			s.hashText.write("\n")
		}
		s.writeProperty(s.w, url, n.GetProperty(url))
	}
}

func writeVarUint64(w bufio.Writer, x uint64) {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, x)
	mustWrite(w, buf[:n])
}

func writeVarInt(w bufio.Writer, x int) {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, int64(x))
	mustWrite(w, buf[:n])
}

func mustWrite(w bufio.Writer, buf []byte) {
	_, err := w.Write(buf)
	if err != nil {
		panic(err)
	}
}

func writeByte(w bufio.Writer, x byte) {
	mustWrite(w, []byte{x})
}

func writeString(w bufio.Writer, str string) {
	panic("TODO")
}

func writeBool(w bufio.Writer, x bool) {
	panic("TODO")
}

func writeFloat64(w bufio.Writer, x float64) {
	panic("TODO")
}

func (s *szContext) writeProperty(w bufio.Writer, name string, value interface{}) {
	info, found := s.properties[name]
	if !found {
		panic("TODO")
	}
	// PropertyID's get prefixed with byte 0
	writeByte(w, 0)
	writeVarUint64(w, uint64(info.id))

	s.writePropertyValue(w, info.def.Arity, info.def.PropertyType, value)
}

func (s *szContext) writePropertyValue(w bufio.Writer, arity schema.Arity, ty schema.PropertyType, value interface{}) {
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

func (s *szContext) writePropertyMany(ty schema.PropertyType, value interface{}, ordered bool) {
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

type hashAcc struct {
	bufio.Writer
}

func (ha *hashAcc) write(x string) {
	_, err := ha.WriteString(x)
	if err != nil {
		panic(err)
	}
}

func (ha *hashAcc) writeStringLiteral(x string, typeIri string, lang string) {
	ha.write("\"")
	ha.write(strconv.Quote(x))
	ha.write("\"")
	if typeIri != "" {
		if lang != "" {
			panic("cannot specify both a data type IRI and a language tag")
		}
		ha.write("^^")
		ha.writeIRI(typeIri)
	} else if lang != "" {
		ha.write("@")
		ha.write(lang)
	}
}

func (ha *hashAcc) writeIRI(x string) {
	ha.write("<")
	ha.write(x)
	ha.write("> ")
}

func (ha *hashAcc) finishLine() {
	ha.write(".\n")
}
