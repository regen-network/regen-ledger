package graph

import (
	"bytes"
	"fmt"
	"github.com/regen-network/regen-ledger/types"
	"golang.org/x/crypto/blake2b"
	"net/url"
	"strconv"
)

// PropertyType indicates the data type of a property
type PropertyType uint32

const (
	// TyObject is the PropertyType of objects or nodes
	TyObject PropertyType = iota
	// TyString is the PropertyType of strings
	TyString
	// TyInteger is the PropertyType of arbitrary-precision integers
	TyInteger
	// TyDouble is the PropertyType of double-precision floating point
	TyDouble
	// TyBool is the PropertyType of boolean values
	TyBool
)

type Arity int

const (
	One Arity = iota
	UnorderedSet
	OrderedSet
)

type Property interface {
	URI() *url.URL
	Arity() Arity
	Type() PropertyType
	fmt.Stringer
}

type Graph interface {
	Nodes() []types.HasURI
	RootNode() Node
	GetNode(id types.HasURI) Node

	WithNode(node Node)
	WithoutNode(id types.HasURI)

	fmt.Stringer
}

type Node interface {
	ID() types.HasURI
	Properties() []Property
	GetProperty(property Property) interface{}
	Classes() []Class

	SetID(id types.HasURI)
	SetProperty(property Property, value interface{})
	DeleteProperty(property Property)

	fmt.Stringer
}

type Class interface {
	// RequiredProperties returns a slice of properties with arity one that are required
	RequiredProperties() []Property
	fmt.Stringer
}

func (ty PropertyType) String() string {
	names := [...]string{
		"Object",
		"String",
		"Integer",
		"Double",
		"Bool",
	}
	if int(ty) > len(names) {
		return ""
	}
	return names[ty]
}

func (a Arity) String() string {
	names := [...]string{
		"One",
		"Unordered",
		"Ordered",
	}
	if int(a) > len(names) {
		return ""
	}
	return names[a]
}

func CanonicalString(g Graph) (string, error) {
	w := new(bytes.Buffer)
	r := g.RootNode()
	if r != nil {
		s, err := CanonicalNodeString(r)
		if err != nil {
			return "", err
		}
		w.WriteString(s)
		w.WriteString("\n\n")
	}
	for _, n := range g.Nodes() {
		s, err := CanonicalNodeString(g.GetNode(n))
		if err != nil {
			return "", err
		}
		w.WriteString(s)
		w.WriteString("\n\n")
	}
	return w.String(), nil
}

func CanonicalNodeString(n Node) (string, error) {
	w := new(bytes.Buffer)
	if n.ID() == nil {
		w.WriteString("_:_")
	} else {
		w.WriteString(n.ID().String())
	}
	first := true
	for _, p := range n.Properties() {
		if !first {
			w.WriteString(" ;")
			first = false
		}
		w.WriteString("\n")
		w.WriteString("  <")
		w.WriteString(p.URI().String())
		w.WriteString("> ")
		vStr, err := ValidatePrintValue(p, n.GetProperty(p))
		if err != nil {
			return "", err
		}
		w.WriteString(vStr)
	}
	w.WriteString(" .")
	return w.String(), nil
}

func Hash(g Graph) []byte {
	h, err := blake2b.New256(nil)
	if err != nil {
		panic(err)
	}
	h.Write([]byte(g.String()))
	return h.Sum(nil)
}

func ValidatePrintValue(prop Property, value interface{}) (string, error) {
	switch value := value.(type) {
	case string:
		return strconv.Quote(value), nil
	case float64:
		return fmt.Sprintf(`"%f"^^<http://www.w3.org/2001/XMLSchema#double>`, value), nil
	case bool:
		return fmt.Sprintf("%t", value), nil
	case []string:
		w := new(bytes.Buffer)
		switch prop.Arity() {
		case UnorderedSet:
			first := false
			for _, v := range value {
				if !first {
					w.WriteString(" ,")
					first = false
				}
				w.WriteString(strconv.Quote(v))
			}
		case OrderedSet:
			w.WriteString("(")
			first := false
			for _, v := range value {
				if !first {
					w.WriteString(" ")
					first = false
				}
				w.WriteString(strconv.Quote(v))
				w.WriteString(" ")
			}
			w.WriteString(")")
		default:
			return "", fmt.Errorf("unexpected arity %s", prop.Arity().String())
		}
		return w.String(), nil
	case []float64:
		w := new(bytes.Buffer)
		switch prop.Arity() {
		case UnorderedSet:
			first := false
			for _, v := range value {
				if !first {
					w.WriteString(" ,")
					first = false
				}
				w.WriteString(fmt.Sprintf(`"%f"^^<http://www.w3.org/2001/XMLSchema#double>`, v))
			}
		case OrderedSet:
			w.WriteString("(")
			first := false
			for _, v := range value {
				if !first {
					w.WriteString(" ")
					first = false
				}
				w.WriteString(fmt.Sprintf(`"%f"^^<http://www.w3.org/2001/XMLSchema#double>`, v))
				w.WriteString(" ")
			}
			w.WriteString(")")
		default:
			return "", fmt.Errorf("unexpected arity %s", prop.Arity().String())
		}
		return w.String(), nil
	case []bool:
		w := new(bytes.Buffer)
		switch prop.Arity() {
		case OrderedSet:
			w.WriteString("( ")
			for _, v := range value {
				w.WriteString(fmt.Sprintf(`%t`, v))
				w.WriteString(" ")
			}
			w.WriteString(")")
		default:
			return "", fmt.Errorf("unexpected arity %s", prop.Arity().String())
		}
		return w.String(), nil
	default:
		return "", fmt.Errorf("unexpected value %+v for property %s", value, prop.String())
	}
}
