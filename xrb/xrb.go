/*
Package xrb defines an efficient RDF binary serialization format
XRN (XRN RDF Binary) that integrates with the Regen Ledger schema module
*/
package xrb

import (
	"encoding/binary"
	"fmt"
	"github.com/regen-network/regen-ledger/x/schema"
	"io"
)

type Node interface {
	ID() string
	GetProperty(url string) interface{}
}

type serializer struct {
	properties map[string]propInfo
}

type propInfo struct {
	id  schema.PropertyID
	def schema.PropertyDefinition
}

type graph struct {
	nodes []node
}

type node struct {
	id         string
	properties map[string]interface{}
}

func (s serializer) SerializeGraph(w io.Writer, g graph) {
	// Write the format version
	mustWriteVarUint64(w, 0)
	// TODO add chain height and maybe highest known property id
	for _, n := range g.nodes {
		s.SerializeNode(w, n)
	}
}

func (s serializer) SerializeNode(w io.Writer, n node) {
	writeString(w, n.id)
	for name, val := range n.properties {
		s.writeProperty(w, name, val)
	}
}

func mustWriteVarUint64(w io.Writer, x uint64) {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, x)
	mustWrite(w, buf[:n])
}

func mustWriteVarInt(w io.Writer, x int) {
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

func writeString(w io.Writer, str string) {
	panic("TODO")
}

func writeBool(w io.Writer, x bool) {
	panic("TODO")
}

func writeFloat64(w io.Writer, x float64) {
	panic("TODO")
}

func (s serializer) writeProperty(w io.Writer, name string, value interface{}) {
	info, found := s.properties[name]
	if !found {
		panic("TODO")
	}

	mustWriteVarUint64(w, uint64(info.id))

	s.writePropertyValue(w, info.def.Many, info.def.PropertyType, value)
}

func (s serializer) writePropertyValue(w io.Writer, many bool, ty schema.PropertyType, value interface{}) {
	if !many {
		s.writePropertyOne(w, ty, value)
	} else {
		s.writePropertyMany(w, ty, value)
	}
}

func (s serializer) writePropertyOne(w io.Writer, ty schema.PropertyType, value interface{}) {
	switch ty {
	case schema.TyString:
		x, ok := value.(string)
		if !ok {
			panic(fmt.Sprintf("Expected string value, got %+v", value))
		}
		writeString(w, x)
	case schema.TyDouble:
		x, ok := value.(float64)
		if !ok {
			panic(fmt.Sprintf("Expected float64 value, got %+v", value))
		}
		writeFloat64(w, x)
	case schema.TyBool:
		x, ok := value.(bool)
		if !ok {
			panic(fmt.Sprintf("Expected bool value, got %+v", value))
		}
		writeBool(w, x)
	case schema.TyInteger:
		panic("Can't handle integer values yet")
	case schema.TyObject:
		panic("Can't handle object values yet")
	default:
	}
}

func (s serializer) writePropertyMany(w io.Writer, ty schema.PropertyType, value interface{}) {
	switch ty {
	case schema.TyString:
		arr, ok := value.([]string)
		if !ok {
			panic(fmt.Sprintf("Expected []string value, got %+v", value))
		}
		mustWriteVarInt(w, len(arr))
		for _, x := range arr {
			writeString(w, x)
		}
	case schema.TyDouble:
		arr, ok := value.([]float64)
		if !ok {
			panic(fmt.Sprintf("Expected []float64 value, got %+v", value))
		}
		mustWriteVarInt(w, len(arr))
		for _, x := range arr {
			writeFloat64(w, x)
		}
	case schema.TyBool:
		arr, ok := value.([]bool)
		if !ok {
			panic(fmt.Sprintf("Expected []bool value, got %+v", value))
		}
		mustWriteVarInt(w, len(arr))
		for _, x := range arr {
			writeBool(w, x)
		}
	case schema.TyInteger:
		panic("Can't handle integer values yet")
	case schema.TyObject:
		panic("Can't handle object values yet")
	default:
	}
}
