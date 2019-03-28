package xrb

import (
	"bytes"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/schema"
	"golang.org/x/crypto/blake2b"
	"net/url"
	"strconv"
)

type Property interface {
	ID() schema.PropertyID
	URI() *url.URL
	Arity() schema.Arity
	Type() schema.PropertyType
	fmt.Stringer
}

type Graph interface {
	Nodes() []types.HasURI
	RootNode() Node
	GetNode(id types.HasURI) Node
	Hash() []byte
	fmt.Stringer
}

type Node interface {
	ID() types.HasURI
	Properties() []Property
	GetProperty(property Property) interface{}
	// TODO GetClasses() []string
	fmt.Stringer
}

type AccAddressID struct {
	sdk.AccAddress
}

func (id AccAddressID) URI() *url.URL {
	uri, err := url.Parse(id.String())
	if err != nil {
		panic(err)
	}
	return uri
}

type HashID struct {
	fragment string
}

func (id HashID) String() string {
	return fmt.Sprintf("_:_.%s", id.fragment)
}

func (id HashID) URI() *url.URL {
	// TODO this should really be a blank node or the actual hash of the graph plus the fragment
	uri, err := url.Parse(fmt.Sprintf("root:#%s", id.fragment))
	if err != nil {
		panic(err)
	}
	return uri
}

type graph struct {
	rootNode  *node
	nodeNames []types.HasURI
	nodes     map[string]*node
}

func (g graph) String() string {
	w := new(bytes.Buffer)
	if g.rootNode != nil {
		w.WriteString(g.rootNode.String())
		w.WriteString("\n\n")
	}
	for _, n := range g.nodes {
		w.WriteString(n.String())
		w.WriteString("\n\n")
	}
	return w.String()
}

func (g graph) Hash() []byte {
	h, err := blake2b.New256(nil)
	if err != nil {
		panic(err)
	}
	h.Write([]byte(g.String()))
	return h.Sum(nil)
}

type node struct {
	id types.HasURI
	// TODO classes    []string
	propertyNames []Property
	properties    map[schema.PropertyID]interface{}
}

func (n node) String() string {
	w := new(bytes.Buffer)
	if n.id == nil {
		w.WriteString("_:_")
	} else {
		w.WriteString(n.id.String())
	}
	first := true
	for _, p := range n.propertyNames {
		if !first {
			w.WriteString(" ;")
			first = false
		}
		w.WriteString("\n")
		w.WriteString("  <")
		w.WriteString(p.URI().String())
		w.WriteString("> ")
		vStr, err := ValidatePrintValue(p, n.properties[p.ID()])
		if err != nil {
			panic(err)
		}
		w.WriteString(vStr)
	}
	w.WriteString(" .")
	return w.String()
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
		case schema.UnorderedSet:
			first := false
			for _, v := range value {
				if !first {
					w.WriteString(" ,")
					first = false
				}
				w.WriteString(strconv.Quote(v))
			}
		case schema.OrderedSet:
			w.WriteString("( ")
			for _, v := range value {
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
		case schema.UnorderedSet:
			first := false
			for _, v := range value {
				if !first {
					w.WriteString(" ,")
					first = false
				}
				w.WriteString(fmt.Sprintf(`"%f"^^<http://www.w3.org/2001/XMLSchema#double>`, v))
			}
		case schema.OrderedSet:
			w.WriteString("( ")
			for _, v := range value {
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
		case schema.OrderedSet:
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

type property struct {
	schema.PropertyDefinition
	id  schema.PropertyID
	uri *url.URL
}

func (p property) ID() schema.PropertyID {
	return p.id
}

func (p property) URI() *url.URL {
	return p.uri
}

func (p property) Arity() schema.Arity {
	return p.PropertyDefinition.Arity
}

func (p property) Type() schema.PropertyType {
	return p.PropertyDefinition.PropertyType
}

func (g graph) RootNode() Node {
	return g.rootNode
}

func (g graph) Nodes() []types.HasURI {
	return g.nodeNames
}

func (g graph) GetNode(url types.HasURI) Node {
	return g.nodes[url.URI().String()]
}

func (n node) ID() types.HasURI {
	return n.id
}

func (n node) Properties() []Property {
	return n.propertyNames
}

func (n node) GetProperty(property Property) interface{} {
	return n.properties[property.ID()]
}

//func (n node) GetClasses() []string {
//	panic("implement me")
//}
